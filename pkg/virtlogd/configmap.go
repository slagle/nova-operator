package virtlogd

import (
	"strings"

	util "github.com/openstack-k8s-operators/lib-common/pkg/util"
	novav1beta1 "github.com/openstack-k8s-operators/nova-operator/api/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TemplatesConfigMap - custom virtlogd config map
func TemplatesConfigMap(cr *novav1beta1.Virtlogd, cmName string) *corev1.ConfigMap {

	cm := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ConfigMap",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cmName,
			Namespace: cr.Namespace,
		},
		Data: map[string]string{
			"virtlogd.conf": util.ExecuteTemplateFile(strings.ToLower(cr.Kind)+"/config/virtlogd.conf", nil),
			"config.json":   util.ExecuteTemplateFile(strings.ToLower(cr.Kind)+"/kolla_config.json", nil),
		},
	}

	return cm
}
