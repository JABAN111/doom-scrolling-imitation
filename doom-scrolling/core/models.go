package core

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}

type Post struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	ImageURL  string `json:"image_url"`
	Caption   string `json:"caption"`
	CreatedAt string `json:"created_at"`
}
