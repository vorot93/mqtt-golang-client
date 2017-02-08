# mqtt-golang-test
Test client for MQTT. Send, receive and reassemble sentences with [Eclipse Mosquitto](http://mosquitto.org) acting as the message broker.

## Usage
The simplest way to run test client and server is to use [docker-compose](https://docs.docker.com/compose). You can customize `docker-compose.yml` to your liking.

```
$ make build
$ sudo docker-compose up
```

## License
See COPYING file for the license information.
