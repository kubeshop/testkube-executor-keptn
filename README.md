# TestKube Executor Keptn

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com) [![Go Report Card](https://goreportcard.com/badge/github.com/kubeshop/testkube-executor-keptn)](https://goreportcard.com/report/github.com/kubeshop/testkube-executor-keptn)

Discord server: https://discord.gg/8RVUapTP

DockerHub repository: https://hub.docker.com/repository/docker/kubeshop/testkube-executor-keptn

This is a Keptn Service Template written in GoLang. Follow the instructions below for writing your own Keptn integration.

### Where to start

If you don't care about the details, your first entrypoint is [eventhandlers.go](eventhandlers.go). Within this file 
 you can add implementation for pre-defined Keptn Cloud events.
 
To better understand all variants of Keptn CloudEvents, please look at the [Keptn Spec](https://github.com/keptn/spec).
 
If you want to get more insights into processing those CloudEvents or even defining your own CloudEvents in code, please 
 look into [main.go](main.go) (specifically `processKeptnCloudEvent`), [deploy/service.yaml](deploy/service.yaml),
 consult the [Keptn docs](https://keptn.sh/docs/) as well as existing [Keptn Core](https://github.com/keptn/keptn) and
 [Keptn Contrib](https://github.com/keptn-contrib/) services.

### Common tasks

* Build the binary: `go build -ldflags '-linkmode=external' -v -o keptn-service-template-go`
* Run tests: `go test -race -v ./...`
* Build the docker image: `docker build . -t keptnsandbox/keptn-service-template-go:dev` (Note: Ensure that you use the correct DockerHub account/organization)
* Run the docker image locally: `docker run --rm -it -p 8080:8080 keptnsandbox/keptn-service-template-go:dev`
* Push the docker image to DockerHub: `docker push keptnsandbox/keptn-service-template-go:dev` (Note: Ensure that you use the correct DockerHub account/organization)
* Deploy the service using `kubectl`: `kubectl apply -f deploy/`
* Delete/undeploy the service using `kubectl`: `kubectl delete -f deploy/`
* Watch the deployment using `kubectl`: `kubectl -n keptn get deployment keptn-service-template-go -o wide`
* Get logs using `kubectl`: `kubectl -n keptn logs deployment/keptn-service-template-go -f`
* Watch the deployed pods using `kubectl`: `kubectl -n keptn get pods -l run=keptn-service-template-go`
* Deploy the service using [Skaffold](https://skaffold.dev/): `skaffold run --default-repo=your-docker-registry --tail` (Note: Replace `your-docker-registry` with your DockerHub username; also make sure to adapt the image name in [skaffold.yaml](skaffold.yaml))


### Testing Cloud Events

We have dummy cloud-events in the form of [RFC 2616](https://ietf.org/rfc/rfc2616.txt) requests in the [test-events/](test-events/) directory. These can be easily executed using third party plugins such as the [Huachao Mao REST Client in VS Code](https://marketplace.visualstudio.com/items?itemName=humao.rest-client).
