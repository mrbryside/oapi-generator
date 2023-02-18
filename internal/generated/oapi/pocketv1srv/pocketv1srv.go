// Package pocketv1srv provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.2 DO NOT EDIT.
package pocketv1srv

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /pockets)
	GetPockets(ctx echo.Context) error

	// (POST /pockets)
	CreatePocket(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetPockets converts echo context to params.
func (w *ServerInterfaceWrapper) GetPockets(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPockets(ctx)
	return err
}

// CreatePocket converts echo context to params.
func (w *ServerInterfaceWrapper) CreatePocket(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreatePocket(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/pockets", wrapper.GetPockets)
	router.POST(baseURL+"/pockets", wrapper.CreatePocket)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9yVwU/cOhDG/xVr3jsmZHnvgnyjVVX1UlB7RKjyOpOsIfGY8YSCUP73yk52F9igUolK",
	"VU8b2ePP3/xsf/sAlvpAHr1E0A8Q7QZ7kz8/MBOnj8AUkMVhHrZUY/qV+4CgwXnBFhnGAnqM0bSPJ6Ow",
	"8y2MYwGMN4NjrEFfTBL7+stiW0/rK7SStD6inJO9RolfMAbyEU+ZzX3SdoJ9dvIvYwMa/qn2LVSz/2pa",
	"/J7RCG4Vku68kcliYwFP624GjHLYMt4F9BG/mZ4Gn+cb4t4IaGg6MgI7XT/06wmGq5chOW+p/zUpb/oF",
	"qAXclWSCKxPNFn2Jd8KmFNNmz7emc7WRtGCHfnx+EFn5uafieb9L57PI9wDc2nTGW3xdm38tZVdD8UrU",
	"xQ7ZAfTk5HpYI3sUjGXLNITyFjk68uW18zVo+Ewep/2dbygZrDFadkEcedBwFtCfBqdiQOsaZ00aVw2x",
	"MirkE1UmuETHSZf23g0GdXb6FQqYNwQNx0ero1UCRwF9WqXh/zxUQDCyyXiqaX3+blEODbUoaluTlThb",
	"+pSa2UcAJJ7THctS/62OpyTygtPxmhC6uZ3qKibpbZT9LCheCppM8anZuU7l8FD8KFVqbMzQyZuZmqJ3",
	"wcLg04WxgrXCfU2guMB2epxRGeXx+0z5APJUdL6d5CkC31F9/2bNLEXsQmvzVRNSpq7h8QsSHnD8jVdg",
	"+b/ipfNXNhWmd/PH3IBUFZHT2wR98QADd6BhIxJ0VXVkTbehKPpkdbKqYLwcfwQAAP//h5jtIvUHAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
