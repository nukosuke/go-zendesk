package mock

import (
	"github.com/chrisjoyce911/go-zendesk/zendesk"
)

var _ zendesk.API = (*Client)(nil)
