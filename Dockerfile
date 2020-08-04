FROM python:2-alpine
WORKDIR /



RUN apk add --no-cache \
  bash \
  curl \
  git \
  jq \
  openssh \
  make

## Install Pip Packages
RUN pip install --upgrade pip awscli aws_role_credentials oktaauth

## Install Terraform
ENV TERRAFORM_VERSION=0.12.24

RUN curl -L -o ./terraform.zip \
    https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
    unzip -d /usr/local/bin ./terraform.zip && \
    rm terraform.zip

## Install Kubectl
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl && \
  chmod +x ./kubectl && \
  mv ./kubectl /bin/kubectl

RUN mkdir ~/.aws && touch ~/.aws/credentials


COPY ./get-token.sh /bin/get-token

ENTRYPOINT bash
