## Okta AWS Portal

This repository runs a Python 2.7 (Alpine) Docker image that acts as a development environment for work requiring the use of AWS accounts whose access is managed with Okta. Most notably, it contains the logic needed to retrieve a temporary session token for AWS accounts (via Okta), the AWS CLI, and Terraform 0.12.24. It additionally mounts a local directory containing Infrastructure code which can be provided as an argument. This repository is meant to be to used as a submodule.

## Toolchain
* [Oktaauth](https://github.com/ThoughtWorksInc/oktaauth)
* [Aws_role_credentials](https://github.com/ThoughtWorksInc/aws_role_credentials)

## Prerequisites
* Docker
* Okta Account
* An AppID from your Okta administrators

## Setup

### Add as Git Submodule into your Repository
Assuming that you have a code repository containing code to provision infrastructure in AWS, you can add this repository as a submodule:
`git submodule add git@github.com:kelseymok/okta-aws-portal.git`

### Build Docker image
Build the Docker image: `./okta-aws-portal/go build`

**NOTE:** You'll only need to do this once unless the Dockerfile or downstream scripts are changed

### Supply Okta Config
Run the following to set up some basic (non-secret) configuration:
```bash
./okta-aws-portal/go setup
```

You'll be prompted for your username and appID (uses Beach AppID by default):
```bash
What is your username (e.g. llamas@arefantastic.com)? meow@catnip.com
What is your appID? my-awesome-app-id
What is your Okta server? puppies.okta.com
What is your app_type? amazon_aws

```

**NOTE:** This should be re-run whenever the AppID changes or you re-clone this repo (the config file is gitignored)

### Run the Docker image
Run the Docker image: `./okta-aws-portal/go run`

This mounts the contents at your project root (assumed at `okta-aws-portal/..`) to a directory at `/app` in the container.

You will be immediately dropped into a bash environment. Run `get-token` to retrieve your AWS credentials on the Docker image. Once the the command has successfully run, your AWS credentials should be mounted in the container. From there, the AWS CLI can be used directly in the terminal or Terraform can be run in any of the directories that require it.

**NOTE:** The `get-token` command might need to be re-run from time to time in order to avoid timeouts and awkward force-unlocking from long-running `terraform apply`s.

#### Testing
* `aws s3 ls` should return a list of AWS S3 buckets
* `terraform` should return a list of possible commands

#### Running AWS commands against Docker container
You can use the Docker container's current AWS session to run AWS CLI commands directly from your host:
```bash
docker exec -it okta-aws-portal aws s3 ls
```