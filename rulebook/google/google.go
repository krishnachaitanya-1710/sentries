package google

var InstMap = map[string]service{}

type service interface {
	ExecuteRules(resources string)
}
