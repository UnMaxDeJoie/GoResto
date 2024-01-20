package entities

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	PwHash     string `json:"pw_hash"`
	Email      string `json:"email"`
	Permission uint8  `json:"permissions"`
}
