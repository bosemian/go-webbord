package api_test

import (
	"testing"

	"github.com/bosemian/go-webbord/api"
)

func TestValidateCreateTopicRequest(t *testing.T) {
	cases := []struct {
		req      api.TopicCreateRequest
		hasError bool
	}{
		{api.TopicCreateRequest{}, true},
		{api.TopicCreateRequest{
			Token: "token",
		}, true},
		{api.TopicCreateRequest{
			Token:   "token",
			ForumID: 999,
		}, true},
		{api.TopicCreateRequest{
			Token:     "token",
			ForumID:   999,
			TopicName: "",
		}, true},
		{api.TopicCreateRequest{
			Token:     "token",
			ForumID:   999,
			TopicName: "Hi",
		}, true},
		{api.TopicCreateRequest{
			Token:     "token",
			ForumID:   999,
			TopicName: "I'm must test before write a real code",
		}, false},
	}

	for _, c := range cases {
		err := c.req.Validate()
		if c.hasError && err == nil {
			t.Errorf("expected has error; but got nil")
		}
		if !c.hasError && err != nil {
			t.Errorf("unexpected has error; but got %v", err)
		}
	}
}

func TestValidateUpdateTopicRequest(t *testing.T) {
	testCases := []struct {
		req      api.TopicUpdateRequest
		hasError bool
	}{
		{api.TopicUpdateRequest{}, true},
		{api.TopicUpdateRequest{
			ForumID: 999,
		}, true},
		{api.TopicUpdateRequest{
			ForumID: 999,
			Token:   "token",
		}, true},
		{api.TopicUpdateRequest{
			ForumID:      999,
			Token:        "token",
			NewTopicName: "Hi",
		}, true},
		{api.TopicUpdateRequest{
			ForumID:      999,
			Token:        "token",
			NewTopicName: "I have to write test case and run it",
		}, false},
	}

	for _, c := range testCases {
		err := c.req.Validate()
		if c.hasError && err == nil {
			t.Errorf("expected has error; but got nil")
		}
		if !c.hasError && err != nil {
			t.Errorf("unexpected has error; but got %v", err)
		}
	}
}
