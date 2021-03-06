package novamigrationtarget

import (
	"strings"

	util "github.com/openstack-k8s-operators/lib-common/pkg/util"
	novav1beta1 "github.com/openstack-k8s-operators/nova-operator/api/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BITSIZE
const (
	BITSIZE int = 4096
)

// Secret func
func Secret(cr *novav1beta1.NovaMigrationTarget) (*corev1.Secret, error) {

	privateKey, err := util.GeneratePrivateKey(BITSIZE)
	if err != nil {
		return nil, err
	}

	publicKey, err := util.GeneratePublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, err
	}

	privateKeyPem := util.EncodePrivateKeyToPEM(privateKey)

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      strings.ToLower(cr.Kind) + "-ssh-keys",
			Namespace: cr.Namespace,
		},
		Type: "Opaque",
		StringData: map[string]string{
			"identity":        privateKeyPem,
			"authorized_keys": publicKey,
		},
	}
	return secret, nil
}
