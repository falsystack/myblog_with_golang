package response

type PostResponse struct {
	id      uint   `json:"id" binding:"required"`
	title   string `json:"title" binding:"required"`
	content string `json:"content" binding:"required"`
}

func NewPostResponse(id uint, title, content string) *PostResponse {
	return &PostResponse{
		id:      id,
		title:   title,
		content: content,
	}
}
