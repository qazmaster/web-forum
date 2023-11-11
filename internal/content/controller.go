package content

import (
	"forum-authentication/internal/content/http"
	"io"
	http2 "net/http"
)

type Handler struct {
	Controller
}

func (h *Handler) ServeHTTP(w http2.ResponseWriter, r *http2.Request) {
	//TODO implement me
}

func NewHTTPController(c *Cases) *Handler {
	return &Handler{
		Controller: http.NewContentHandler(c),
	}
}

type Controller interface {
	CreatePost(w io.Writer, r any) error
	CreateComment(w io.Writer, r any) error
	GetPost(w io.Writer, r any) error
	GetPostsByCategory(w io.Writer, r any) error
	GetCreatedPostsByUser(w io.Writer, r any) error
	GetVotedPostsByUser(w io.Writer, r any) error
}
