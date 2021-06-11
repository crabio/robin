package authgoogle

type UserProfile struct {
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	Email     string `json:"email"`
	Locale    string `json:"locale"`
}
