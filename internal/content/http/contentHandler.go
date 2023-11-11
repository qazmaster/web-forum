package http

import (
	"forum-authentication/internal/content"
	"io"
)

type contentHandler struct {
	cases *content.Cases
}

func NewContentHandler(c *content.Cases) *contentHandler {
	return &contentHandler{
		cases: c,
	}
}

func (c contentHandler) CreatePost(w io.Writer, r any) error {
	//TODO implement me
	//rw := w.(http.ResponseWriter)
	//req := r.(*http.Request)
	return nil
}

func (c contentHandler) CreateComment(w io.Writer, r any) error {
	//TODO implement me
	return nil
}

func (c contentHandler) GetPost(w io.Writer, r any) error {
	//TODO implement me
	return nil
}

func (c contentHandler) GetPostsByCategory(w io.Writer, r any) error {
	//TODO implement me
	return nil
}

func (c contentHandler) GetCreatedPostsByUser(w io.Writer, r any) error {
	//TODO implement me
	return nil
}

func (c contentHandler) GetVotedPostsByUser(w io.Writer, r any) error {
	//TODO implement me
	return nil
}
