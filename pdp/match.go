package pdp

import (
	"encoding/xml"
)

type Match struct {
	XMLName             xml.Name `xml:"Match"`
	MatchId             string   `xml:"MatchId,attr"`
	AttributeValue      AttributeValue
	AttributeDesignator AttributeDesignator
	AttributeSelector   AttributeSelector
}
