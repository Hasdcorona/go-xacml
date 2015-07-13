package pdp

import (
	"encoding/xml"
)

type PolicyDefaults struct {
	XMLName xml.Name `xml:"PolicyDefaults"`
	XPathVersion string `xml:"XPathVersion"`
}