package request

import "encoding/json"

type UpdatePost struct {
	ID      json.Number `json:"id"`
	Title   string      `json:"title"`
	Content string      `json:"content"`
}
