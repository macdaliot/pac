#cloud-boothook
#!/bin/bash

cloud-init-per once docker_options echo 'OPTIONS="--storage-opt dm.basesize=900G"' >> /etc/sysconfig/docker
echo ECS_CLUSTER=${ecs_cluster_name} >> /etc/ecs/ecs.config
