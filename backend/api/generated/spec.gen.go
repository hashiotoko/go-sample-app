// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7RVXWsrRRj+K+XVy+1u0t6EvfILSi9EEPWmFBl33yRTdnfGmUkwhAWTvWmtH6WIoihi",
	"kSrWhoJQWk7P+TPTpOf8i8PMbpImadMc0l7NJMz78Tzv87zbhoDFnCWYKAl+G2RQx5jY64csxEi6n0oU",
	"5id+RWIeob3GhEafkzAUKCX40JAoyu8UD9yAxeAADcEvO5CQGMGHmx/Pri9O+3/9+yr7HlIHuGAchaIo",
	"78jXhhBlIChXlCXgg87+1NmVzk5090hnezr7T3cvwQHV4ia3VIImNZPV1JwNPjbB3XOdXW1+AA5UmYiJ",
	"Ah9ootbXxmloorCGwuTJm56XqX/w3WwDqQMCv2xQgSH4W6abIpUzhXB7FMq+2MFAmZofo+QskSjddzl1",
	"Pytb3qX7vkCi0NyHD+Y3dv38t8HuwWD3oP/NH7rTy7nS2e86+z8nbZL5RjFcEkUfVcHfasPbAqvgw1ve",
	"WBdeIQrvtiLSbWdOGzNk2EKzuM0zmlSZ6SGiARb4CtVwQZtEoaFHUWW0BzW2Kq3MVgnn4EAThczLl9yS",
	"WzJPGceEcAo+rNu/HOBE1S1cj3DqNcteTSAqMzO/DTVUs5QOvv1nsH9uhZfp7jOL60J3D/s//NR/8bPu",
	"/KK7+2BLCWJCNkPwYaPIOjzBUFBM1ZRYK5XMEbBEYWKLEs4jGtgM3o40lYcONLdpeaXTjH9SxxVDMkq1",
	"UidyRTaCADHE0LX8y0YcE9EawbF/DikwA5H34p+Q1MXXL4//fgi8lesGKnsui5wqjG3gwmocKYsIQVpL",
	"kXUH9k4vx25XF5MPEKa7h7kN51E1NjbkTkGp3mNh6414epo1+jjrb8HNNxmmRAPTJcUzTzOLb9nHUpDu",
	"9HIxzLrPa5tjM0zvteHN3mW/96vd4xPyWtiJdvkJEqOybt9qAzV5zUKE0cc57wKm5+Dc4vTBj6b5HDzZ",
	"1CacvsRc7qZzbG77GkVzyNW0AU5t0InOTm4Oz/pHGTjQEBH4UFeK+54XsYBEdSaVX6lUKpBup68DAAD/",
	"/5mUkd5fCQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
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
	res := make(map[string]func() ([]byte, error))
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
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
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
