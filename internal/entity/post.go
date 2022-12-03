package entity

type Post struct {
	ID         uint64   `json:"id,omitempty"`
	UserID     uint64   `json:"user_id,omitempty"`
	Title      string   `json:"title,omitempty"`
	Text       string   `json:"text,omitempty"`
	Categories []string `json:"categories,omitempty"`
	// Likes      uint64   `json:"likes,omitempty"`
	// Dislikes uint64   `json:"dislikes,omitempty"`
}
