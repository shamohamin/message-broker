package src_publisher


import (
	"time"

	topics "shamohamin.github/publisher_subscriber_pattern/topics"
	message_b "shamohamin.github/publisher_subscriber_pattern/src_message_broker"
)

type Publisher struct {
	ID 		 			uint32
	ExecutionDuration 	time.Duration
	InputChan 			chan <-message_b.Message
	DoneChan			chan <-struct{}
	SubTopic 			topics.Topic
}


func (pub *Publisher) Execution() {
	timeout := time.After(pub.ExecutionDuration)
	for {
		select {
		case <-timeout: // end of Subscriber Job
			pub.Done <- struct{}{}
			break
		default:
			print("break\n")
			break
		}
	}
}