package src_subscriber

import (
	"fmt"
	"time"

	msg 		"shamohamin.github/publisher_subscriber_pattern/src_message"
	topics 		"shamohamin.github/publisher_subscriber_pattern/topics"
)

const MAX_LEN_OF_MESSAGE_KEEP = 1000

type Subscriber struct {
	ID							int
	OverChan 					chan bool
	ExecutionDuration 			time.Duration
	StartTime 					time.Time
	EndTime 					time.Time
	OutputChan 					chan string
	Messages					[]string
	SubTopic 					topics.Topic
	MessageBrokerCommandChan	chan <- msg.Message
}


func (sub *Subscriber) Execution() {
	sub.StartTime = time.Now()
	sub.Messages = make([]string, 0)

	for {
		select {
		case <-sub.OverChan:
			sub.SetExecuationTime()
			fmt.Println("AFRIN DONE --> ", sub.ID)
			return
		case mes := <-sub.OutputChan:
			// keeping the messages
			fmt.Printf("SubID(%d); msg(%s); topic(%q); \n", sub.ID, mes, sub.SubTopic)
			sub.AppendMessage(mes)
		default:
			continue
		}
	}
}

func (sub *Subscriber) AppendMessage(message string) {
	if len(sub.Messages) >= MAX_LEN_OF_MESSAGE_KEEP {
		// remove the first item
		a := make([]string, 0)
		a = append(a, sub.Messages[1:]...)
		a = append(a, message)
		sub.Messages = a
	}else {
		sub.Messages = append(sub.Messages, message)
	}
}


func (sub *Subscriber) SetExecuationTime() {
	sub.EndTime = time.Now()
	sub.ExecutionDuration = sub.EndTime.Sub(sub.StartTime)
}