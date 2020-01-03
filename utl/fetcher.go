// Package utl (fetcher.go) :
// These methods are for retrieving data from URL.
package utl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

// RequestParams : Parameters for FetchAPI
type RequestParams struct {
	Method        string
	APIURL        string
	Data          io.Reader
	Contenttype   string
	ContentLength string
	ContentRange  string
	Accesstoken   string
	Dtime         int64
}

// errHandlingFromFetch : Add error messages to Msgar.
func (p *FileInf) errHandlingFromFetch(body []byte) {
	var em map[string]interface{}
	json.Unmarshal(body, &em)
	erMsgBase1 := em["error"].(map[string]interface{})
	erMsgBase2 := erMsgBase1["errors"].([]interface{})[0].(map[string]interface{})
	erCode := erMsgBase1["code"].(float64)
	erLoc := erMsgBase2["location"].(string)
	erMsg := erMsgBase2["message"].(string)
	p.Msgar = append(p.Msgar, fmt.Sprintf("Status code is %d, location is '%s', Error message is '%s'.", int(erCode), erLoc, erMsg))
}

// FetchAPI : For fetching data to URL.
func (r *RequestParams) FetchAPI() ([]byte, error) {
	req, err := http.NewRequest(
		r.Method,
		r.APIURL,
		r.Data,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", r.Contenttype)
	req.Header.Set("Authorization", "Bearer "+r.Accesstoken)
	client := &http.Client{
		Timeout: time.Duration(r.Dtime) * time.Second,
	}
	res, err := client.Do(req)
	if err != nil || res.StatusCode-300 >= 0 {
		var msg []byte
		var er string
		if res == nil {
			msg = []byte(err.Error())
			er = err.Error()
		} else {
			errmsg, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v. ", err)
				os.Exit(1)
			}
			msg = errmsg
			er = "Status Code: " + strconv.Itoa(res.StatusCode)
		}
		return msg, errors.New(er)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return body, err
}

// FetchAPIRaw : For fetching data to URL. Raw data (http.Response) from API is returned.
func (r *RequestParams) FetchAPIRaw() (*http.Response, error) {
	req, err := http.NewRequest(
		r.Method,
		r.APIURL,
		r.Data,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", r.Contenttype)
	req.Header.Set("Authorization", "Bearer "+r.Accesstoken)
	client := &http.Client{
		Timeout: time.Duration(r.Dtime) * time.Second,
	}
	res, err := client.Do(req)
	if err != nil || res.StatusCode-300 >= 0 {
		var er string
		if res == nil {
			er = err.Error()
		} else {
			er = "Status Code: " + strconv.Itoa(res.StatusCode)
		}
		return res, errors.New(er)
	}
	return res, err
}

// FetchAPIres : For fetching data to URL.
func (r *RequestParams) FetchAPIres() (*http.Response, error) {
	req, err := http.NewRequest(
		r.Method,
		r.APIURL,
		r.Data,
	)
	if err != nil {
		return nil, err
	}
	if len(r.ContentLength) > 0 {
		req.Header.Set("Content-Length", r.ContentLength)
	}
	if len(r.ContentRange) > 0 {
		req.Header.Set("Content-Range", r.ContentRange)
	}
	if len(r.Contenttype) > 0 {
		req.Header.Set("Content-Type", r.Contenttype)
	}
	if len(r.Accesstoken) > 0 {
		req.Header.Set("Authorization", "Bearer "+r.Accesstoken)
	}
	client := &http.Client{
		Timeout: time.Duration(r.Dtime) * time.Second,
	}
	return client.Do(req)
}
