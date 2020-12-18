package pod

import (
	context "context"
	//factory "knative.dev/pkg/client/injection/kube/informers/factory"
	factory "github.com/tektoncd/pipeline/pkg/client/injection/clusterInformer/factory"
	v1 "k8s.io/client-go/informers/core/v1"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
	logging "knative.dev/pkg/logging"
)

func init() {
	injection.Default.RegisterInformer(withInformer)
}

// Key is used for associating the Informer inside the context.Context.
type Key struct{}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := factory.Get(ctx)
	inf := f.Core().V1().Pods()
	return context.WithValue(ctx, Key{}, inf), inf.Informer()
}

// Get extracts the typed informer from the context.
func Get(ctx context.Context) v1.PodInformer {
	untyped := ctx.Value(Key{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch k8s.io/client-go/informers/core/v1.PodInformer from context.")
	}
	return untyped.(v1.PodInformer)
}
