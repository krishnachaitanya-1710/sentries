package globalvar

var AppName = "sentries"
var Resources []string
var SkipRules []string
var RuleCount int
var ResourceCount int
var Version = "0.1.0"
var ResourceArns []string
var RuleNames []string
var ViolationContext []string
var ComplianceStatus bool
var ResponseFile = "sentries_response.json"

type RuleConfig struct {
	ruleVersion, ruleName, ruleCategory, ruleId, ruleDescription, severity string
}
type ComplianceRules struct {
	rules []RuleConfig
}

type ResourceConfig struct {
	complianceStatus, violationContext string
}

type ComplianceScanResult struct {
	resources []ResourceConfig
}

type ComplianceSummary struct {
	ComplianceStatus     bool
	TotalRulesApplied    int
	TotalResourceScanned int
	RuleNames            []string
	ResourceArns         []string
	ViolationContext     []string
}
