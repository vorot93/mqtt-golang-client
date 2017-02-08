package main

import (
	"flag"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	var ip = flag.String("ip", "tcp://127.0.0.1:1883", "MQTT broker address")
	var server = flag.Bool("server", false, "Run in server mode")
	var debug = flag.Bool("debug", false, "Run in debug mode")
	var maxInterval = flag.Int("maxinterval", 5, "Max interval for sending words")

	flag.Parse()
	var flags = flag.Args()
	if len(flags) == 0 {
		fmt.Println("Need a message to send/receive")
		return
	}
	var message = flags[0]

	var opts = mqtt.NewClientOptions()
	opts.AddBroker(*ip)

	var client = mqtt.NewClient(opts)
	var token = client.Connect()
	token.Wait()
	if token.Error() != nil {
		panic(token.Error())
	}

	if *server {
		serverMain(client, message, *debug)
	} else {
		clientMain(client, message, *maxInterval, *debug)
	}
}
