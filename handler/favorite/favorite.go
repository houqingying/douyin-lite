package favorite

type Video struct {
	author_id int64 `json:"id,omitempty"`
	// Author         repository.User `json:"author"`
	play_url       string `json:"play_url" json:"play_url,omitempty"`
	cover_url      string `json:"cover_url,omitempty"`
	favorite_count int64  `json:"favorite_count,omitempty"`
	comment_count  int64  `json:"comment_count,omitempty"`
	// status     bool   `json:"status,omitempty"`
	title string `json:"title,omitempty"`
}
