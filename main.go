package main

import (
	"fmt"


	topic "shamohamin.github/publisher_subscriber_pattern/topics"
)

func main() {
	fmt.Println(topic.NewTopic("topic-1"))
}