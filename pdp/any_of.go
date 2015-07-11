package pdp

import (
	"encoding/xml"
)

type AnyOf struct {
	XMLName xml.Name `xml:"AnyOf"`
	AllOfs []AllOf `xml:"AllOf"`
}