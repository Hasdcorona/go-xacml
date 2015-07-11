package pdp

import (
	"encoding/xml"
)

type AllOf struct {
	XMLName xml.Name `xml:"AllOf"`
	Matches []Match `xml:"Match"`
}