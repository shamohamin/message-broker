package src_message


type MessageQueue struct {
	messages []*Message
	index	 int
	size	 int
}

func NewMessageQueue(sizeQ int) *MessageQueue {
	return &MessageQueue{
		messages: make([]*Message, sizeQ),
		index: 0,
		size: sizeQ
	}
}


func (qu *MessageQueue) Enqueue(msg *Message) {
	if index >= qu.size - 1 {
		// removing the last element
		qu.messages = qu.messages[:qu.index]
	}
	qu.index++ 
	qu.messages[qu.index] = msg
} 

func (qu *MessageQueue) Dequeue() *Message {
	msg := qu.messages[0]
	msg.index--
	qu.messages = qu.messages[1:]
	return msg
}