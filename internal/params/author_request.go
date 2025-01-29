package params

type AuthorRequest struct {
	UserID uint64 `json:"user_id"`
	Name   string `json:"name"`
	Bio    string `json:"bio"`
}
