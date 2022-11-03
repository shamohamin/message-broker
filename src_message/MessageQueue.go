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
		size: sizeQ,
	}
}


func (qu *MessageQueue) Enqueue(msg *Message) {
	if qu.index >= qu.size {
		// overiding the first element
		// first element is the oldest one
		qu.messages[0] = msg
		return
	}
	
	qu.messages[qu.index] = msg
	qu.index++ 
} 

func (qu *MessageQueue) Dequeue() *Message {
	if qu.index == 0 {
		return nil
	}
	msg := qu.messages[0]
	qu.index--
	qu.messages = qu.messages[1:]
	qu.messages = append(qu.messages, nil)
	return msg
}