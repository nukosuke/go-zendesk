package zendesk

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
)

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

// Upload is the API response received from zendesk whenc creating attachments
type Upload struct {
	Attachment  Attachment   `json:"attachment"`
	Attachments []Attachment `json:"attachments"`
	Token       string       `json:"token"`
}

type result struct {
	body []byte
	err  error
	resp *http.Response
}

// UploadWriter is used to write a zendesk attachment
type UploadWriter interface {
	io.Writer
	Close() (Upload, error)
}

type writer struct {
	*Client
	once     sync.Once
	w        io.WriteCloser
	filename string
	token    string
	c        chan result
}

func (wr *writer) open() error {
	r, w := io.Pipe()
	wr.c = make(chan result)

	wr.w = w
	path := "/uploads.json"
	req, err := http.NewRequest(http.MethodPost, wr.baseURL.String()+path, r)
	if err != nil {
		return err
	}

	wr.prepareRequest(req)
	req.Header.Set("Content-Type", "application/binary")

	q := req.URL.Query()
	if wr.token != "" {
		q.Add("token", wr.token)
	}

	q.Add("filename", wr.filename)
	req.URL.RawQuery = q.Encode()

	go func() {
		resp, err := wr.httpClient.Do(req)
		if err != nil {
			wr.c <- result{
				err: err,
			}
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			wr.c <- result{
				err: err,
			}
			return
		}

		wr.c <- result{
			body: body,
			resp: resp,
		}
	}()

	return nil
}

func (wr *writer) Write(p []byte) (n int, err error) {
	wr.once.Do(func() {
		err = wr.open()
	})

	if err != nil {
		return 0, err
	}

	return wr.w.Write(p)
}

func (wr *writer) Close() (Upload, error) {
	defer close(wr.c)
	err := wr.w.Close()
	if err != nil {
		return Upload{}, err
	}

	result := <-wr.c
	if result.err != nil {
		return Upload{}, result.err
	}

	resp, body := result.resp, result.body
	if resp.StatusCode != http.StatusCreated {
		return Upload{}, Error{
			resp: resp,
			body: body,
		}
	}

	var data struct {
		Upload Upload `json:"upload"`
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return Upload{}, err
	}

	return data.Upload, nil
}

// AttachmentAPI an interface containing all of the attachment related zendesk methods
type AttachmentAPI interface {
	UploadAttachment(filename string, token string) UploadWriter
}

// UploadAttachment returns a writer that can be used to create a zendesk attachment
// ref: https://developer.zendesk.com/rest_api/docs/support/attachments#upload-files
func (z *Client) UploadAttachment(filename string, token string) UploadWriter {
	return &writer{
		Client:   z,
		filename: filename,
		token:    token,
	}
}
