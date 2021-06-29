package aws

var InstMap = map[string]service{}
var awsRegion = "us-east-1"

type service interface {
	ExecuteRules(resources string)
}
