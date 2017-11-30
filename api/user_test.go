package api_test

import (
	"testing"

	"github.com/bosemian/go-webbord/api"
)

func TestUserValidateSignInReqiest(t *testing.T) {
	testCases := []struct {
		req      api.UserSignInRequest
		hasError bool
	}{
		{api.UserSignInRequest{}, true},
		{api.UserSignInRequest{
			Username: "speedboy",
			Password: "",
		}, true},
		{api.UserSignInRequest{
			Username: "",
			Password: "speedboy",
		}, true},
		{api.UserSignInRequest{
			Username: "speed",
			Password: "speedboy",
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
