package httprequest

import (
	"errors"
	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

type Request *http.Request
type Response *http.Response

const (
	POST    = "POST"
	GET     = "GET"
	HEAD    = "HEAD"
	PUT     = "PUT"
	DELETE  = "DELETE"
	PATCH   = "PATCH"
	OPTIONS = "OPTIONS"
)

const (
	ContentTypeJSON       = "json"
	ContentTypeXml        = "xml"
	ContentTypeUrlencoded = "urlencoded"
	ContentTypeFrom       = "form"
	ContentTypeFromData   = "form-data"
	ContentTypeHTML       = "html"
	ContentTypeText       = "text"
	ContentTypeMultipart  = "multipart"
)

var NoMethodSpecified = errors.New("No method specified ")

type Client struct {
	Url               string
	Method            string
	Header            http.Header
	ContentType       string
	Data              map[string]interface{}
	SliceData         []interface{}
	FormData          url.Values
	QueryData         url.Values
	FileData          []File
	BounceToRawString bool
	RawString         string
	Client            *http.Client
	Transport         *http.Transport
	Cookies           []*http.Cookie
	Errors            []error
	BasicAuth         struct{ Username, Password string }
	NotResetClient    bool
	isClone           bool
	logger           *logrus.Logger
}

type File struct {
	Filename  string
	FieldName string
	Data      []byte
}

func (c *Client)SetLogger(log *logrus.Logger) *Client  {
	c.logger = log
}

func (c *Client) ResetClient() *Client {
	if c.NotResetClient{
		return c
	}
	c.Url = ""
	c.Method = ""
	c.Header = http.Header{}
	c.Data = make(map[string]interface{})
	c.SliceData = []interface{}{}
	c.FormData = url.Values{}
	c.QueryData = url.Values{}
	c.FileData = make([]File, 0)
	c.BounceToRawString = false
	c.RawString = ""
	c.ContentType = ""
	//c.TargetType = TypeJSON
	c.Cookies = make([]*http.Cookie, 0)
	c.Errors = nil

	return c
}

func (c *Client)Get(url string) *Client {
	gorequest.New().Get("zz").SendSlice(nil).End()
	c.ResetClient()
	c.Method = GET
	c.Url = url
	c.Errors = nil
	return c 
}

func doHttp()  {
	
}
