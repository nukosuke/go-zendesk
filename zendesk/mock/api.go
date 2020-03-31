package mock

import (
	"github.com/nukosuke/go-zendesk/zendesk"
)

var _ zendesk.API = (*Client)(nil)
