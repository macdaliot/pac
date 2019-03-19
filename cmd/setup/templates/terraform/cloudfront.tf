# DO NOT DELETE
# disabled until needed due to long create and destroy cycles
#
# module "pac_cloudfront_oai" {
#   source = "./modules/cloudfront/origin_access_identity"

#   project_name = "${var.project_name}"
# }

# module "pac_cloudfront" {
#   source = "./modules/cloudfront/distribution"

#   s3_brdn = "${module.pac_bucket.brdn}"
#   project_name = "${var.project_name}"
#   oia_path = "${module.pac_cloudfront_oai.oai_path}"
#   env = "integration"
#   price_class = "PriceClass_100"
#   host_name = "${var.hosted_zone}"
# }

