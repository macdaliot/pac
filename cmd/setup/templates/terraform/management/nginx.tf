#======================================================================================================================
# DNS
#======================================================================================================================
resource "aws_route53_record" "nginx" {
  zone_id = data.terraform_remote_state.dns.outputs.main_zone_id
  name    = "monitoring.${var.project_name}.${var.hosted_zone}"
  type    = "A"
  records = [aws_instance.nginx.public_ip]
  ttl     = 300
}

#======================================================================================================================
# SECURITY GROUP
#======================================================================================================================
resource "aws_security_group" "nginx" {
  name                   = "${var.project_name}-management-nginx"
  description            = "controls access to nginx"
  vpc_id                 = aws_vpc.management_vpc.id
  revoke_rules_on_delete = true

  ingress {
    protocol    = "tcp"
    from_port   = "22"
    to_port     = "22"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    protocol    = "tcp"
    from_port   = "80"
    to_port     = "80"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    protocol    = "tcp"
    from_port   = "443"
    to_port     = "443"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = var.project_name
  }
}

#======================================================================================================================
# INSTANCES
#======================================================================================================================
resource "aws_instance" "nginx" {
  # The connection block tells our provisioner how to
  # communicate with the resource (instance)
  connection {
    # The default username for our AMI
    user        = "ubuntu"
    type        = "ssh"
    private_key = data.terraform_remote_state.bootstrap.outputs.management_tls.private_key_pem
    host        = self.public_ip

    # The connection will use the local SSH agent for authentication.
  }

  associate_public_ip_address = true
  instance_type               = "t2.small"

  # Lookup the correct AMI based on the region
  # we specified 
  ami = data.aws_ami.ubuntu.id

  # The name of our SSH keypair we created above.
  key_name = "jumpbox-${var.project_name}"

  # Our Security group to allow HTTP and SSH access
  vpc_security_group_ids = [aws_security_group.nginx.id]

  # We're going to launch into the same subnet as our ELB. In a production
  # environment it's more common to have a separate private subnet for
  # backend instances.
  subnet_id = aws_subnet.private.0.id

  # Copy the certificate into the dir where nginx conf is expecting it to be
  provisioner "file" {
    content     = acme_certificate.certificate.certificate_pem
    destination = "/tmp/certificate.pem"
  }

  provisioner "file" {
    content     = acme_certificate.certificate.issuer_pem
    destination = "/tmp/issuer.pem"
  }

  provisioner "file" {
    source      = "./default.conf"
    destination = "/tmp/default.conf"
  }

  provisioner "file" {
    source      = ".htpasswd"
    destination = "/tmp/.htpasswd"
  }

  # We run a remote provisioner on the instance after creating it.
  # In this case, we just install nginx and start it. By default,
  # this should be on port 80
  provisioner "remote-exec" {
    inline = [
      "host=$(hostname)",
      "hostIp=$(hostname -I)",
      "sudo -- sh -c \"echo '$hostIp $host' >> /etc/hosts\"",
      "sudo apt-get -y update",
      "sudo apt-get -y install nginx",
      "sudo mkdir /etc/certs/",
      "sudo cat /tmp/issuer.pem >> /tmp/certificate.pem",
      "sudo mv /tmp/certificate.pem /etc/certs",
      "sudo unlink /etc/nginx/sites-enabled/default",
      "sudo mv /tmp/.htpasswd /etc/nginx/.htpasswd",
      "sudo touch /etc/nginx/.htpasswd",
      "sudo -- sh -c \"echo '${acme_certificate.certificate.private_key_pem}' > /etc/certs/privKey.pem\"",
      "sudo mv /tmp/default.conf /etc/nginx/sites-enabled/",
      "sudo sed -i \"s/SERVERNAME/monitoring.${var.project_name}.${var.hosted_zone}/g\" /etc/nginx/sites-enabled/default.conf",
      "sudo sed -i \"s/KIBANA_URL/${aws_elasticsearch_domain.es.endpoint}/g\" /etc/nginx/sites-enabled/default.conf",
      "RESOLVER_IP=$(grep -o \"[0-9]\\{1,3\\}\\.[0-9]\\{1,3\\}\\.[0-9]\\{1,3\\}\\.[0-9]\\{1,3\\}\" /etc/resolv.conf | head -n 1 $1)",
      "sudo sed -i \"s/RESOLVER_IP/$RESOLVER_IP/g\" /etc/nginx/sites-enabled/default.conf",
      "sudo service nginx start",
      "sudo nginx -s reload",
      "sudo update-rc.d nginx defaults",
    ]
  }
}

#----------------------------------------------------------------------------------------------------------------------
# LOAD BALANCING
#----------------------------------------------------------------------------------------------------------------------
resource "aws_lb_target_group" "nginx" {
  name        = "${var.project_name}-management-nginx"
  port        = "80"
  protocol    = "HTTP"
  vpc_id      = aws_vpc.management_vpc.id
  target_type = "instance"

  tags = {
    Name             = var.project_name
    pac-project-name = var.project_name
    environment      = "management"
  }
}

resource "aws_lb_target_group_attachment" "nginx" {
  target_group_arn = aws_lb_target_group.nginx.arn
  target_id        = aws_instance.nginx.id
  port             = 80
}