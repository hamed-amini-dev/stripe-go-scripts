package request

/*
Define new customize http request
*/
import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

type GenericRequest struct {
	Method      string
	URL         string
	Host        string
	Headers     http.Header
	QueryParams url.Values
	Body        []byte
	PathParams  map[string]string
	RemoteAddr  string
	r           *http.Request
}

// NewGenericHTTPRequestFromHTTPRequest- Returns a new Generic HTTP Request from http.request
func NewGenericHTTPRequestFromHTTPRequest(r *http.Request) (*GenericRequest, error) {
	gr := new(GenericRequest)

	var err error
	gr.Headers = r.Header
	gr.URL = r.URL.String()
	gr.Host = r.Host
	gr.PathParams = mux.Vars(r)
	gr.QueryParams = r.URL.Query()
	gr.Method = r.Method
	gr.RemoteAddr = r.RemoteAddr
	gr.r = r

	reader := bufio.NewReader(r.Body)
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	gr.Body = body
	return gr, nil
}

// String- Returns the request as a string | Compacts the request by omiting
// the insignificant space characters and appends it to compactBody which is
// of type *bytes.Buffer.

func (req *GenericRequest) String() string {
	request, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	compactBody := new(bytes.Buffer)
	err = json.Compact(compactBody, req.Body)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("httpRequest: %v, RequestBody: %v", string(request), compactBody.String())
}

func (req *GenericRequest) GetPathParameter(key string) string {
	params := mux.Vars(req.r)
	return params[key]
}
