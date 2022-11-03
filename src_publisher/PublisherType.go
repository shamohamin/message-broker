package src_publisher


import (
	"fmt"
	"time"

	topics "shamohamin.github/publisher_subscriber_pattern/topics"
	msg 	"shamohamin.github/publisher_subscriber_pattern/src_message"
	dataGen	"shamohamin.github/publisher_subscriber_pattern/data_generator" 
)

type Publisher struct {
	ID 		 					uint32
	IsOver						bool
	ExecutionDuration 			time.Duration
	StartTime					time.Time
	EndTime						time.Time
	InputChan 					chan <-msg.Message
	DoneChan					chan <-struct{}
	PubTopic 					topics.Topic
	DataGen						dataGen.DataGenerator
	MessageBrokerCommandChan	chan <- msg.Message
}


func (pub *Publisher) Execution() {
	pub.StartTime = time.Now()
	pub.IsOver = false
	timeout := time.After(pub.ExecutionDuration)// works for certain amount of time
	for !pub.IsOver {
		select {
		case <-timeout: 
			// end of Publisher Job
			pub.IsOver = true
			pub.EndTime = time.Now()
			pub.DoneChan <- struct{}{}	
		default:
			time.Sleep(500 * time.Millisecond) // sleeping
			// generating data
			data := pub.DataGen.Generate()
			strData := fmt.Sprintf("data-publisher-ID(%d)--(%0.6f)", pub.ID, data)
			pub.InputChan <- msg.Message{
				Topic: pub.PubTopic,
				Content: strData,
			}
			continue
		}
	}
}