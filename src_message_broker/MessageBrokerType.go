package src_message_broker

import (
	"fmt"
	"log"
	"sync"
	"math/rand"

	pubs 		"shamohamin.github/publisher_subscriber_pattern/src_publisher"
	subs 		"shamohamin.github/publisher_subscriber_pattern/src_subscriber"
	topics 		"shamohamin.github/publisher_subscriber_pattern/topics"
	msg 		"shamohamin.github/publisher_subscriber_pattern/src_message"
)

const DEFAULT_NUM_OF_PUBLISHER_CHAN = 3

type MessageBroker struct {
	ID 						int
	PubDoneChan				chan struct{}
	Publishers				[]*pubs.Publisher
	SubscribersMap			map[topics.Topic][]*subs.Subscriber
	MessageQueues			map[topics.Topic][]*msg.Message
	PublisherMessageChan 	[]chan msg.Message
	WorkerCount 			uint32
	PubLock 				*sync.Mutex
	SubLock					*sync.Mutex
	MsgLock					*sync.Mutex
}

func NewMessageBroker() MessageBroker {
	pubMessageChans := make([]chan msg.Message, 0)
	for i := 0; i < DEFAULT_NUM_OF_PUBLISHER_CHAN; i++ {
		pubMessageChans = append(pubMessageChans, make(chan msg.Message))
	}

	return MessageBroker{
		ID: int(12345),
		PubDoneChan: make(chan struct{}),
		Publishers: make([]*pubs.Publisher, 0),
		SubscribersMap: make(map[topics.Topic][]*subs.Subscriber),
		MessageQueues: make(map[topics.Topic][]*msg.Message),
		PublisherMessageChan: pubMessageChans,
		WorkerCount: uint32(0),
		PubLock: new(sync.Mutex),
		SubLock: new(sync.Mutex),
		MsgLock: new(sync.Mutex),
	}
}

func (msgB *MessageBroker) AddTopic(t topics.Topic) {
	msgB.SubLock.Lock()
	msgB.MsgLock.Lock()
	if _, ok := msgB.SubscribersMap[t]; !ok {
		// check if topic already exits or not.
		msgB.SubscribersMap[t] =  make([]*subs.Subscriber, 0)
		msgB.MessageQueues[t]  =  make([]*msg.Message, 0)
	}
	msgB.MsgLock.Unlock()
	msgB.SubLock.Unlock()
}

func (msgB *MessageBroker) AddPublisher(pub *pubs.Publisher) {
	msgB.PubLock.Lock()
	pub.InputChan = msgB.PublisherMessageChan[rand.Intn(len(msgB.PublisherMessageChan))]
	pub.DoneChan  = msgB.PubDoneChan
	msgB.Publishers = append(msgB.Publishers, pub)
	msgB.WorkerCount += 1
	msgB.PubLock.Unlock()

	go pub.Execution() // start publisher
}

func (msgB *MessageBroker) AddSubscriber(sub *subs.Subscriber) {
	msgB.SubLock.Lock()
	if _, ok := msgB.SubscribersMap[sub.SubTopic]; !ok {
		log.Fatalf("topic(%q)-doesnot exist in message broker. try adding this topic.", sub.SubTopic)
		return
	}
	msgB.SubscribersMap[sub.SubTopic] = append(msgB.SubscribersMap[sub.SubTopic], sub)
	msgB.SubLock.Unlock()
}



func (msgB *MessageBroker) Start() {
	Over := false
	go func() {
		for msgB.WorkerCount > 0 {
			select {
			case <- msgB.PubDoneChan:
				msgB.PubLock.Lock()
				msgB.WorkerCount -= 1
				fmt.Println("DONE AFRIN")
				msgB.PubLock.Unlock()
			default:
				continue
			}
		}
		Over = true
	}()

	for {
		for i := 0; i < len(msgB.PublisherMessageChan); i++ {
			select {
			case msg := <-msgB.PublisherMessageChan[i]:
				fmt.Println(msg)
				msg.Topic
			default:
				continue
			}
		}

		if Over {
			break
		}
	}
}

func (msgB MessageBroker) String() string {
	return fmt.Sprintf("\nID: %d\npublishers: %q\nSubscriberMap: %q\n***\n", msgB.ID, msgB.Publishers, msgB.SubscribersMap)
}