package src_message_broker

import (
	topics "shamohamin.github/publisher_subscriber_pattern/topics"
)

type Message struct {
	Topic 		topics.Topic
	Content		string
}