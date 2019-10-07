python setup.py bdist_wheel
# The name of our algorithm
algorithm_name=$3

#docker image prune -a -f

account=$(aws sts get-caller-identity --query Account --output text)

# Get the region defined in the current configuration (default to us-west-2 if none defined)
region=$(aws configure get region)

DOCKER_LOGIN=`aws ecr get-login --region $region --no-include-email`; eval $DOCKER_LOGIN

fullname="$1$2:$3"

# build docker image
docker build -t $fullname --build-arg 'py_version=3' --build-arg "region=$region" --build-arg $1  --build-arg 'pipeline_data_bucket=terraform.{{ .projectName }}-state' --build-arg 'project_name=bdso-fork' --build-arg "sagemaker_image_name=$2" --build-arg 'sagemaker_role=arn:aws:iam::342720212717:role/service-role/AmazonSageMaker-ExecutionRole-20190624T153300' --build-arg 'stack_name=qa-bdso-fork' --build-arg 'environment=dev' -f ./docker/0.20.0/base/Dockerfile.dockerfile .

docker push $fullname
# docker run -it -p 6689:6689 -p 22:22 -e ACCESS_KEY=$access_key -e SECRET_KEY=$secret_key 342720212717.dkr.ecr.us-east-2.amazonaws.com/scikit-sagemaker:latest