package auth

type payloadLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Company  int    `json:"company"`
}
