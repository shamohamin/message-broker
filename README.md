# message-broker
This repository represents a simple implementation of the message broker structure in Golang. 

# Run the program
```bash

go run main.go

```

## Running message broker and adding topic to it

```go
// inside main.go file.


// Creating topics
topic, err := topics.NewTopic("topic-1")
topic2, _ := topics.NewTopic("topic-2")
if err != nil {
  log.Fatalf("your requested topic doesn't exits.")
  os.Exit(1)
  return
}

// create a instance from message broker
messageBroker := mes_b.NewMessageBroker()
messageBroker.AddTopic(topic)
messageBroker.AddTopic(topic2)

```

## Adding Publisher

```go
// inside main.go file.


// Define your topic 
topic, err := topics.NewTopic("topic-1")
if err != nil {
		log.Fatalf("your requested topic doesn't exits.")
		os.Exit(1)
		return
}
// Define your data_generating method
// Add your Topic
// Execution time of publisher
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
 // Add publisher to the message broker
 messageBroker.AddPublisher(&p)

```

## Adding Subscriber

```go
// inside main.go file.


// Add your topic and put it the subscriber topic
s := sub.Subscriber{
  ID: 1,
  SubTopic: topic,
}

// Finally add the subscriber to the message broker
messageBroker.AddPublisher(&p)


```
