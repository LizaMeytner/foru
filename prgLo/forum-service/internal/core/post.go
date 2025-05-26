package core

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ForumCore struct {
	posts map[string]Post
}

func NewForumCore() *ForumCore {
	return &ForumCore{
		posts: map[string]Post{
			"1": {ID: "1", Title: "Welcome", Content: "First post!"},
		},
	}
}

func (c *ForumCore) CreatePost(post Post) error {
	c.posts[post.ID] = post
	return nil
}
