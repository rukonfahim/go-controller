package main

import (
	"context"
	"flag"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	mygroupv1 "mygroup.com/api/v1" // import the package containing the generated clientset for your CRD
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(mygroupv1.AddToScheme(scheme)) // add your CRD's scheme to the runtime scheme 
}

type MyReconciler struct {
	client.Client // embed the controller-runtime client 
}

func (r *MyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var myresource mygroupv1.MyResource // create an instance of your custom resource 
	if err := r.Get(ctx, req.NamespacedName, &myresource); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	fmt.Printf("Image: %s\n", myresource.Spec.Image)
	fmt.Printf("Port : %d\n", myresource.Spec.Port)

	return ctrl.Result{}, nil // return no error 
}

func main() {
	flag.Parse()

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:scheme,
	})
	if err != nil {
		panic(err)
	}

	if err := ctrl.NewControllerManagedBy(mgr).
		 For(&mygroupv1.MyResource{}). // watch for changes to instances of MyResource 
		 Complete(&MyReconciler{Client:mgr.GetClient()}); err != nil { // specify the reconciler 
		 panic(err)
	}

	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		 panic(err)
	}
}