package api_test

import (
	"testing"

	"github.com/bosemian/go-webbord/api"
)

func TestValidateCreateComment(t *testing.T) {
	req := api.CommentCreateRequest{}
	err := req.Validate()
	if err == nil {
		t.Errorf("expected error; but got nil")
	}

	req = api.CommentCreateRequest{
		Token: "token",
	}
	err = req.Validate()
	if err == nil {
		t.Errorf("expected error; but got nil")
	}

	req = api.CommentCreateRequest{
		Token:         "token",
		CommentDetail: "HelloWorld",
	}
	err = req.Validate()
	if err != nil {
		t.Errorf("expected nil; but got %v", err)
	}
}

func TestValidateUpdateComment(t *testing.T) {
	req := api.CommentUpdateRequest{}
	err := req.Validate()
	if err == nil {
		t.Errorf("expected error; but got nil")
	}

	req = api.CommentUpdateRequest{
		api.CommentCreateRequest{
			Token: "token",
		},
	}

	err = req.Validate()
	if err == nil {
		t.Errorf("expected error; but got nil")
	}

	req = api.CommentUpdateRequest{
		api.CommentCreateRequest{
			Token:         "token",
			CommentDetail: "Updated Comment",
		},
	}

	err = req.Validate()
	if err != nil {
		t.Errorf("expected nil; but got %v", err)
	}

}
