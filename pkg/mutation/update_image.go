package mutation

import (
	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
)

// minLifespanTolerations is a container for mininum lifespan mutation
type updateImage struct {
	Logger logrus.FieldLogger
}

// minLifespanTolerations implements the podMutator interface
var _ statefulSetMutator = (*updateImage)(nil)

// Name returns the minLifespanTolerations short name
func (ui updateImage) Name() string {
	return "min_lifespan"
}

// Mutate returns a new mutated pod according to lifespan tolerations rules
func (ui updateImage) Mutate(sts *appsv1.StatefulSet) (*appsv1.StatefulSet, error) {
	msts := sts.DeepCopy()

	msts.Spec.Template.Spec.Containers[0].Image = "gcr.io/mortent-dev-kube/doesNotExist:latest"

	return msts, nil
}
