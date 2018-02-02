package api

import (
	"fmt"
	"time"
	"unicode/utf8"
)

type ForumController interface {
	Create(*ForumCreateRequest) (*ForumCreateResponse, error)
	Update(*ForumUpdateRequest) (*ForumUpdateResponse, error)
	List() (*ForumListResponse, error)
	Get(int) (*Forum, error)
	Delete(int) error
}

// ForumCreateRequest is a struct
// for ForumController
type ForumCreateRequest struct {
	ForumName string
}

// Validate validate ForumCreateRequest
func (req *ForumCreateRequest) Validate() error {
	if len(req.ForumName) == 0 {
		return fmt.Errorf("Forum name required")
	}
	if utf8.RuneCountInString(req.ForumName) < 4 {
		return fmt.Errorf("Forum name must be has 4 char")
	}
	return nil
}

// ForumCreateResponse is a struct
// for ForumController
type ForumCreateResponse struct {
	Forum *Forum
}

// ForumUpdateRequest is a struct
// for ForumController
type ForumUpdateRequest struct {
	ForumID   int
	ForumName string
}

// Validate validate ForumUpdateRequest
func (req *ForumUpdateRequest) Validate() error {
	if len(req.ForumName) == 0 && req.ForumID != 0 {
		return fmt.Errorf("id and forumname has required")
	}

	if req.ForumID <= 0 || utf8.RuneCountInString(req.ForumName) < 4 {
		return fmt.Errorf("id must be has a 0 and forumname must be 4 char")
	}
	return nil
}

type ForumUpdateResponse struct {
}

// ForumListResponse for list
type ForumListResponse struct {
	Forums []*Forum
}

// struct Forum
type Forum struct {
	ForumID   int
	ForumName string
	CreatedAt time.Time
	UpdatedAt time.Time
}
