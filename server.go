package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
	"time"
)

type Message struct {
	msg  string
	time time.Time
}

func O1average(accum_sum int64, accum_n int64, addition int64) (int64, int64, float64) {
	accum_sum += addition
	accum_n += 1
	return accum_sum, accum_n, float64(accum_sum) / float64(accum_n)
}

func reassembleFunc(in_chan chan string, words []string, out_chan chan Message, debug bool) {
	var msg = strings.Join(words, " ")
	var sentencelen = len(words)
	var i = 0
	for word := range in_chan {
		if debug {
			fmt.Println(word)
		}
		if word != words[i] {
			// Shouldn't miss the beginning of new attempt to reassemble
			if word == words[0] {
				i = 1
			} else {
				i = 0
			}
			continue
		}
		i += 1
		if i == sentencelen {
			out_chan <- Message{msg: msg, time: time.Now()}
			i = 0
		}
	}
}

func receiveFunc(in_chan chan Message, debug bool) {
	var sum, n int64

	for {
		var listenTime = time.Now().Unix()
		var m, ok = <-in_chan
		if !ok {
			break
		}

		fmt.Println(m.msg)

		var avg float64;
		sum, n, avg = O1average(sum, n, m.time.Unix() - listenTime)
		fmt.Printf("Average receive time: %f s\n", avg)
	}
}

func serverMain(client mqtt.Client, message string, debug bool) {
	var closeChan = make(chan struct{})

	var wordChan = make(chan string)
	var recvMessageChan = make(chan Message)

	var words = strings.Split(message, " ")

	for _, w := range words {
		var word = w
		var topic = fmt.Sprintf("topic_%s", word)
		if debug {
			fmt.Printf("Subscribing for topic %s\n", topic)
		}
		var token = client.Subscribe(topic, 2, func(mqtt.Client, mqtt.Message) { wordChan <- word })
		token.Wait()
		if (token.Error()) != nil {
			panic(token.Error())
		}
	}

	go reassembleFunc(wordChan, words, recvMessageChan, debug)
	go receiveFunc(recvMessageChan, debug)

	<-closeChan
}
