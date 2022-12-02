package mutation

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
	"github.com/wI2L/jsondiff"
	appsv1 "k8s.io/api/apps/v1"
)

// Mutator is a container for mutation
type Mutator struct {
	Logger *logrus.Entry
}

// NewMutator returns an initialised instance of Mutator
func NewMutator(logger *logrus.Entry) *Mutator {
	return &Mutator{Logger: logger}
}

// podMutators is an interface used to group functions mutating pods
type statefulSetMutator interface {
	Mutate(*appsv1.StatefulSet) (*appsv1.StatefulSet, error)
	Name() string
}

// MutatePodPatch returns a json patch containing all the mutations needed for
// a given pod
func (m *Mutator) MutateStatefulSetPatch(sts *appsv1.StatefulSet) ([]byte, error) {
	var stsName string
	if sts.Name != "" {
		stsName = sts.Name
	} else {
		if sts.ObjectMeta.GenerateName != "" {
			stsName = sts.ObjectMeta.GenerateName
		}
	}
	log := logrus.WithField("sts_name", stsName)

	// list of all mutations to be applied to the pod
	mutations := []statefulSetMutator{
		updateImage{Logger: log},
	}

	msts := sts.DeepCopy()

	// apply all mutations
	for _, m := range mutations {
		var err error
		msts, err = m.Mutate(msts)
		if err != nil {
			return nil, err
		}
	}

	// generate json patch
	patch, err := jsondiff.Compare(sts, msts)
	if err != nil {
		return nil, err
	}

	patchb, err := json.Marshal(patch)
	if err != nil {
		return nil, err
	}

	return patchb, nil
}
