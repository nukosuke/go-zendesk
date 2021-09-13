package zendesk

// Credential is interface of API credential
type Credential interface {
	Email() string
	Secret() string
	Type() string
}

// BasicAuthCredential is type of credential for Basic authentication
type BasicAuthCredential struct {
	email    string
	password string
}

// NewBasicAuthCredential creates BasicAuthCredential and returns its pointer
func NewBasicAuthCredential(email string, password string) *BasicAuthCredential {
	return &BasicAuthCredential{
		email:    email,
		password: password,
	}
}

// Email is accessor which returns email address
func (c BasicAuthCredential) Email() string {
	return c.email
}

// Secret is accessor which returns password
func (c BasicAuthCredential) Secret() string {
	return c.password
}

// Type is accessor which returns the type
func (c BasicAuthCredential) Type() string {
	return "Basic"
}

// APITokenCredential is type of credential for API token authentication
type APITokenCredential struct {
	email    string
	apiToken string
}

// NewAPITokenCredential creates APITokenCredential and returns its pointer
func NewAPITokenCredential(email string, apiToken string) *APITokenCredential {
	return &APITokenCredential{
		email:    email,
		apiToken: apiToken,
	}
}

// Email is accessor which returns email address
func (c APITokenCredential) Email() string {
	return c.email + "/token"
}

// Secret is accessor which returns API token
func (c APITokenCredential) Secret() string {
	return c.apiToken
}

// Type is accessor which returns the type
func (c APITokenCredential) Type() string {
	return "API"
}

// BearerTokenCredential is type of credential for Bearer token authentication
type BearerTokenCredential struct {
	bearerToken string
}

// NewBearerTokenCredential creates BearerTokenCredential and returns its pointer
func NewBearerTokenCredential(bearerToken string) *BearerTokenCredential {
	return &BearerTokenCredential{
		bearerToken: bearerToken,
	}
}

// Email is accessor which returns email address
func (c BearerTokenCredential) Email() string {
	return ""
}

// Secret is accessor which returns Bearer token
func (c BearerTokenCredential) Secret() string {
	return c.bearerToken
}

// Type is accessor which returns the type
func (c BearerTokenCredential) Type() string {
	return "Bearer"
}
