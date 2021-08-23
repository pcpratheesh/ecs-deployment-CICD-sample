# ecs-task-definition-ci-cd
This repository is simply update the docker image inside an aws ecs cluster task.

### Work flow
When a new update pushed into **staging**, **pre-production**, or **production** branches, the pipline configuration will trigger. It will build a new docker image and pushed to the ECR REPO. Then it will make a new task definition and update the cluster service

# Github Secrets
For working this i have added the aws credentials as github secrets. You can add the secrets from settings tab under repository. [Read more](https://docs.github.com/en/actions/reference/encrypted-secrets)

Have to configure following secret credentials
- AWS_ACCESS_KEY_ID : AWS access key
- AWS_SECRET_ACCESS_KEY : AWS secret key
- AWS_ECR_REPO : docker image container repo
- AWS_ECS_CLUSTER : cluster name

### Execution
run main.go under cmd folder to execute in local

### Build docker image
docker build -t golang-sample-app -f ./Dockerfile .
#### dev.workflow.yml
Will execute the test cases for each every push requests

## deployment.workflow.yml
From this, actual deployment works