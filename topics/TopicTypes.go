package topics


import (
	"fmt"
	"errors"
)

var topicsValue map[int]string
var topicsName  map[string]int
const defaultTopicValues []int = {1, 2, 3}

func init() {
	for i := 0; i < len(defaultTopicValues); i++ {
		topicName = fmt.Sprintf("topic-%d", defaultTopicValues[i])
		topicsValue[defaultTopicValues[i]] = topicName
		topicsName[topicName] = defaultTopicValues[i]
	}
}

type Topic struct {
	TopicName 			string
	TopicNumericValue 	int
}

func NewTopic(val interface{}) (Topic, error) {
	switch tt := val.(type) {
	case int:
		if strVal, ok := topicsValue[val]; ok { // value does exists
			return Topic{TopicName: strVal, TopicNumericValue: val}, nil
		}
	case string:
		if intVal, ok := topicsName[val]; ok {
			return Topic{topicName: val, TopicNumericValue: intVal}
		}
	default:
		return interface{}{}, errors.New("value Type must be string or integer")
	}

	return interface{}{}, errors.New(fmt.Sprintf("value(%d) doesn't exist in the topicsValue", val))
}