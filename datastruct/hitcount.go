package datastruct

import "encoding/xml"

//Response: represents the parent XML node from the HTTP response
type Response struct {
	XMLName  xml.Name `xml:"response"`
	Response string   `xml:"status,attr"`
	Result   Result   `xml:"result"`
	VsysName string   `xml:"name,attr"`
	RuleBase string   `xml:"rule-base"`
	Rules    []Rules  `xml:"rules"`
}

//Result: Response child node
type Result struct {
	XMLName      xml.Name     `xml:"result"`
	RuleHitCount RuleHitCount `xml:"rule-hit-count"`
}

//RuleHitCount: Result child node
type RuleHitCount struct {
	XMLName xml.Name `xml:"rule-hit-count"`
	Vsys    Vsys     `xml:"vsys"`
}

//Vsys: Rule-Hit-Count child node
type Vsys struct {
	XMLName    xml.Name   `xml:"vsys"`
	VsysNumber VsysNumber `xml:"entry"`
}

//VsysNumber: Vsys child node
type VsysNumber struct {
	XMLName  xml.Name `xml:"entry"`
	Name     string   `xml:"name,attr"`
	Rulebase Rulebase `xml:"rule-base"`
}

//Rulebase: VsysNumber child node
type Rulebase struct {
	XMLName  xml.Name `xml:"rule-base"`
	RuleType RuleType `xml:"entry"`
}

//RuleType: Rulebase child node
type RuleType struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`
	Rules   Rules    `xml:"rules"`
}

//Vsys: RuleType child node
type Rules struct {
	XMLName     xml.Name      `xml:"rules"`
	RuleDetails []RuleDetails `xml:"entry"`
}

//RuleDetails: Rules child nodes
type RuleDetails struct {
	XMLName      xml.Name `xml:"entry"`
	Name         string   `xml:"name,attr"`
	HitCountInt  int      `xml:"hit-count"`
	LastHitTS    int64      `xml:"last-hit-timestamp"`
	LastRstTS    int      `xml:"last-reset-timestamp"`
	FirstHitTS   int      `xml:"first-hit-timestamp"`
	RuleCreateTS int      `xml:"rule-creation-timestamp"`
	RuleModeTS   int      `xml:"rule-modification-timestamp"`
}
