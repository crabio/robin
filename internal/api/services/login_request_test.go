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

func TestGoogleLoginRequest(t *testing.T) {
	// Define Mock data
	googleAuthProvider := authgoogle.Mock{}

	controllerData := apiresources.ControllerData{
		GoogleAuthProvider: &googleAuthProvider,
	}

	// Define input data
	request := proto_resources.LoginRequest{
		Provider: proto_resources.AuthProviderEnum_google,
	}

	// Execute function
	response, err := LoginRequest(controllerData, &request)

	// Check result
	assert.NoError(t, err)
	assert.Equal(t, response, &proto_resources.LoginResponse{
		Url: "test google auth URL",
	})
}
