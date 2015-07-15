package pdp

import (
	"encoding/xml"
)

type PolicyIdReference struct {
	XMLName  xml.Name `xml:"PolicyIdReference"`
	PolicyId string   `xml:"PolicyIdReference"`
}
