package post

type Post struct {
	UUID      string `json:"uuid,omitempty"`
	Title     string `json:"title,omitempty"`
	Content   string `json:"content,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"update_at,omitempty"`
}
