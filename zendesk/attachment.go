package zendesk

// Attachment is struct for attachment payload
// https://developer.zendesk.com/rest_api/docs/support/attachments.html
type Attachment struct {
	ID          int64   `json:"id"`
	FileName    string  `json:"file_name"`
	ContentURL  string  `json:"content_url"`
	ContentType string  `json:"content_type"`
	Size        int64   `json:"size"`
	Thumbnails  []Photo `json:"thumbnails"`
	Inline      bool    `json:"inline"`
}

// Photo is thumbnail which is included in attachment
type Photo struct {
	ID          int64  `json:"id"`
	FileName    string `json:"file_name"`
	ContentURL  string `json:"content_url"`
	ContentType string `json:"content_type"`
	Size        int64  `json:"size"`
}
