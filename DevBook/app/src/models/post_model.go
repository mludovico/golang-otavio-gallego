package models

type Post struct {
	ID          uint64 `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	AuthorNick  string `json:"authorNick,omitempty"`
	AuthorID    uint64 `json:"author_id,omitempty"`
	Likes       uint64 `json:"likes"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdateAt    string `json:"updated_at,omitempty"`
	ILiked      bool   `json:"i_liked"`
}
