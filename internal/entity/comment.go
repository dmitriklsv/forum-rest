package entity

type Comment struct {
	ID     uint64 `json:"id,omitempty"`
	UserID uint64 `json:"user_id,omitempty"`
	PostID uint64 `json:"post_id,omitempty"`
	Text   string `json:"text,omitempty"`
	Rating int64  `json:"rating"`
	// Likes    uint64 `json:"likes,omitempty"`
	// Dislikes uint64 `json:"dislikes,omitempty"`
}
