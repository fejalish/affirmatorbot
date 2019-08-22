# Affirmator affirms you!

For those days when you just need to be told that you're a good person and, dangit, people like you.

```
@Affirmator tell me I'm beautiful
@Affirmator tell me I'm awesome incarnate
@Affirmator tell me I'm a good person and people like me
```

You could even go to the dark side...

```
@Affirmator tell me I'm a lazy, piece of...
```

# Setup

The slackbot container needs a `SLACK_BOT_TOKEN` ENV var to run. This value must be the "Bot User OAuth Access Token" from your Slack app.

# Docker Compose

## Setup

Ensure `docker` and `docker-compose` are installed locally.

Copy `development.env.sample` and save it as `development.env`. Edit the file and put the token as the `SLACK_BOT_TOKEN` value.

## Commands

To run locally via docker-compose run the following commands.

```
docker-compose build
docker-compose up
```

# Kubernetes

## Setup

Ensure `minikube` and `docker` are installed locally.

Copy `k8s/secrets.env.sample` and save it as `k8s/secrets.yaml`.

`base64` encode the slack token, like the following:

```
echo -n 'slack-token-here' | base64
```

Edit the file and put the token as the `slack_bot_token` value.

## Commands

To run locally via minikube run the following commands.

```
cd k8s
kubectl apply -f ./secrets.yaml
kubectl apply -f ./deployment.yaml
```

Note: the deployment uses a published version of the docker image in this repo. If you wish to build and test your own image edit the deployment and set the container image to the name of your local image. You will also need to update minikube to look locally for images.

```
minikube addons configure registry-creds
minikube addons enable registry-creds
eval $(minikube docker-env)
```

To cleanup run the following:

```
kubectl delete -f ./deployment.yaml
kubectl delete -f ./secrets.yaml
```

# Improvements

- dependency management
- testing
  - linting
  - post-commit hooks
  - unit tests
  - intergation tests
  - infr tests
- data store
  - persistent data storage
  - migrations
- deployment
  - build pipeline
  - deploy pipeline
- monitoring
  - alerting
- go reload for quicker dev...

