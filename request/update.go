package request

type UpdatePost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
