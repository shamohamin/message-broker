package topics


import (
	"fmt"
	"errors"
)

var topicsValue  = make(map[int]string)
var topicsName   = make(map[string]int)
var defaultTopicValues = [3]int{1, 2, 3}

func init() {
	
	for i := 0; i < len(defaultTopicValues); i++ {
		topicName := fmt.Sprintf("topic-%d", defaultTopicValues[i])
		topicsValue[defaultTopicValues[i]] = topicName
		topicsName[topicName] = defaultTopicValues[i]
	}
}

type Topic struct {
	TopicName 			string
	TopicNumericValue 	int
}

func (t Topic) String() string {
	return fmt.Sprintf("topic-name(%s)", t.TopicName)
}

func NewTopic(val interface{}) (Topic, error) {
	switch val.(type) {
	case int:
		val, _ := val.(int) 
		if strVal, ok := topicsValue[val]; ok { // value does exists
			return Topic{TopicName: strVal, TopicNumericValue: int(val)}, nil
		}
	case string:
		val, _ := val.(string)
		if intVal, ok := topicsName[string(val)]; ok {
			return Topic{TopicName: string(val), TopicNumericValue: intVal}, nil
		}
	default:
		return Topic{}, errors.New("value Type must be string or integer")
	}

	return Topic{}, errors.New(fmt.Sprintf("value(%q) doesn't exist in the topicsValue", val))
}

func AddTopic(val int) (error) {
	if _, ok := topicsValue[val]; ok {
		return errors.New(fmt.Sprintf("val(%q) already exits.", val))
	}
	topicName := fmt.Sprintf("topic-%d", val)	
	topicsValue[val] = topicName
	topicsName[topicName] = val
	return nil
}