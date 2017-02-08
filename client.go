package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"math/rand"
	"strings"
	"time"
)

func publishFunc(client mqtt.Client, word string, maxInterval int, debug bool) {
	var topic = fmt.Sprintf("topic_%s", word)
	for {
		var token = client.Publish(topic, 2, false, word)
		token.Wait()
		if token.Error() != nil {
			panic(token.Error())
		}
		if debug {
			fmt.Println(word)
		}
		time.Sleep(time.Duration(rand.Float64() * float64(maxInterval) * float64(time.Second)))
	}
}

func clientMain(client mqtt.Client, message string, maxInterval int, debug bool) {
	var closeChan = make(chan struct{})

	for _, w := range strings.Split(message, " ") {
		go publishFunc(client, w, maxInterval, debug)
	}

	<-closeChan
}
