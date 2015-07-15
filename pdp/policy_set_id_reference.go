package pdp

import (
	"encoding/xml"
)

type PolicySetIdReference struct {
	XMLName         xml.Name `xml:"PolicySetIdReference"`
	Version         string   `xml:"Version,attr"`
	EarliestVersion string   `xml:"EarliestVersion,attr"`
	LatestVersion   string   `xml:"LatestVersion,attr"`
	Reference       string   `xml:"PolicySetIdReference"`
}
