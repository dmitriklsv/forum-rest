package entity

type Category struct {
	ID     uint64 `json:"id,omitempty"`
	PostID uint64 `json:"post_id,omitempty"`
	Name   string `json:"name,omitempty"`
}
