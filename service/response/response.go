package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// response

type Response struct {
	Code  int32  `json:"code"`
	Msg   string `json:"msg"`
	Data  any    `json:"data,omitempty"`
	Debug any    `json:"debug,omitempty"`
}

func Ok(c *gin.Context) {
	// Result(200, "success", map[string]any{}, nil, c)
	Result(200, "success", nil, nil, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(200, "success", data, nil, c)
}

func Fail(code int32, message string, debug any, c *gin.Context) {
	Result(code, message, nil, debug, c)
}

func Result(code int32, msg string, data interface{}, debug any, c *gin.Context) {

	// contentType := c.Request.Header.Get("Content-Type")
	// var isJson, isProtobuf bool

	// c.ShouldBindJSON()
	// isJson = strings.Contains(contentType, "application/json")
	// isProtobuf = strings.Contains(contentType, "application/x-protobuf")
	// c.GetRawData()
	if debug == nil && c.Request.Method == http.MethodPost {
		debug, _ = c.GetRawData()
	}

	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
		debug,
	})
	// return

	// var jsonContentType = []string{"application/json; charset=utf-8"}
	// var jsonpContentType = []string{"application/javascript; charset=utf-8"}
	// var jsonAsciiContentType = []string{"application/json"}
	// var protobufContentType = []string{"application/x-protobuf"}

	// MIMEJSON              = "application/json"
	// MIMEHTML              = "text/html"
	// MIMEXML               = "application/xml"
	// MIMEXML2              = "text/xml"
	// MIMEPlain             = "text/plain"
	// MIMEPOSTForm          = "application/x-www-form-urlencoded"
	// MIMEMultipartPOSTForm = "multipart/form-data"
	// MIMEPROTOBUF          = "application/x-protobuf"
	// MIMEMSGPACK           = "application/x-msgpack"
	// MIMEMSGPACK2          = "application/msgpack"
	// MIMEYAML              = "application/x-yaml"
	// MIMETOML              = "application/toml"
}