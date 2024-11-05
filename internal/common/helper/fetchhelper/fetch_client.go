package fetchhelper

import (
	"bytes"
	"encoding/json"
	"errors"
	"example/internal/common/helper/querystringhelper"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

type (
	FetchClient interface {
		PostJSON(path string, body any, target any) error
		PostFile(path string, args any, target any, filename string, file io.Reader) error
		Get(path string, args any, target any) error
	}

	fetchClient struct {
		HttpClient *http.Client
		BaseURL    string
	}

	ClientOptions struct {
		HttpClient *http.Client
		BaseURL    string
	}
)

func NewFetchClient(options *ClientOptions) FetchClient {
	return &fetchClient{
		HttpClient: options.HttpClient,
		BaseURL:    options.BaseURL,
	}
}

func (f *fetchClient) PostJSON(path string, body any, target any) error {
	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("Got error while encode json data, err: %v", err)
		return err
	}
	bodyData := bytes.NewBuffer(jsonData)

	url := fmt.Sprintf("%s/%s", f.BaseURL, path)
	urlWithParams := url

	req, err := http.NewRequest(http.MethodPost, urlWithParams, bodyData)
	if err != nil {
		log.Fatalf("Got error while creating request, err: %v", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	return f.do(req, url, target)
}
func (f *fetchClient) PostFile(path string, args any, target any, filename string, file io.Reader) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}

	params, err := querystringhelper.FromStruct(args)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/%s", f.BaseURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, params.Encode())

	req, err := http.NewRequest("POST", urlWithParams, body)
	if err != nil {
		log.Fatalf("Got error while creating request, err: %v", err)
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return f.do(req, url, target)
}
func (f *fetchClient) Get(path string, args any, target any) error {
	params, err := querystringhelper.FromStruct(args)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s%s", f.BaseURL, path)
	urlWithParams := fmt.Sprintf("%s?%s", url, params.Encode())

	req, err := http.NewRequest(http.MethodGet, urlWithParams, nil)
	if err != nil {
		log.Fatalf("Got error while creating request, err: %v", err)
		return err
	}

	return f.do(req, url, target)
}

func (f *fetchClient) do(req *http.Request, url string, target any) error {
	log.Printf("Request to url => %s", url)
	resp, err := f.HttpClient.Do(req)
	if err != nil {
		log.Fatalf("Got error while calling to url => %s \nerr: %v", url, err)
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Got error while close body, err:  %v", err)
		}
	}(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return makeHTTPClientError(url, resp)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("Got error while reading body from response, err: %v", err))
	}

	err = json.Unmarshal(b, target)
	if err != nil {
		return errors.New(fmt.Sprintf("Got error while decoding body from body, err: %v", err))
	}
	return nil
}

type httpClientError struct {
	msg  string
	code int
}

func makeHTTPClientError(url string, resp *http.Response) error {
	body, _ := io.ReadAll(resp.Body)
	msg := fmt.Sprintf("HTTP request failure on %s:\n%d: %s", url, resp.StatusCode, string(body))

	return &httpClientError{
		msg:  msg,
		code: resp.StatusCode,
	}
}

func (e *httpClientError) Error() string { return e.msg }
