package validation

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
)

// nameValidator is a container for validating the name of pods
type nameValidator struct {
	Logger logrus.FieldLogger
}

// nameValidator implements the podValidator interface
var _ deploymentValidator = (*nameValidator)(nil)

// Name returns the name of nameValidator
func (n nameValidator) Name() string {
	return "name_validator"
}

// Validate inspects the name of a given deployment and returns validation.
// The returned validation is only valid if the deployment name does not contain some
// bad string.
func (n nameValidator) Validate(dep *appsv1.Deployment) (validation, error) {
	badString := "offensive"

	if strings.Contains(dep.Name, badString) {
		v := validation{
			Valid:  false,
			Reason: fmt.Sprintf("deployment name contains %q", badString),
		}
		return v, nil
	}

	return validation{Valid: true, Reason: "valid name"}, nil
}
