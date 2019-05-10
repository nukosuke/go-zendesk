package mock

import (
	"github.com/nukosuke/go-zendesk/v0.3/zendesk"
)

var _ zendesk.API = (*Client)(nil)
