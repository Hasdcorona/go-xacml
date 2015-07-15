package pdp

import (
	"encoding/xml"
)

type ResourceMatch struct {
	XMLName                     xml.Name `xml:"ResourceMatch"`
	MatchId                     string   `xml:"MatchId,attr"`
	AttributeValue              AttributeValue
	ResourceAttributeDesignator ResourceAttributeDesignator
}
