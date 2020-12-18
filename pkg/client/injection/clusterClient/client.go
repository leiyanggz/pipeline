package clusterClient

import (
	context "context"
	"flag"
	"log"

	kubernetes "k8s.io/client-go/kubernetes"
	rest "k8s.io/client-go/rest"
	injection "knative.dev/pkg/injection"
	"knative.dev/pkg/injection/sharedmain"
	logging "knative.dev/pkg/logging"
)

func init() {
	injection.Default.RegisterClient(withClient)
}

// Key is used as the key for associating information w ith a context.Context.
type Key struct{}

func withClient(ctx context.Context, cfg *rest.Config) context.Context {
	clusterKubeconfig := flag.Lookup("cluster-kubeconfig").Value.String()
	clusterCfg, err := sharedmain.GetConfig("", clusterKubeconfig)
	if err != nil {
		log.Fatalf("Error building fcpKubeconfig: %v", err)
	}
	return context.WithValue(ctx, Key{}, kubernetes.NewForConfigOrDie(clusterCfg))
}

// Get extracts the kubernetes.Interface client from the context.
func Get(ctx context.Context) kubernetes.Interface {
	untyped := ctx.Value(Key{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch k8s.io/client-go/kubernetes.Interface from context.")
	}
	return untyped.(kubernetes.Interface)
}
