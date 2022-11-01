package main

import (
	"os"
	"log"	
	"time"

	topics "shamohamin.github/publisher_subscriber_pattern/topics"
	pub "shamohamin.github/publisher_subscriber_pattern/src_publisher"
)

func main() {
	topic, err := topics.NewTopic("topic-1")
	if err != nil {
		log.Fatalf("your requested topic doesn't exits.")
		os.Exit(1)
		return
	}
	
	done := make(chan struct{})
	pub1 := pub.Publisher{
		ID: 1,
		ExecutionDuration: time.Second * 2,
		PubTopic: topic,
		DoneChan: done,
		StartTime: time.Now(),		
		IsOver: false,
	}
	go pub1.Execution()
	<-done
}