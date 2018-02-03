package common

const (
	BasicAuth AuthType = iota
	APIToken
)

type AuthType int

type Credential struct {
	AuthType  AuthType
	Email     string
	Password  string
	APIToken  string
}
