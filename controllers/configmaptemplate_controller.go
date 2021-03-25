/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	k8sv1alpha1 "github.com/spigwitmer/configmap-template/api/v1alpha1"
)

// ConfigMapTemplateReconciler reconciles a ConfigMapTemplate object
type ConfigMapTemplateReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=k8s.camtap.io,resources=configmaptemplates,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=k8s.camtap.io,resources=configmaptemplates/status,verbs=get;update;patch

func (r *ConfigMapTemplateReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("configmaptemplate", req.NamespacedName)

	// your logic here
	cmt := k8sv1alpha1.ConfigMapTemplateList{}
	err := r.List(ctx, &cmt)
	if err != nil {
		log.Error(err, "Error listing ConfigMapTemplates")
		return ctrl.Result{}, err
	}
	log.V(1).Info("ConfigMapTemplate count", "key", "configmaptemplate-count", "count", len(cmt.Items))

	return ctrl.Result{}, nil
}

func (r *ConfigMapTemplateReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&k8sv1alpha1.ConfigMapTemplate{}).
		Complete(r)
}
