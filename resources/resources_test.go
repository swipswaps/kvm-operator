package resources

import (
	//"reflect"
	"log"
	"testing"

	"github.com/giantswarm/clusterspec"
)

func TestComponent_Creation(t *testing.T) {
	// Master: Service, Ingress8080, Ingress2379, Ingress6443, Deployment
	// Flannel-client: Deployment
	// Worker: Deployment, service
	expectedObjects := 8

	cluster := &clusterspec.Cluster{
		Spec: clusterspec.ClusterSpec{
			Customer:  "test",
			ClusterID: "test",
			Replicas:  int32(1),
		},
	}
	objects, err := ComputeResources(cluster)
	if err != nil {
		t.Fatalf("Error when computing cluster resources")
	}

	for _, obj := range objects {
		log.Println("obj desired resources for cluster: %v", obj)
	}

	if len(objects) != expectedObjects {
		t.Fatalf("Number of objects in expected output differed from received units: %d != %d", len(objects), expectedObjects)
	}
}
