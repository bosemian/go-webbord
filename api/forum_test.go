package api_test

import (
	"testing"

	"github.com/bosemian/go-webbord/api"
)

func TestValidateForumCreateRequest(t *testing.T) {
	cases := []struct {
		req      api.ForumCreateRequest
		hasError bool
	}{
		{api.ForumCreateRequest{}, true},
		{api.ForumCreateRequest{ForumName: "a"}, true},
		{api.ForumCreateRequest{ForumName: "123"}, true},
		{api.ForumCreateRequest{ForumName: "test123"}, false},
	}

	for _, c := range cases {
		err := c.req.Validate()
		if c.hasError && err == nil {
			t.Errorf("expected has error; got nil")
		}
		if !c.hasError && err != nil {
			t.Errorf("expexted not error; got %v", err)
		}
	}

}

func TestValidateForumUpdateRequest(t *testing.T) {
	testCases := []struct {
		req      api.ForumUpdateRequest
		hasError bool
	}{
		{api.ForumUpdateRequest{}, true},
		{api.ForumUpdateRequest{
			ForumID:   -1,
			ForumName: "aaaa",
		}, true},
		{api.ForumUpdateRequest{
			ForumID:   01234,
			ForumName: "123",
		}, true},
		{api.ForumUpdateRequest{
			ForumID:   123456789,
			ForumName: "TestForumName",
		}, false},
	}

	for _, tc := range testCases {
		err := tc.req.Validate()
		if tc.hasError && err == nil {
			t.Errorf("expected has error; got nil")
		}

		if !tc.hasError && err != nil {
			t.Errorf("expected has not error; got %v", err)
		}
	}
}
