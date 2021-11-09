# golang-rss-service
A GoLang service utilizing the golang-rss-reader-package. This service serves as a proxy to parse URL via API calls. 

## Usage

### No Docker
In order to run it locally you need to:
`go get .` +`go build .` + `go run .`. These commands will build the app and run it on `http://localhost:8080`. The required route is `/v1/feeds` (being a POST request).

### Docker
There is also a Dockerfile that you can build and run the app. To do so follow these steps: <br/>
`docker build . -t <image_name>` <br/>
`docker run -p 8080:8080 <image_name>` <br/>
This will run the app on `http://localhost:8080`. You can change the ports by changing the first 8080 to a port of your liking.

## Validation
The service uses a validation framework, so that correct responses are sent. The only rule here is that the urls parameter in the POST is required.

## Logging
The service creates a logger.log file to log any errors during execution.