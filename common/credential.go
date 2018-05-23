package common

// Credential is interface of API credential
type Credential interface{}

// Basicauthcredential is type of credential for Basic authentication
type BasicAuthCredential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// APITokenCredential is type of credential for API token authentication
type APITokenCredential struct {
	Email    string `json:"email"`
	APIToken string `json:"api_token"`
}
