package common

const (
	BasicAuth AuthType = iota
	APIToken
)

type Credentials struct {
	AuthType AuthType
	Email    string
	Password string
	APIToken string
}
