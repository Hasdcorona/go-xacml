package rule

import (
	"encoding/xml"
	"pdp/target"
	"pdp/effect"
	"pdp/condition"
	"pdp/obligation"
	"pdp/advice"
)

type Rule struct {
	XMLName xml.Name `xml:"Rule"`
	RuleId string `xml:"RuleId,attr"`
	Effect Effect `xml:"Effect,attr"`
	Description string
	Target Target
}