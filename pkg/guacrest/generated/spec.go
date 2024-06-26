// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package generated

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

	"H4sIAAAAAAAC/+RWzW4bNxB+lQFbQG2xlYy0vfgWu0ljIGmMyEAPSQ40d6RlzB2uh6RUJdC7F0Nq5ZWy",
	"atyDT73tkjOcb/6+mS/K+LbzhBSDOv+iOs26xYic/6710pKO1tO8QyMnNQbDtpMjda5uGoRuLwPG08Iu",
	"E5e/hWeIDcJ9Qt5MPxDATzC51kuc2884gdChsQuLIQtRam+RwS+AMSQXAzDGxIT1TvEycfA8AftwA7cb",
	"6BhX1qcARjsXQFM9eHjd6Cj4EKLfaX0gVSkr2DMsVSnSLapz1R26WqlgGmx1jgn7DjlazDEpQOQrbjrR",
	"DJEtLdW2Ur1zg0tLEZfIarut+iN/+wlNVFs5Ygydp1BevtD1HzriWm/kz3iKSFE+ddc5azK42acgkf8y",
	"gPc940Kdq+9mD5mcldswe8HsuZj6OnMBeYUMSMYnishYgyZAUZFUEppoaSmxkwzVOmq41eYOqRZnL3T9",
	"Du8Thvj0aC90DVyMVRCSaUAHWLBvwdJKO1uDZ2htCIJ3UMLbSl2JZ6TdPDtbLDw53t4oFKuwE6zUdWL3",
	"2v7HkB3W30NPXtHCfwvikfQRBBuxDd98IrFTD+WrmfVGleK9T5axVufvj1ENzHwcLfzDeD0HZ0OU7u8S",
	"u5Bf35kXdPusHUbiDYaglzjSikfgesGvoVQj4TyEdukpakuFpUzu/R2bsMUVQus5cyCGKVwtRIoRNCOQ",
	"z3cVwJ/4dyysAWvrHNwikHXTTEWHHj1IjvLLjY/aXUqzPo5hShbG4rMVEhR3KTlXKd8h6c6qc/XL9Gx6",
	"Jrh0bDKkmSbtNsGGWY0dUo1kdmCXmGEI/hK/WjIp0p/x96FsdTBV3o9X24PI7GjqbKuxsRM8R/Bcl6ER",
	"8yAyd5KH3cBYZMIgs5nAz3AzuIe1jU3WaOyywRAHw6f3MfavBOMZjeb69CvOr+WRtx3SfP4S9hrl6+TA",
	"EQfUsE4jJxyOHaTUSvXuHclDaff4oJb3Sf14NE+enZ2d6u293Gzfp9tK/foYhQHvbyv122NUxjg46z57",
	"lLl+KGZWSG2reSMcK2myi03OQetDBNt2nqOmCAelKmqzBrWLzeeTdfsq3182aO7UeBgfTdcjvXY8HWrR",
	"3m0+uylsAxSMx34WZGAE2kChuJVr6lRnHlp9aanO+pE1BRvtCg/i1LeRpS7Faa71VzrML96+yVuVfL+e",
	"P5eFq0cvLJeCbA3Fk/1rG2B0OTyhsV1Wl/ur8NaYxIxkMB/C9d3yxX3SbvTV6GEhmI2nYKXPpVNW2iHF",
	"fRcWEj1M5bsdNT8tB/3VoHA9OEt3AW4xrhEJyCcK0KYQheVbXaMsqbVdCkN4BptjvAEwmvYSn7L4BoQY",
	"4Ac7xWkegz9OYZ5X2Q1M5GoiEdHO+TWkvOhIbiTyOkLtaRKhY7+yNZZk7GyWpIaUR0JJa40LnVyUkoNJ",
	"kZtM4cZDQM2myWt1Ylf1Znt3+sW6np4kNYnGpafa5iiNsVmx12uMsNgo2wue3pc9R/dVMD210svwq/6l",
	"M0dNHQbu0cb2fp029//i574NT/GNMOP2nwAAAP///dINuX8OAAA=",
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
