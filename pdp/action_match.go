package pdp

import (
	"encoding/xml"
)

type ActionMatch struct {
	XMLName xml.Name `xml:"ActionMatch"`
	MatchId string `xml:"MatchId,attr"`
	AttributeValue AttributeValue
	ActionAttributeDesignator ActionAttributeDesignator
}