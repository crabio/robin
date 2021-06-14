package apiservices

import (
	// External

	"testing"

	// Internal
	apiresources "github.com/iakrevetkho/robin/internal/api/resources"
	authgoogle "github.com/iakrevetkho/robin/internal/auth/google"
	proto_resources "github.com/iakrevetkho/robin/internal/proto_resources"
	"github.com/stretchr/testify/assert"
)

func TestGoogleAuthRequest(t *testing.T) {
	// Define Mock data
	googleAuthProvider := authgoogle.Mock{}

	controllerData := apiresources.ControllerData{
		GoogleAuthProvider: &googleAuthProvider,
	}

	// Define input data
	request := proto_resources.AuthRequest{
		Provider:                proto_resources.AuthProviderEnum_google,
		AuthProviderUrlResponse: "",
	}

	// Execute function
	response, err := AuthRequest(controllerData, &request)

	// Check result
	assert.NoError(t, err)
	assert.Equal(t, response, &proto_resources.AuthResponse{
		FirstName: "Big",
		LastName:  "Bo",
		Email:     "bigbo@bigbo.com",
		Locale:    "past",
	})
}
