package content

import (
	"encoding/json"
	"io"
	"net/http"
)

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("web-forum mock"))
	})
	//http.HandleFunc("/register", h.User.Register)
	//http.HandleFunc("/login", h.User.Login)
	http.HandleFunc("/create-post", makeHTTPHandleFunc(h.Controller.CreatePost))
	//http.HandleFunc("/create-comment", h.Comment.CreateComment)
	//http.HandleFunc("/associate-category", h.Category.AssociateCategory)
	//http.HandleFunc("/like-dislike", h.Reaction.LikeDislike)
	return mux
}

type apiFunc func(io.Writer, any) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

type ApiError struct {
	Error string `json:"error"`
}
