# resource "aws_vpc" "application" {
#   cidr_block = "10.1.0.0/16"

#   tags {
#     name = "${var.project_name}-application-vpc"
#   }
# }

# # Create var.az_count public subnets, each in a different AZ
# resource "aws_subnet" "public" {
#   count                   = "${var.az_count}"
#   cidr_block              = "${cidrsubnet(aws_vpc.application.cidr_block, 8, var.az_count + count.index)}"
#   availability_zone       = "${data.aws_availability_zones.available.names[count.index]}"
#   vpc_id                  = "${aws_vpc.application.id}"
#   map_public_ip_on_launch = false

#   tags {
#     name = "${var.project_name}-public-${count.index}"
#   }
# }

# # IGW for the public subnet
# resource "aws_internet_gateway" "gw" {
#   vpc_id = "${aws_vpc.application.id}"
# }

# # Route the public subnet traffic through the IGW
# resource "aws_route" "internet_access" {
#   route_table_id         = "${aws_vpc.application.main_route_table_id}"
#   destination_cidr_block = "0.0.0.0/0"
#   gateway_id             = "${aws_internet_gateway.gw.id}"
# }

# resource "aws_alb" "main" {
#   name            = "${var.project_name}-load-balancer"
#   subnets         = ["${aws_subnet.public.*.id}"]
#   security_groups = ["${aws_security_group.lb.id}"]
# }

# module "cnames" {
#   source = "./modules/route53_record"

#   zone_id = "${module.route53_zone.child_zone_id}"
#   names = ["api", "jenkins", "selenium", "sonar"]
#   type = "CNAME"
#   records = ["${aws_alb.main.dns_name}"]
# }

# #
# # https://www.terraform.io/docs/providers/aws/r/lb_listener.html
# # alb listener
# #
# resource "aws_lb_listener" "api" {
#   load_balancer_arn = "${aws_alb.main.arn}"
#   port              = "80"
#   protocol          = "HTTP"

#   default_action {
#     type             = "redirect"

#     redirect {
#       port        = "80"
#       protocol    = "HTTP"
#       status_code = "HTTP_301"
#       path        = "/#{host}/api"
#       query       = "#{query}"
#     }
#   }
# }

module "pac_bucket" {
  source = "./modules/s3"

  project_name = "${var.project_name}"
}

# resource "aws_route_table" "application_vpc" {
#   vpc_id = "${aws_vpc.application.id}"

#   route {
#     cidr_block = "${aws_vpc.management.cidr_block}"
#     vpc_peering_connection_id = "${module.peer_vpcs.id}"
#   }

#   tags = {
#     Name = "${var.project_name} vpc peering route table"
#   }
# }

# # DO NOT DELETE
# # disabled until needed due to long create and destroy cycles
# #
# # module "pac_cloudfront_oai" {
# #   source = "./modules/cloudfront/origin_access_identity"

# #   project_name = "${var.project_name}"
# # }

# # module "pac_cloudfront" {
# #   source = "./modules/cloudfront/distribution"

# #   s3_brdn = "${module.pac_bucket.brdn}"
# #   project_name = "${var.project_name}"
# #   oia_path = "${module.pac_cloudfront_oai.oai_path}"
# #   env = "integration"
# #   price_class = "PriceClass_100"
# #   host_name = "${var.hosted_zone}"
# # }