package httpc

import (
	"encoding/xml"
	"io"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v2"
)

// Copy copies the response body into any io.Writer
func Copy(w io.Writer) func(resp *http.Response) error {
	return func(resp *http.Response) error {
		_, err := io.Copy(w, resp.Body)
		return err
	}
}

// ParseJSON parses the response body as JSON into a struct
func ParseJSON(v interface{}) func(resp *http.Response) error {
	return func(resp *http.Response) error {
		return jsoniter.NewDecoder(resp.Body).Decode(v)
	}
}

// ParseYAML parses the response body as YAML into a struct
func ParseYAML(v interface{}) func(resp *http.Response) error {
	return func(resp *http.Response) error {
		return yaml.NewDecoder(resp.Body).Decode(v)
	}
}

// ParseXML parses the response body as XML into a struct
func ParseXML(v interface{}) func(resp *http.Response) error {
	return func(resp *http.Response) error {
		return xml.NewDecoder(resp.Body).Decode(v)
	}
}
