FROM ubuntu:latest
RUN apt-get update
RUN apt-get install -y wget; wget -q https://packages.microsoft.com/config/ubuntu/19.04/packages-microsoft-prod.deb -O packages-microsoft-prod.deb
RUN dpkg -i packages-microsoft-prod.deb
RUN apt update
RUN apt-get install -y apt-transport-https ca-certificates gnupg
RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 3FA7E0328081BFF6A14DA29AA6A19B38D3D831EF

RUN echo "deb https://download.mono-project.com/repo/ubuntu stable-xenial main" | tee /etc/apt/sources.list.d/mono-official-stable.list
RUN apt-get update
RUN wget http://security.ubuntu.com/ubuntu/pool/main/libp/libpng/libpng12-0_1.2.54-1ubuntu1.1_amd64.deb
RUN dpkg -i libpng12-0_1.2.54-1ubuntu1.1_amd64.deb
RUN apt-get install -y libpng12-0 libgdiplus libmono-system-web4.0-cil libmono-system-web-services4.0-cil libmono-system-servicemodel4.0a-cil
RUN apt-get install -y libmono-system-windows-forms4.0-cil libmono-system-drawing4.0-cil msbuild mono-runtime mono-devel
RUN apt-get install -y dotnet-sdk-2.2 git curl
# RUN git clone https://github.com/fsprojects/Paket.git; mozroots --import --sync
# RUN chmod +x Paket/build.sh Paket/install.sh; Paket/build.sh; Paket/install.sh

RUN dotnet --version

RUN apt-get install -y python3-pip python3-venv; pip3 --version;

RUN pip3 install sagemaker awscli wget

RUN which aws

RUN apt-get install -y unzip

RUN wget https://releases.hashicorp.com/terraform/0.12.7/terraform_0.12.7_linux_amd64.zip; unzip terraform_0.12.7_linux_amd64.zip

RUN mkdir -p /home/temp/scripts; mv terraform /home/temp/scripts/; chmod -R a+rwx /home/temp

ENV PATH="/home/temp/scripts:${PATH}"
