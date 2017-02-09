package http

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	ContentTypeTextXml = []byte("text/xml")
	ContentTypeHtml    = []byte("text/html; charset=utf-8")
	ContentTypeTextCss = []byte("text/css; charset=utf-8")
	ContentTypeXJS     = []byte("application/x-javascript")
	ContentTypeJS      = []byte("text/javascript")
	ContentTypeJson    = []byte("application/json; charset=utf-8")
	ContentTypeForm    = []byte("application/x-www-form-urlencoded")
	ContentTypeImg     = []byte("image/png")
)

// Request - send an http Request
func Request(method, url string, body io.Reader, deadline, dialTimeout time.Duration, contentType []byte) ([]byte, int, error) {
	client := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(deadline)
				c, err := net.DialTimeout(netw, addr, dialTimeout)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, 0, err
	}

	if contentType != nil {
		req.Header.Set("Content-Type", contentType)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	return data, resp.StatusCode, nil
}

// HttpResponse - htt Response
type HttpResponse struct {
	Code    int                    `json:"Code"`
	Message string                 `json:"Message"`
	Data    map[string]interface{} `json:"Data,omitempty"`
}

// NewHttpResponse -
func NewHttpResponse() *HttpResponse {
	return &HttpResponse{
		Code:    0,
		Message: "success",
		Data:    make(map[string]interface{}),
	}
}

// Response - write data to resp
func (h *HttpResponse) Response(resp http.ResponseWriter) {
	resp.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(h)
	resp.Write(data)
}

// ResponseWithErr - write data to resp with error
func (h *HttpResponse) ResponseWithErr(resp http.ResponseWriter, err error) {
	resp.WriteHeader(http.StatusOK)
	if err != nil {
		h.Error(err)
	}

	data, _ := json.Marshal(h)
	resp.Write(data)
}

// Error - set Error
func (h *HttpResponse) Error(err error) {
	h.Code = 1
	h.Message = err.Error()
}
