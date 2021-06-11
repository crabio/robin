package auth

import (
	// External
	// Internal
	"github.com/iakrevetkho/robin/internal/resources"
)

type IProvider interface {
	// Method for getting URL to external provider for authentication
	GetAuthURL() string
	// Method for processing redirect from auth provider after user is authenticated
	ProcessAuthRedirect(authCode string) (userProfile *resources.UserProfile, err error)
}
