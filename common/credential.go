package common

// Credential is interface of API credential
type Credential interface {
	Email() string
	Secret() string
}

// BasicAuthCredential is type of credential for Basic authentication
type BasicAuthCredential struct {
	email    string
	password string
}

// APITokenCredential.Email is accessor which returns email address
func (c BasicAuthCredential) Email() string {
	return c.email
}

// APITokenCredential.Secret is accessor which returns password
func (c BasicAuthCredential) Secret() string {
	return c.password
}

// APITokenCredential is type of credential for API token authentication
type APITokenCredential struct {
	email    string
	apiToken string
}

// APITokenCredential.Email is accessor which returns email address
func (c APITokenCredential) Email() string {
	return c.email
}

// APITokenCredential.Secret is accessor which returns API token
func (c APITokenCredential) Secret() string {
	return c.apiToken
}
