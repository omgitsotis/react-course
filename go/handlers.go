package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

func addChannel(client *Client, data interface{}) {
	var channel Channel
	mapstructure.Decode(data, &channel)
	fmt.Printf("%#+v\n", channel)
	channel.ID = "abc123"
	msg := Message{"channel add", channel}
	client.send <- msg
}
