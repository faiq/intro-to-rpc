package handler

import (
	"github.com/faiq/intro-to-rpc/gen-go/service"
	"github.com/faiq/intro-to-rpc/server/request"
)

type TagsHandler struct {
	accessToken string
}

func NewTagsHandler(token string) *TagsHandler {
	return &TagsHandler{
		accessToken: token,
	}
}

func (t *TagsHandler) Generate(img service.Image) ([]string, error) {
	return request.GetImageTags([]byte(img), t.accessToken)
}
