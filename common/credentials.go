package common

const (
	BasicAuth AuthType = iota
	APIToken
)

type AuthType int

type Credential struct {
	Subdomain string
	AuthType  AuthType
	Email     string
	Password  string
	APIToken  string
}
