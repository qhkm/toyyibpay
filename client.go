package toyyibpay

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ajg/form"
)

// Client contains all resource for toyyibpay
type Client struct {
	UserSecretKey string
	Backend       *http.Client
}

// ErrorResponse ...
type ErrorResponse struct {
	Response *http.Response `json:"-"`
	Name     string         `json:"name"`
	DebugID  string         `json:"debug_id"`
	Message  string         `json:"message"`
}

// NewClient creates new client for toyyibpay
func NewClient(secretKey string) (*Client, error) {

	if secretKey == "" {
		return nil, errors.New("secretKey are required to create a Client")
	}
	client := &Client{
		UserSecretKey: secretKey,
		Backend:       &http.Client{},
	}
	return client, nil
}

// NewRequest makes api request to toyyibpay API
func (c *Client) NewRequest(task string, payload interface{}) (*http.Request, error) {

	// var buf io.Reader
	var b url.Values
	var err error
	if payload != nil {
		b, err = form.EncodeToValues(payload)
		if err != nil {
			return nil, err
		}
		// buf = bytes.NewBuffer()
	}

	url := GetAPIPath(task)
	return http.NewRequest("POST", url, strings.NewReader(b.Encode()))
}

// MockCallWithJSONResponse mocks call to create single bill
func (c *Client) MockCallWithJSONResponse(req *http.Request, resVal interface{}) error {
	mockData := []byte(`[{"BillCode": "05e0hahi"}]`)
	err := json.NewDecoder(bytes.NewBuffer(mockData)).Decode(&resVal)
	if err != nil {
		return err
	}
	return nil
}

// CallWithJSONResponse makes API request to toyyibpay account and return unmarshalled response value
func (c *Client) CallWithJSONResponse(req *http.Request, resVal interface{}) error {
	var (
		err  error
		resp *http.Response
		data []byte
	)

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req = req.WithContext(ctx)
	// Set default headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// API request
	resp, err = c.Backend.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		// errResp := &ErrorResponse{Response: resp}
		errResp := errors.New("unexpected error occured")
		data, err = ioutil.ReadAll(resp.Body)

		if err == nil && len(data) > 0 {
			json.Unmarshal(data, errResp)
		}
		return err
	}

	dec := json.NewDecoder(resp.Body)
	for {
		if err := dec.Decode(&resVal); err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}

	return nil
}

// CallWithHTMLResponse makes API request to toyyibpay account and return unmarshalled response value
func (c *Client) CallWithHTMLResponse(req *http.Request, resVal interface{}) error {
	var (
		err  error
		resp *http.Response
		data []byte
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req = req.WithContext(ctx)

	// Set default headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err = c.Backend.Do(req)
	htmlResponse, err := ioutil.ReadAll(resp.Body)
	stringifiedHTMLResponse := string(htmlResponse)

	returnValue := resVal.(*RunBillResponse)
	returnValue.Body = &stringifiedHTMLResponse

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		// errResp := &ErrorResponse{Response: resp}
		errResp := errors.New("unexpected error occured")
		data, err = ioutil.ReadAll(resp.Body)

		if err == nil && len(data) > 0 {
			json.Unmarshal(data, errResp)
		}
		return err
	}

	dec := json.NewDecoder(resp.Body)
	for {
		if err := dec.Decode(&resVal); err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}

	return nil
}
