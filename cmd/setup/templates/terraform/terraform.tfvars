az_count     = 3
# project_name = "testalicious"
app_count    = 1
app_port     = 8080
services = [
    {
        service_name = "sonar_db"
        app_image    = "118104210923.dkr.ecr.us-east-2.amazonaws.com/${var.project_name}-sonar-db"
    },
    {
        service_name = "sonar_app"
        app_image    = "118104210923.dkr.ecr.us-east-2.amazonaws.com/${var.project_name}-sonarqube"
    }
]