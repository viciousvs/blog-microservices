package post

type Post struct {
	UUID      string `json:"uuid"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"update_at"`
}
