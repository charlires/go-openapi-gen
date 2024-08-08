package post

import (
	"net/http"

	"github.com/charlires/go-openapi-gen/examples/stdlib/models"
)

type TopicPostRequest struct {
	models.Topic
}

type TopicPostResponse struct {
	OK bool `json:"ok"`
}

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
