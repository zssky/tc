package http

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
	"fmt"
	"strings"
	"github.com/zssky/log"
)

var (
	ContentTypeTextXml = "text/xml"
	ContentTypeHtml    = "text/html; charset=utf-8"
	ContentTypeTextCss = "text/css; charset=utf-8"
	ContentTypeXJS     = "application/x-javascript"
	ContentTypeJS      = "text/javascript"
	ContentTypeJson    = "application/json; charset=utf-8"
	ContentTypeForm    = "application/x-www-form-urlencoded"
	ContentTypeImg     = "image/png"
)

// PostJSON - send an http post json Request.
func PostJSON(url, token string, data interface{}, deadline, dialTimeout time.Duration) ([]byte, int, error) {
	buf, err := json.Marshal(data)
	fmt.Println(string(buf))
	fmt.Println("---------------------------------------")
	if err != nil {
		return nil, 0, err
	}

	return Request(http.MethodPost, url, bytes.NewBuffer(buf), deadline, dialTimeout,
		map[string]string{"Content-Type": ContentTypeJson, "token": token})
}

// PostJSONWithHeader - send an http post json Request with header.
func PostJSONWithHeader(url string, header map[string]string, data interface{}, deadline, dialTimeout time.Duration) ([]byte, int, error) {
	buf, err := json.Marshal(data)
	fmt.Println(string(buf))
	fmt.Println("---------------------------------------")
	if err != nil {
		return nil, 0, err
	}

	header["Content-Type"] = ContentTypeJson

	return Request(http.MethodPost, url, bytes.NewBuffer(buf), deadline, dialTimeout, header)
}

// PostHex - send an http post json Request.
func PostHex(url, token string, data string, deadline, dialTimeout time.Duration) ([]byte, int, error) {
	return Request(http.MethodPost, url, bytes.NewBuffer([]byte(data)), deadline, dialTimeout,
		map[string]string{"Content-Type": ContentTypeJson, "token": token})
}

// Request - send an http Request
func Request(method, url string, body io.Reader, deadline, dialTimeout time.Duration, header map[string]string) ([]byte, int, error) {
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

	if header != nil {
		for key, value := range header {
			req.Header.Set(key, value)
		}
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

// PostForm post form data
func PostForm(url string, header, para map[string]string) ([]byte, int, error) {
	// write string buffer
	var r http.Request
	r.ParseForm()

	// add form key value
	for key, value := range para {
		r.Form.Add(key, value)
	}

	body := strings.TrimSpace(r.Form.Encode())
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		return nil, 0, err
	}

	// add header
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	for key, value := range header {
		req.Header.Add(key, value)
	}

	log.Debugf("header %v", req.Header)

	// request client
	resp, err := http.DefaultClient.Do(req)
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
