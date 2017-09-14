package api

import (
	"time"
)

type CommentController interface {
	Create(*CommentCreateRequest) (*CommentCreateResponse, error)
	List(int) (*CommentList, error)
	Update(int) (*Comment, error)
	Delete(int) (error)
}

type CommentCreateRequest struct {
	Token string
	CommentDetail string
}

type CommentCreateResponse struct {

}

type CommentList struct {
	Comments []*Comment
}

type Comment struct {
	CommentID     int
	TopicID       int
	CommentDetail string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
