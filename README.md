# zarvis-example-go

Simple http server example written in [Go](https://golang.org) and use [Skaffold](https://skaffold.dev) to build.

[![Deploy](https://zarvis.ai/api/open/button.svg)](https://zarvis.ai/api/open)

## Run locally

 1. Start [Minikube](https://kubernetes.io/docs/setup/learning-environment/minikube/)
 2. Configure kubectl current context to `minikube` (`kubectl config use-context minikube`)
 3. Get [Skaffold](https://github.com/GoogleContainerTools/skaffold) and run

    ```
    $ skaffold dev --port-forward
    ```

    The command will build, deploy (to minikube) and port-forward automatically. Browse `http://localhost:12222`.

## Deploy to Zarvis

  1. Clone this repository
  2. [Install Zarvis Github App](https://zarvis.ai/projects/settings) on your cloned repository
  3. Select project from [zarvis.ai](https://zarvis.ai), go to `Deploy` tab, and select branch to deploy.
  4. Once deployed, http endpoint will be created.
