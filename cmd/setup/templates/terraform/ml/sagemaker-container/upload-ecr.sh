
# The name of our algorithm
algorithm_name=leotest-scikit-gensim

region=$(aws configure get region)

$(aws ecr get-login --region $region --no-include-email)

account=$(aws sts get-caller-identity --query Account --output text)

docker push "${account}.dkr.ecr.${region}.amazonaws.com/${algorithm_name}"
