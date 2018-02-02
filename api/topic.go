package api

import (
	"fmt"
	"time"
	"unicode/utf8"
)

type TopicController interface {
	Create(*TopicCreateRequest) (*TopicCreateResponse, error)
	Update(*TopicUpdateRequest) (*TopicUpdateResponse, error)
	List() (*TopicList, error)
	Get(int) (*Topic, error)
	Delete(int) error
}

type TopicCreateRequest struct {
	Token     string
	ForumID   int
	TopicName string
}

// Validate TopicCreateRequest
func (t *TopicCreateRequest) Validate() error {
	if len(t.Token) == 0 {
		return fmt.Errorf("token is empty")
	}
	if t.ForumID <= 0 {
		return fmt.Errorf("forumid empty")
	}
	if len(t.TopicName) == 0 {
		return fmt.Errorf("topicname is empty")
	}
	if utf8.RuneCountInString(t.TopicName) < 6 {
		return fmt.Errorf("topicname required at least 6 character")
	}
	return nil
}

type TopicCreateResponse struct {
	*Topic
}

type TopicUpdateRequest struct {
	Token        string
	ForumID      int
	NewTopicName string
}

// Validate test evry request update Topic
func (req *TopicUpdateRequest) Validate() error {
	if req.ForumID <= 0 {
		return fmt.Errorf("forumid is empty")
	}
	if len(req.Token) == 0 {
		return fmt.Errorf("token is empty")
	}
	if len(req.NewTopicName) == 0 {
		return fmt.Errorf("newtopicname is empty")
	}
	if utf8.RuneCountInString(req.NewTopicName) < 6 {
		return fmt.Errorf("newtopicname required at least 6 character")
	}
	return nil
}

type TopicUpdateResponse struct {
	*Topic
}

type TopicList struct {
	Topics []*Topic
}

type Topic struct {
	TopicID   int
	ForumID   int
	TopicName string
	CreatedAt time.Time
	UpdatedAt time.Time
}
