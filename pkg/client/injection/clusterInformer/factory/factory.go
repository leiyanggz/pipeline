package clusterInformerFactory

import (
	context "context"

	client "github.com/tektoncd/pipeline/pkg/client/injection/clusterClient"
	informers "k8s.io/client-go/informers"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
	logging "knative.dev/pkg/logging"
)

func init() {
	injection.Default.RegisterInformerFactory(withInformerFactory)
}

// Key is used as the key for associating information with a context.Context.
type Key struct{}

func withInformerFactory(ctx context.Context) context.Context {
	c := client.Get(ctx)
	opts := make([]informers.SharedInformerOption, 0, 1)
	if injection.HasNamespaceScope(ctx) {
		opts = append(opts, informers.WithNamespace(injection.GetNamespaceScope(ctx)))
	}
	return context.WithValue(ctx, Key{},
		informers.NewSharedInformerFactoryWithOptions(c, controller.GetResyncPeriod(ctx), opts...))
}

// Get extracts the InformerFactory from the context.
func Get(ctx context.Context) informers.SharedInformerFactory {
	untyped := ctx.Value(Key{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch k8s.io/client-go/informers.SharedInformerFactory from context.")
	}
	return untyped.(informers.SharedInformerFactory)
}