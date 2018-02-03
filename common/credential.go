package common

const (
	// BasicAuth is type of Basic authentication
	BasicAuth AuthType = iota
	// APIToken is type of API access token
	APIToken
)

// AuthType is enum of API authentication type
type AuthType int

// Credential has data to authenticate user for API
type Credential struct {
	AuthType  AuthType
	Email     string
	Password  string
	APIToken  string
}
