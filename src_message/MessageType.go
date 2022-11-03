package src_message

import (
	"fmt"
	
	topics "shamohamin.github/publisher_subscriber_pattern/topics"
)

type Message struct {
	Topic 		topics.Topic
	Content		string
}

func (msg Message) String() string {
	return fmt.Sprintf("{{%q, Content: %s}}", msg.Topic, msg.Content)
}