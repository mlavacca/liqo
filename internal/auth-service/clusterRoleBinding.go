package auth_service

import (
	"context"
	"github.com/liqotech/liqo/pkg/discovery"
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (authService *AuthServiceCtrl) createClusterRoleBinding(sa *v1.ServiceAccount, clusterRole *rbacv1.ClusterRole, remoteClusterId string) (*rbacv1.ClusterRoleBinding, error) {
	rb := &rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: sa.Name,
			Labels: map[string]string{
				discovery.LiqoManagedLabel: "true",
				discovery.ClusterIdLabel:   remoteClusterId,
			},
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      sa.Name,
				Namespace: sa.Namespace,
			},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: rbacv1.SchemeGroupVersion.Group,
			Kind:     "ClusterRole",
			Name:     clusterRole.Name,
		},
	}
	return authService.clientset.RbacV1().ClusterRoleBindings().Create(context.TODO(), rb, metav1.CreateOptions{})
}
