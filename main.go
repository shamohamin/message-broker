package main

import (
	"os"
	"log"	
	"time"

	topics "shamohamin.github/publisher_subscriber_pattern/topics"
	mes_b "shamohamin.github/publisher_subscriber_pattern/src_message_broker"
	pub "shamohamin.github/publisher_subscriber_pattern/src_publisher"
	sub "shamohamin.github/publisher_subscriber_pattern/src_subscriber"
	gen_data "shamohamin.github/publisher_subscriber_pattern/data_generator"
)



func main() {
	topic, err := topics.NewTopic("topic-1")
	topic2, _ := topics.NewTopic("topic-2")
	if err != nil {
		log.Fatalf("your requested topic doesn't exits.")
		os.Exit(1)
		return
	}
	messageBroker := 	mes_b.NewMessageBroker()
	messageBroker.AddTopic(topic)
	messageBroker.AddTopic(topic2)

	s := sub.Subscriber{
		ID: 1,
		SubTopic: topic,
	}

	s2 := sub.Subscriber{
		ID: 2,
		SubTopic: topic2,
	}

	s3 := sub.Subscriber{
		ID: 3,
		SubTopic: topic,
	}

	p := pub.Publisher{
		ID: 1,
		PubTopic: topic,
		ExecutionDuration: time.Second * 8,
		DataGen: gen_data.DataGenerator{
			InitialValueX: 0.0,
			Step: 1.0,
			FuncExc: gen_data.X2,
		},
	}

	p2 := pub.Publisher{
		ID: 2,
		PubTopic: topic2,
		ExecutionDuration: time.Second * 8,
		DataGen: gen_data.DataGenerator{
			InitialValueX: 0.0,
			Step: 1.0,
			FuncExc: gen_data.X3,
		},
	}
	messageBroker.AddSubscriber(&s)
	messageBroker.AddSubscriber(&s2)
	messageBroker.AddSubscriber(&s3)
	
	messageBroker.AddPublisher(&p)
	messageBroker.AddPublisher(&p2)
	log.Println(messageBroker)

	go func() {
		p3 := pub.Publisher{
			ID: 3,
			PubTopic: topic2,
			ExecutionDuration: time.Second * 6,
			DataGen: gen_data.DataGenerator{
				InitialValueX: 0.0,
				Step: 1.0,
				FuncExc: gen_data.X4,
			},
		}
	
		<-time.After(time.Second * 3)
		messageBroker.AddPublisher(&p3)
	}()

	messageBroker.Start()

	
	// done := make(chan struct{})
	// pub1 := pub.Publisher{
	// 	ID: 1,
	// 	ExecutionDuration: time.Second * 2,
	// 	PubTopic: topic,
	// 	DoneChan: done,
	// 	StartTime: time.Now(),		
	// 	IsOver: false,
	// }
	// go pub1.Execution()
	// <-done
}