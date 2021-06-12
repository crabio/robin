package authgoogle

import (
	// External

	"context"
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	// Internal
	"github.com/iakrevetkho/robin/internal/config"
	"github.com/iakrevetkho/robin/internal/resources"
)

type Provider struct {
	oAuthConf *oauth2.Config
	// URL which should be sent to user to auth via Google
	authURL string
}

func New(config config.Config) (provider *Provider, err error) {
	secretsFilePath := filepath.Join(config.SecretsFolderPath, config.Auth.Google.SecretFileName)
	// Read JSON file with Google API secrets for your service
	data, err := ioutil.ReadFile(secretsFilePath)
	if err != nil {
		return
	}
	log.Debugf("Google secrets have been read from '%s'", secretsFilePath)
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
		authURL:   oauthConf.AuthCodeURL("state", authWithChooseAccount),
	}
	return
}

// Method for getting URL to external provider for authentication
func (p *Provider) GetAuthURL() string {
	return p.authURL
}

// Method for processing redirect from auth provider after user is authenticated
func (p *Provider) ProcessAuthRedirect(authCode string) (userProfile *resources.UserProfile, err error) {
	log.Debugf("Process Google auth redirect with code '%s'", authCode)
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
	googleUserProfile := &UserProfile{}
	err = json.Unmarshal(bodyBytes, &googleUserProfile)
	if err != nil {
		return
	}
	log.Debugf("Google user profile: %+v", googleUserProfile)

	// Convert Google user profile to overall user profile
	userProfile = &resources.UserProfile{
		FirstName: googleUserProfile.FirstName,
		LastName:  googleUserProfile.LastName,
		Email:     googleUserProfile.Email,
		Locale:    googleUserProfile.Locale,
	}

	return
}
