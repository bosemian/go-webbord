package ctrl

import "github.com/bosemian/go-webbord/api"

func NewForumController() api.ForumController {
	return &forumCtrl{}
}

type forumCtrl struct{}

func (f *forumCtrl) Create(req *api.ForumCreateRequest) (*api.ForumCreateResponse, error) {
	return nil, nil
}

func (f *forumCtrl) Update(req *api.ForumUpdateRequest) (*api.ForumUpdateResponse, error) {
	return nil, nil
}

func (f *forumCtrl) List(req *api.ForumListRequest) (*api.ForumListResponse, error) {
	return nil, nil
}

func (f *forumCtrl) Get(forumID int) (*api.Forum, error) {
	return nil, nil
}

func (f *forumCtrl) Delete(forumID int) error {
	return nil
}
