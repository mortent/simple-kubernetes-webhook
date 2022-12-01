package validation

import (
	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
)

// nameValidator is a container for validating the name of pods
type replicasValidator struct {
	Logger logrus.FieldLogger
}

// nameValidator implements the podValidator interface
var _ deploymentValidator = (*replicasValidator)(nil)

// Name returns the name of nameValidator
func (n replicasValidator) Name() string {
	return "replicas_validator"
}

// Validate inspects the name of a given deployment and returns validation.
// The returned validation is only valid if the deployment name does not contain some
// bad string.
func (n replicasValidator) Validate(dep *appsv1.Deployment) (validation, error) {
	if dep.Spec.Replicas != nil && *dep.Spec.Replicas == 1 {
		v := validation{
			Valid:  false,
			Reason: "deployment replicas is 1",
		}
		return v, nil
	}
	return validation{Valid: true, Reason: "valid replicas"}, nil
}
