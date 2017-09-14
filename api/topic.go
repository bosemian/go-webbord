package api

import (
	"time"
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

type TopicCreateResponse struct {
	*Topic
}

type TopicUpdateRequest struct {
	*Topic
}

type TopicUpdateResponse struct {
	*Topic
}

type TopicList struct {
	Topics []*Topic
}

type Topic struct {
	TopicId   int
	ForumId   int
	TopicName string
	CreatedAt time.Time
	UpdatedAt time.Time
}
