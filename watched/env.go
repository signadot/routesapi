package watched

import "os"

const (
	SignadotEnvPrefix       = "SIGNADOT_"
	BaselineEnvPrefix       = SignadotEnvPrefix + "BASELINE_"
	BaselineKindEnvVar      = BaselineEnvPrefix + "KIND"
	BaselineNamespaceEnvVar = BaselineEnvPrefix + "NAMESPACE"
	BaselineNameEnvVar      = BaselineEnvPrefix + "NAME"
)

func GetRouteServerAddr() string {
	rsa := os.Getenv(SignadotEnvPrefix + "ROUTESERVER")
	if rsa != "" {
		return rsa
	}
	return "routeserver.signadot.svc:7777"
}
