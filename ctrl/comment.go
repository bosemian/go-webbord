package ctrl

import "github.com/bosemian/go-webbord/api"

func NewCommentController() api.CommentController {
	return &commentCtrl{}
}

type commentCtrl struct{}

func (c *commentCtrl) Create(req *api.CommentCreateRequest) (*api.CommentCreateResponse, error) {
	return nil, nil
}

func (c *commentCtrl) Update(req *api.CommentUpdateRequest) (*api.CommentUpdateResponse, error) {
	return nil, nil
}

func (c *commentCtrl) List() (*api.CommentList, error) {
	return nil, nil
}

func (c *commentCtrl) Get(commentID int) (*api.Comment, error) {
	return nil, nil
}

func (c *commentCtrl) Delete(commentID int) error {
	return nil
}
