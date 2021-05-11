# ecs-task-definition-ci-cd
This repository is simply update the docker image inside an aws ecs cluster task. For this have just added a simple golang web app.

# Github Secrets
For working this i have added the aws credentials as github secrets. You can add the secrets from settings tab under repository. [Read more](https://docs.github.com/en/actions/reference/encrypted-secrets)

Create new variables as follow
- AWS_ACCESS_KEY_ID : access key
- AWS_SECRET_ACCESS_KEY : secret key
- AWS_ECR_REPO  : ECR repository name
- AWS_ECS_CLUSTER  : ECS cluster name

# Execution
run main.go under cmd folder to execute in local

# Build docker image
docker build -t golang-sample-app -f ./Dockerfile .

# Pipline 
Added pipeline files into .github folder.

## dev.workflow.yml
Will execute the test cases for each every push requests

## deployment.workflow.yml
From this, actual deployment works