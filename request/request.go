package request

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/medivhzhan/weapp/v3/logger"
)

type Request struct {
	http *http.Client
	// 获取日志记录器
	logger      func() logger.Logger
	contentType ContentType
}

func NewRequest(http *http.Client, ctp ContentType, logger func() logger.Logger) *Request {
	return &Request{
		http:        http,
		logger:      logger,
		contentType: ctp,
	}
}

func (cli *Request) Get(url string, response interface{}) error {
	cli.logger().Info(context.Background(), "request url: %s", url)

	resp, err := cli.http.Get(url)
	if err != nil {
		cli.logger().Error(context.Background(), "get error: %s", err)
		return err
	}
	defer resp.Body.Close()

	switch cli.contentType {
	case ContentTypeXML:
		return xml.NewDecoder(resp.Body).Decode(response)
	case ContentTypeJSON:
		return json.NewDecoder(resp.Body).Decode(response)
	default:
		return errors.New("invalid content type")
	}
}

func (cli *Request) GetWithBody(url string) (*http.Response, error) {
	cli.logger().Info(context.Background(), "request url: %s", url)
	rsp, err := cli.http.Get(url)
	if err != nil {
		cli.logger().Error(context.Background(), "get with body error: %s", url)
		return nil, err
	}

	return rsp, nil
}

func (cli *Request) Post(url string, params interface{}, response interface{}) error {
	resp, err := cli.PostWithBody(url, params)
	if err != nil {
		cli.logger().Error(context.Background(), "post error: %s", err)
		return err
	}
	defer resp.Body.Close()

	switch cli.contentType {
	case ContentTypeXML:
		return xml.NewDecoder(resp.Body).Decode(response)
	case ContentTypeJSON:
		return json.NewDecoder(resp.Body).Decode(response)
	default:
		return errors.New("invalid content type")
	}
}

func (cli *Request) PostWithBody(url string, params interface{}) (*http.Response, error) {
	cli.logger().Info(context.Background(), "request url: %s", url)
	cli.logger().Info(context.Background(), "request params: %+v", params)
	buf := new(bytes.Buffer)
	if params != nil {
		switch cli.contentType {
		case ContentTypeXML:
			err := xml.NewEncoder(buf).Encode(params)
			if err != nil {
				return nil, err
			}
		case ContentTypeJSON:
			enc := json.NewEncoder(buf)
			enc.SetEscapeHTML(false)
			err := enc.Encode(params)
			if err != nil {
				return nil, err
			}
		default:
			return nil, errors.New("invalid content type")
		}
	}

	rsp, err := cli.http.Post(url, cli.contentType.String(), buf)
	if err != nil {
		cli.logger().Error(context.Background(), "post with body error: %s", url)
		return nil, err
	}

	return rsp, nil
}

func (cli *Request) FormPostWithFile(url, field, filename string, response interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return cli.FormPost(url, field, filename, file, response)
}

func (cli *Request) FormPost(url, field, filename string, reader io.Reader, response interface{}) error {
	cli.logger().Info(context.Background(), "request url: %s", url)
	// Prepare a form that you will submit to that URL.
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	fw, err := w.CreateFormFile(field, filename)
	if err != nil {
		return err
	}

	if _, err = io.Copy(fw, reader); err != nil {
		return err
	}

	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return err
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := cli.http.Do(req)
	if err != nil {
		cli.logger().Error(context.Background(), "form post error: %s", err)
		return err
	}
	defer resp.Body.Close()

	switch cli.contentType {
	case ContentTypeXML:
		return xml.NewDecoder(resp.Body).Decode(response)
	case ContentTypeJSON:
		return json.NewDecoder(resp.Body).Decode(response)
	default:
		return errors.New("invalid content type")
	}
}
