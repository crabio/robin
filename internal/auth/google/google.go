package authgoogle

import (
	// External

	"context"
	"encoding/json"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	// Internal
	"github.com/iakrevetkho/robin/internal/config"
)

// Constants
const (
	profileEndroint = "https://www.googleapis.com/oauth2/v2/userinfo"
)

// Semi-constants
var (
	authWithChooseAccount oauth2.AuthCodeOption = oauth2.SetAuthURLParam("prompt", "select_account")
)

type Provider struct {
	oAuthConf *oauth2.Config
	// URL which should be sent to user to auth via Google
	AuthURL string
}

func New(config config.Config) (provider *Provider, err error) {
	// Read JSON file with Google API secrets for your service
	data, err := ioutil.ReadFile(config.Auth.Google.SecretFileName)
	if err != nil {
		return
	}
	oauthConf, err := google.ConfigFromJSON(data)
	if err != nil {
		return
	}
	// Add scopes for requesting data
	oauthConf.Scopes = []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	}
	oauthConf.RedirectURL = config.Auth.Google.RedirectURL

	provider = &Provider{
		oAuthConf: oauthConf,
		AuthURL:   oauthConf.AuthCodeURL("state", authWithChooseAccount),
	}
	return
}

// Method for processing redirect from auth provider after user is authenticated
func (p *Provider) ProcessAuthRedirect(authCode string) (userProfile *UserProfile, err error) {
	// Handle the exchange code to initiate a transport.
	token, err := p.oAuthConf.Exchange(context.Background(), authCode)
	if err != nil {
		log.Error(err)
	}

	client := p.oAuthConf.Client(context.Background(), token)

	// Get user email from Google
	response, err := client.Get(profileEndroint)
	if err != nil {
		return
	}

	// Read response body data
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	log.Debugf("Google auth response data: %s", string(bodyBytes))

	// Parse user profile data from response
	userProfile = &UserProfile{}
	err = json.Unmarshal(bodyBytes, &userProfile)
	if err != nil {
		return
	}
	log.Debugf("Google user profile: %+v", userProfile)

	return
}
