package authgoogle

import (
	// External

	"github.com/iakrevetkho/robin/internal/resources"
)

type Mock struct {
}

// Method for getting URL to external provider for authentication
func (p *Mock) GetAuthURL() string {
	return "test google auth URL"
}

// Method for processing redirect from auth provider after user is authenticated
func (p *Mock) ProcessAuthRedirect(authCode string) (userProfile *resources.UserProfile, err error) {
	return &resources.UserProfile{
		FirstName: "Big",
		LastName:  "Bo",
		Email:     "bigbo@bigbo.com",
		Locale:    "past",
	}, nil
}
