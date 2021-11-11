# golang-rss-service
A GoLang service utilizing the golang-rss-reader-package. This service serves as a proxy to parse URL via API calls. The services utilizes JWT tokens to secure the incoming requests. The user is made in memory for the time-being. Future improvements may include the use of a database storage. 

## Usage

### No Docker
In order to run it locally you need to:
`go get .` +`go build .` + `go run .`. These commands will build the app and run it on `http://localhost:8080`. The required route is `/v1/feeds` (being a POST request).

### Docker
There is also a Dockerfile that you can build and run the app. To do so follow these steps: <br/>
`docker build . -t <image_name>` <br/>
`docker run --env-file .env -p 3000:3000 <image_name> ` <br/>
This will run the app on `http://localhost:3000`. You can change the ports by changing the first 8080 to a port of your liking.

### JWT
In order to access the feeds API, you first need to login. For testing purposes the test user is as follows: <br>
`username: emerchantpay password: password` <br>
After logging in an access token will be provided which needs to be added to the as a Bearer Token in the header of the feed request. JWT tokens have a lifespan of 15 minutes. Refresh tokens are introduced as well they last for 7 days.

## Validation
The service uses a validation framework, so that correct responses are sent. The only rule here is that the urls parameter in the POST is required.

## Logging
The service creates a logger.log file to log any errors during execution.