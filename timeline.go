package main

import "fmt"

type timeline struct {
	Event string
}

func main() {

	fmt.Println("Hello")
	myTimeline := timeline{}

	myTimeline.Event = "I am a string"

	fmt.Println(myTimeline.Event)
}
