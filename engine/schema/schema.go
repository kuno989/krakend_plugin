package schema

type Key struct {
	UserId    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	ExpiresAt string `json:"expires_at"`
	ApiKey    string `json:"api_key"`

}
