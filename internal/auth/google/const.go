package authgoogle

import (
	// External

	"golang.org/x/oauth2"
	// Internal
)

// Constants
const (
	profileEndroint = "https://www.googleapis.com/oauth2/v2/userinfo"
)

// Semi-constants
var (
	authWithChooseAccount oauth2.AuthCodeOption = oauth2.SetAuthURLParam("prompt", "select_account")
)
