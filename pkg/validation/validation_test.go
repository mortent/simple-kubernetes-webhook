package validation

import (
	"io/ioutil"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestValidatePod(t *testing.T) {
	v := NewValidator(logger())

	pod := &appsv1.Deployment{
		ObjectMeta: v1.ObjectMeta{
			Name: "lifespan",
		},
		Spec: appsv1.DeploymentSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:  "lifespan",
						Image: "busybox",
					}},
				},
			},
		},
	}

	val, err := v.ValidateDeployment(pod)
	assert.Nil(t, err)
	assert.True(t, val.Valid)
}

func logger() *logrus.Entry {
	mute := logrus.StandardLogger()
	mute.Out = ioutil.Discard
	return mute.WithField("logger", "test")
}
