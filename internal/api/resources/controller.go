package resources

import (
	// External
	// Internal
	auth_google "github.com/iakrevetkho/robin/internal/auth/google"
)

type ControllerData struct {
	GoogleAuthProvider *auth_google.Provider
}
