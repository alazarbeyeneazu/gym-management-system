package models

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	CreatedAt   string `json:"created_at"`
	Password    string `json:"password"`
	State       int64  `json:"state"`
}
type RestResponse struct {
	Error  string `json:"error"`
	Status string `json:"stats"`
	User   User   `json:"user"`
}
