# golang-rss-service
A GoLang service utilizing the golang-rss-reader-package. This service serves as a proxy to parse URL via API calls. 

## Usage
In order to run it locally you need to:
`go get .` +`go build .` + `go run .`. These commands will build the app and run it on `http://localhost:8080`. The required route is `/v1/feeds` (being a POST request).

## Validation
The service uses a validation framework, so that correct responses are sent. The only rule here is that the urls parameter in the POST is required.

## Logging
The service creates a logger.log file to log any errors during execution.