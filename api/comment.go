package api

import (
	"fmt"
	"time"
)

type CommentController interface {
	Create(*CommentCreateRequest) (*CommentCreateResponse, error)
	List() (*CommentList, error)
	Update(*CommentUpdateRequest, error) (*CommentUpdateResponse, error)
	Delete(int) error
}

type CommentCreateRequest struct {
	Token         string
	CommentDetail string
}

// Validate check comment request
func (c *CommentCreateRequest) Validate() error {
	if len(c.Token) == 0 {
		return fmt.Errorf("required token")
	}
	if len(c.CommentDetail) == 0 {
		return fmt.Errorf("Comment detail is empty")
	}
	return nil
}

type CommentCreateResponse struct {
	Comment
}

type CommentList struct {
	Comments []*Comment
}

type CommentUpdateRequest struct {
	CommentCreateRequest
}

// Validate CommentUpdateRequest
func (c *CommentUpdateRequest) Validate() error {
	if len(c.Token) == 0 {
		return fmt.Errorf("required token")
	}
	if len(c.CommentDetail) == 0 {
		return fmt.Errorf("Comment detail is empty")
	}
	return nil
}

type CommentUpdateResponse struct {
}

type Comment struct {
	CommentID     int
	TopicID       int
	CommentDetail string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
