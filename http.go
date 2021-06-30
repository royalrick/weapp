package weapp

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func (cli *client) getJSON(url string, response interface{}) error {

	resp, err := cli.httpClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(response)
}

// postJSON perform a HTTP/POST request with json body
func (cli *client) postJSON(url string, params interface{}, response interface{}) error {
	resp, err := cli.postJSONWithBody(url, params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(response)
}

// postJSONWithBody return with http body.
func (cli *client) postJSONWithBody(url string, params interface{}) (*http.Response, error) {
	b := new(bytes.Buffer)
	if params != nil {
		enc := json.NewEncoder(b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(params)
		if err != nil {
			return nil, err
		}
	}

	return cli.httpClient.Post(url, "application/json; charset=utf-8", b)
}

func (cli *client) postFormByFile(url, field, filename string, response interface{}) error {
	// Add your media file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return cli.postForm(url, field, filename, file, response)
}

func (cli *client) postForm(url, field, filename string, reader io.Reader, response interface{}) error {
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

	resp, err := cli.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(response)
}
