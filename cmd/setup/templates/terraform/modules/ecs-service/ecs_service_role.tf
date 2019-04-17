resource "aws_iam_role" "tr-ecs-service-role" {
    name                = "tr-ecs-service-role"
    path                = "/"
    assume_role_policy  = "${data.aws_iam_policy_document.tr-ecs-service-policy.json}"
}

resource "aws_iam_role_policy_attachment" "tr-ecs-service-role-attachment" {
    role       = "${aws_iam_role.tr-ecs-service-role.name}"
    policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceRole"
}

data "aws_iam_policy_document" "tr-ecs-service-policy" {
    statement {
        actions = ["sts:AssumeRole"]

        principals {
            type        = "Service"
            identifiers = ["ecs.amazonaws.com"]
        }
    }
}