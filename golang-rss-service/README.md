# GoLang RSS service
A GoLang service utilizing the golang-rss-reader-package. 

# Specification
This service should serve as a proxy to parse URL via API calls. The parsing is done by utilizing the golang-rss-reader-package. An .env file is provided with env variables.
1. The service should utilize JSON API calls and accept URL arrays and parse their RSS feeds.
2. The returning response should be in the form:
```
{
    "items": [
        {
            "title": string,
            "link": string,
            "source": string,
            "source_url": string,
            "publish_date": string,
            "description": text
        }
    ]
}
```

## Installation

Inside the .env file you need to add some variables (these are set-up for docker compose use):

```
ACCESS_SECRET=jdnfksdmfksd
REFRESH_SECRET=mcmvmkmsdnfsdmfdsjf
REDIS_DSN=localhost:6379
QUEUE_NAME="rss"
AMQP_TRANSPORT="amqp://guest:guest@rabbitmq:5672/"
LARAVEL_APP_URL=http://laravel.test:80
```

### Docker (preferred)

***You will need docker and docker-compose installed in order to continue with this step***

There is also a Dockerfile + docker-compose.yaml that you can build and run the app. To do so follow these steps: <br/>
- `docker compose build .`
- `docker compose up`

This will run the app on `http://localhost:3000`. <br>

#### Notes on docker
As RabbitMQ has been implemented, the docker compose has a rabbitmq image (it has a redis image as well more on that in the [JWT section](#JWT)). The RabbitMQ image depends on the laravel-sail RSS app network, so it is important to first build the Laravel APP and then the GoLang one. The service here acts as the consumer, where the laravel-app is the producer.

### Local

***If you run it locally, you will need `redis` installed, as well as `rabbitMQ`.***

In order to run it locally you need to:
`go get .` +`go build .` + `go run .`. These commands will build the app and run it on `http://localhost:3000`. The required route is `/v1/feeds` (being a POST request). You will need to change the following .env file variables:
```
AMQP_TRANSPORT="amqp://guest:guest@rabbitmq:5672/" -> "amqp://guest:guest@localhost:5672/" (if rabbitMQ is ran locally)
LARAVEL_APP_URL=http://laravel.test:80 (http://localhost:80 if laravel-app is ran locally)
```
## Notes on the service
### JWT
JWT tokens have been introduced in order to introduce a basic level of security (in this case using an in-memory user). In order to access the feeds API, you first need to login. For testing purposes the test user is as follows: <br>
`username: emerchantpay password: password` <br>
After logging in an access token will be provided which needs to be added to the as a Bearer Token in the authorization header of the feed request. JWT tokens have a lifespan of 15 minutes. Refresh tokens are introduced as well they last for 7 days. These tokens are stored in redis as key-value pairs (this feature is more to show that I can integrate redis in GoLang). If you are using docker compose, redis is added as an image.

#### JWT Notes
As the communication between the laravel app and the golang app is through JSON API calls, the refresh token functionality is not utilized. It is added for future developments where a GUI is added to remove the need of re-using the /login method, by refreshing the access token. Logout method is added as well to remove the JWT acces token as soon as the user stops using them, instead of waiting for their expiration time  (improved security).

## Validation
The service uses a validation framework, so that correct responses are sent. The only rule here is that the urls parameter in the POST is required.

## Logging
The service creates a logger.log file to log any errors during execution. If docker compose is used the file will be inside the container. 

## Contact

For any questions feel free to contact @ ***bkolev95@gmail.com***