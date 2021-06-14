package resources

// Response that Web API server receives after auth provider redirect
type AuthProviderResponse struct {
	AuthCode string `schema:"code,required"`
}
