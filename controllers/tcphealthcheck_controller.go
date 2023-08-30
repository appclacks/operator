/*
Copyright 2023.

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

	goclient "github.com/appclacks/cli/client"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	healthchecksv1alpha1 "appclacks.com/operator/api/v1alpha1"
)

// TCPHealthcheckReconciler reconciles a TCPHealthcheck object
type TCPHealthcheckReconciler struct {
	client.Client
	Scheme    *runtime.Scheme
	Appclacks *goclient.Client
}

//+kubebuilder:rbac:groups=healthchecks.appclacks.com,resources=tcphealthchecks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=healthchecks.appclacks.com,resources=tcphealthchecks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=healthchecks.appclacks.com,resources=tcphealthchecks/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the TCPHealthcheck object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *TCPHealthcheckReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	healthcheck := &healthchecksv1alpha1.TCPHealthcheck{}
	if err := r.Get(ctx, req.NamespacedName, healthcheck); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	err := reconcile(ctx, log, r.Client, r.Appclacks, healthcheck, healthcheck.Name)
	return ctrl.Result{}, err
}

// SetupWithManager sets up the controller with the Manager.
func (r *TCPHealthcheckReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&healthchecksv1alpha1.TCPHealthcheck{}).
		Complete(r)
}
