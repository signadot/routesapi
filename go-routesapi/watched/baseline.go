package watched

import (
	"fmt"
	"os"
)

// Baseline is the type of a baseline workload.  We use
// this instead of the grpc routes.Baseline because it
// can be used as a map key, whereas the grpc one cannot.
type Baseline struct {
	Kind      string `json:"kind"`
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

// BaselineFromEnv grabs the baseline from the environment
// if possible.  Otherwise, it returns nil and a non-nil error.
func BaselineFromEnv() (*Baseline, error) {
	k := os.Getenv(BaselineKindEnvVar)
	if k == "" {
		return nil, fmt.Errorf("%s not found in env", BaselineKindEnvVar)
	}
	ns := os.Getenv(BaselineNamespaceEnvVar)
	if ns == "" {
		return nil, fmt.Errorf("%s not found in env", BaselineNamespaceEnvVar)
	}
	n := os.Getenv(BaselineNameEnvVar)
	if n == "" {
		return nil, fmt.Errorf("%s not found in env", BaselineNameEnvVar)
	}
	return &Baseline{Kind: k, Namespace: ns, Name: n}, nil
}
