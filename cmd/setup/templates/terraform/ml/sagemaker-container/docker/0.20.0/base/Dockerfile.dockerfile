FROM ubuntu:18.04

ARG py_version
ARG region
ARG image_repo
ARG pipeline_data_bucket
ARG project_name
ARG sagemaker_image_name
ARG sagemaker_role
ARG stack_name
ARG environment

ENV ML_PIPELINE_ENVS="[\"dev\", \"e2e\"]" TEST="false" REGION=${region} IMAGE_REPO=${image_repo} PIPELINE_DATA_BUCKET=${pipeline_data_bucket}
ENV PROJECT_NAME=${project_name} SAGEMAKER_IMAGE_NAME=${sagemaker_image_name} SAGEMAKER_ROLE=${sagemaker_role} STACK_NAME=${stack_name} ENVIRONMENT=${environment}

# Validate that arguments are specified
RUN test $py_version || exit 1

# Install python and other scikit-learn runtime dependencies
# Dependency list from http://scikit-learn.org/stable/developers/advanced_installation.html#installing-build-dependencies
RUN apt-get update && \
    apt-get -y install build-essential libatlas-base-dev git wget curl nginx jq && \
    if [ $py_version -eq 2 ]; \
       then apt-get -y install python-dev python-setuptools \
                     python-numpy python-scipy libatlas3-base; \
       else apt-get -y install python3-dev python3-setuptools \
                     python3-numpy python3-scipy libatlas3-base; fi

# Install pip
RUN cd /tmp && \
     curl -O https://bootstrap.pypa.io/get-pip.py && \
     if [ $py_version -eq 2 ]; \
        then python2 get-pip.py; \
        else python3 get-pip.py; fi && \
     rm get-pip.py

# Python won’t try to write .pyc or .pyo files on the import of source modules
# Force stdin, stdout and stderr to be totally unbuffered. Good for logging
ENV PYTHONDONTWRITEBYTECODE=1 PYTHONUNBUFFERED=1 PYTHONIOENCODING=UTF-8 LANG=C.UTF-8 LC_ALL=C.UTF-8

# Install dependencies from pip
# Install Scikit-Learn; 0.20.0 supports both python 2.7+ and 3.4+
RUN pip install --no-cache -I scikit-learn==0.20.0 retrying

LABEL com.amazonaws.sagemaker.capabilities.accept-bind-to-port=true

# Added for dedupe
RUN apt-get -y install libffi6 libffi-dev

COPY dist/sagemaker_sklearn_container-1.0-py2.py3-none-any.whl /sagemaker_sklearn_container-1.0-py2.py3-none-any.whl
RUN pip install --no-cache /sagemaker_sklearn_container-1.0-py2.py3-none-any.whl && \
    rm /sagemaker_sklearn_container-1.0-py2.py3-none-any.whl

ENV SAGEMAKER_TRAINING_MODULE sagemaker_sklearn_container.training:main
ENV SAGEMAKER_SERVING_MODULE sagemaker_sklearn_container.serving:main

RUN apt-get update
RUN apt install openjdk-8-jdk -y
RUN apt-get install ssh -y
RUN apt-get install -y python3-venv
RUN useradd -m -s /bin/bash jenkins
RUN pip3 install awscli
RUN mkdir /home/jenkins/.aws/
COPY jenkins /home/jenkins/.ssh/id_rsa
COPY jenkins.pub /home/jenkins/.ssh/id_rsa.pub
RUN echo "[default]" > /home/jenkins/.aws/credentials
RUN echo "[default]" > /home/jenkins/.aws/config
RUN echo "region = "${REGION} >> /home/jenkins/.aws/config
RUN echo "output = json" >> /home/jenkins/.aws/config
RUN cat /home/jenkins/.ssh/id_rsa.pub > /home/jenkins/.ssh/authorized_keys
RUN echo "PubkeyAuthentication yes" >> /etc/ssh/sshd_config
RUN echo "PasswordAuthentication no" >> /etc/ssh/sshd_config
RUN echo "Port 6689" >> /etc/ssh/sshd_config
RUN mkdir -p /jdk; cp -r /usr/lib/jvm/java-1.8.0-openjdk-amd64/* /jdk
RUN chown -R jenkins:jenkins /home/jenkins/
RUN apt-get install -y unzip
RUN wget https://releases.hashicorp.com/terraform/0.12.7/terraform_0.12.7_linux_amd64.zip; unzip terraform_0.12.7_linux_amd64.zip
RUN mv terraform /usr/local/bin; chmod a+rwx /usr/local/bin/terraform
CMD env | grep _ >> /etc/environment; echo "aws_access_key_id = "$ACCESS_KEY >> /home/jenkins/.aws/credentials; echo "aws_secret_access_key = "$SECRET_KEY >> /home/jenkins/.aws/credentials; service ssh start; tail -f /dev/null
