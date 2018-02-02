package ctrl

import "github.com/bosemian/go-webbord/api"

func NewTopicController() api.TopicController {
	return &topicCtrl{}
}

type topicCtrl struct{}

func (t *topicCtrl) Create(req *api.TopicCreateRequest) (*api.TopicCreateResponse, error) {
	return nil, nil
}

func (t *topicCtrl) Update(req *api.TopicUpdateRequest) (*api.TopicUpdateResponse, error) {
	return nil, nil
}

func (t *topicCtrl) List() (*api.TopicList, error) {
	return nil, nil
}

func (t *topicCtrl) Get(topicID int) (*api.Topic, error) {
	return nil, nil
}

func (t *topicCtrl) Delete(topicID int) error {
	return nil
}
