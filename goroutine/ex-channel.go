//
// Pass from one goroutine to another via channel
//
// "sends" and "receives" block until the sender and receiver are ready
// no need to use any other sychronization to wait for the sending message is handled

package main

import (
	"fmt"
)

func main() {

	// create a new channel with make(chan val-type)
	messages := make(chan string)

	go func() {
		// send a value into a channel using "channel <-" syntax
		// here we send "hello world" to the "messages" channel
		// from a new goroutine
		messages <- "hello world"
	}()

	// use "<-channel" to receive a value from channel
	msg := <-messages
	fmt.Println("Received message from channel:", msg)

}
