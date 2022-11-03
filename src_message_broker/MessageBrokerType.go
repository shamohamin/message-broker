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
	TopicsVal				[]topics.Topic
	SubscribersMap			map[topics.Topic][]*subs.Subscriber
	MessageQueues			map[topics.Topic]*msg.MessageQueue
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
		MessageQueues: make(map[topics.Topic]*msg.MessageQueue),
		PublisherMessageChan: pubMessageChans,
		WorkerCount: uint32(0),
		TopicsVal: make([]topics.Topic, 0),
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
		msgB.TopicsVal = append(msgB.TopicsVal, t)
		msgB.SubscribersMap[t] =  make([]*subs.Subscriber, 0)
		msgB.MessageQueues[t]  =  msg.NewMessageQueue(1000)
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
	sub.OutputChan = make(chan string)
	sub.OverChan   = make(chan bool)
	if _, ok := msgB.SubscribersMap[sub.SubTopic]; !ok {
		log.Fatalf("topic(%q)-doesnot exist in message broker. try adding this topic.", sub.SubTopic)
		return
	}
	msgB.SubscribersMap[sub.SubTopic] = append(msgB.SubscribersMap[sub.SubTopic], sub)
	msgB.SubLock.Unlock()
	// start the subscriber
	go sub.Execution()
}


func (msgB *MessageBroker) handlingSubscribers(timeoutSubs chan bool) {
	for {
		msgB.PubLock.Lock()
		if msgB.WorkerCount <= 0 {
			msgB.PubLock.Unlock()
			// send the subscribers to sleep
			for _, subscribers := range msgB.SubscribersMap {
				for i := 0; i < len(subscribers); i++ {
					subscribers[i].OverChan <- true
				}
			}

			timeoutSubs <- true
			break
		}
		msgB.PubLock.Unlock()

		msgB.SubLock.Lock()
		msgB.MsgLock.Lock()
		for _, val := range msgB.TopicsVal {
			// getting the queue for this topic
			if msgQueue, ok := msgB.MessageQueues[val]; ok {
				// getting the val of last item added
				if	valMsg := msgQueue.Dequeue(); valMsg != nil {
					if subscribers, found := msgB.SubscribersMap[val]; found {
						for i := 0; i < len(subscribers); i++ {
							subscribers[i].OutputChan <- valMsg.Content
						}
					}
				}
			}
		}
		msgB.MsgLock.Unlock()
		msgB.SubLock.Unlock()
	}
}


func (msgB *MessageBroker) Start() {
	Over := false
	timeoutSubscribers := make(chan bool)
	go msgB.handlingSubscribers(timeoutSubscribers)

	go func() {
		for msgB.WorkerCount > 0 {
			select {
			case <- msgB.PubDoneChan:
				msgB.PubLock.Lock()
				msgB.WorkerCount -= 1
				msgB.PubLock.Unlock()
			default:
				continue
			}
		}
		Over = true
		//msgB.OverChan <- true
	}()

	for {
		for i := 0; i < len(msgB.PublisherMessageChan); i++ {
			select {
			case msg := <-msgB.PublisherMessageChan[i]:
				msgB.MsgLock.Lock()
				msgB.MessageQueues[msg.Topic].Enqueue(&msg)
				msgB.MsgLock.Unlock()
			default:
				continue
			}
		}

		if Over {
			break
		}
	}
	<-timeoutSubscribers
}

func (msgB MessageBroker) String() string {
	return fmt.Sprintf("\nID: %d\npublishers: %q\nSubscriberMap: %q\n***\n", msgB.ID, msgB.Publishers, msgB.SubscribersMap)
}


func TestMessageQueue() {
	topic, err := topics.NewTopic("topic-1")
	if err != nil {
		log.Fatalf("your requested topic doesn't exits.")
		return
	}
	mq := msg.NewMessageQueue(10)
	for i := 0; i < 11; i++ {
		mq.Enqueue(&msg.Message{
			Topic: topic,
			Content: fmt.Sprintf("%d", i),
		})
	}
	fmt.Println(mq)

	for i := 0; i < 10; i++ {
		fmt.Println(mq.Dequeue())
	}

	fmt.Println(mq)
	for i := 0; i < 10; i++ {
		mq.Enqueue(&msg.Message{
			Topic: topic,
			Content: fmt.Sprintf("%d", i),
		})
	}
	fmt.Println(mq)
}