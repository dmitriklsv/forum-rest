package config

import "time"

const (
	DefaultTimeout = 5 * time.Second
)

type ctx string

var (
	UserID     ctx = "user_ID"
	PostID     ctx = "post_ID"
	CategoryID ctx = "category_ID"
	Categories ctx = "categories"
	Filter     ctx = "filter"
)
