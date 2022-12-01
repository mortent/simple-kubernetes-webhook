package validation

import (
	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
)

// Validator is a container for mutation
type Validator struct {
	Logger *logrus.Entry
}

// NewValidator returns an initialised instance of Validator
func NewValidator(logger *logrus.Entry) *Validator {
	return &Validator{Logger: logger}
}

// podValidators is an interface used to group functions mutating pods
type deploymentValidator interface {
	Validate(*appsv1.Deployment) (validation, error)
	Name() string
}

type validation struct {
	Valid  bool
	Reason string
}

// ValidatePod returns true if a pod is valid
func (v *Validator) ValidateDeployment(dep *appsv1.Deployment) (validation, error) {
	var deploymentName string
	if dep.Name != "" {
		deploymentName = dep.Name
	} else {
		if dep.ObjectMeta.GenerateName != "" {
			deploymentName = dep.ObjectMeta.GenerateName
		}
	}
	log := logrus.WithField("dep_name", deploymentName)
	log.Print("delete me")

	// list of all validations to be applied to the pod
	validations := []deploymentValidator{
		nameValidator{v.Logger},
		replicasValidator{v.Logger},
	}

	// apply all validations
	for _, v := range validations {
		var err error
		vp, err := v.Validate(dep)
		if err != nil {
			return validation{Valid: false, Reason: err.Error()}, err
		}
		if !vp.Valid {
			return validation{Valid: false, Reason: vp.Reason}, err
		}
	}

	return validation{Valid: true, Reason: "valid pod"}, nil
}
