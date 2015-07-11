package pdp

import (
	"encoding/xml"
)

type Target struct {
	XMLName xml.Name `xml:"Target"`
	Targets []AnyOf `xml:"AnyOf"`
}