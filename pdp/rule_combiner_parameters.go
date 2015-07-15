package pdp

import (
	"encoding/xml"
)

type RuleCombinerParameters struct {
	XMLName            xml.Name            `xml:"RuleCombinerParameters"`
	RuleIdRef          string              `xml:"RuleIdRef,attr"`
	CombinerParameters []CombinerParameter `xml:"CombinerParameter"`
}
