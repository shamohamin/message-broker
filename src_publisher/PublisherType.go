package src_publisher


import (
	"time"

	topics "shamohamin.github/publisher_subscriber_pattern/topics"
	message_b "shamohamin.github/publisher_subscriber_pattern/src_message_broker"
)

type Publisher struct {
	ID 		 					uint32
	IsOver						bool
	ExecutionDuration 			time.Duration
	StartTime					time.Time
	EndTime						time.Time
	InputChan 					chan <-message_b.Message
	DoneChan					chan <-struct{}
	PubTopic 					topics.Topic
	MessageBrokerCommandChan	chan <- message_b.Message
}


func (pub *Publisher) Execution() {
	timeout := time.After(pub.ExecutionDuration)// works for certain amount of time
	for !pub.IsOver {
		select {
		case <-timeout: // end of Publisher Job
			pub.IsOver = true
			pub.EndTime = time.Now()
			pub.DoneChan <- struct{}{}	
		default:
			time.Sleep(500 * Time.NanoSecond) // sleeping
			// generating data
			continue
		}
	}
}