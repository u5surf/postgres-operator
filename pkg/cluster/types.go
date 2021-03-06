package cluster

import (
	acidv1 "github.com/zalando-incubator/postgres-operator/pkg/apis/acid.zalan.do/v1"
	"k8s.io/api/apps/v1beta1"
	"k8s.io/api/core/v1"
	policybeta1 "k8s.io/api/policy/v1beta1"
	"k8s.io/apimachinery/pkg/types"
	"time"
)

// PostgresRole describes role of the node
type PostgresRole string

const (
	// Master role
	Master PostgresRole = "master"

	// Replica role
	Replica PostgresRole = "replica"
)

type PodEventType string

// Possible values for the EventType
const (
	PodEventAdd    PodEventType = "ADD"
	PodEventUpdate PodEventType = "UPDATE"
	PodEventDelete PodEventType = "DELETE"
)

// PodEvent describes the event for a single Pod
type PodEvent struct {
	ResourceVersion string
	PodName         types.NamespacedName
	PrevPod         *v1.Pod
	CurPod          *v1.Pod
	EventType       PodEventType
}

// Process describes process of the cluster
type Process struct {
	Name      string
	StartTime time.Time
}

// WorkerStatus describes status of the worker
type WorkerStatus struct {
	CurrentCluster types.NamespacedName
	CurrentProcess Process
}

// ClusterStatus describes status of the cluster
type ClusterStatus struct {
	Team                string
	Cluster             string
	MasterService       *v1.Service
	ReplicaService      *v1.Service
	MasterEndpoint      *v1.Endpoints
	ReplicaEndpoint     *v1.Endpoints
	StatefulSet         *v1beta1.StatefulSet
	PodDisruptionBudget *policybeta1.PodDisruptionBudget

	CurrentProcess Process
	Worker         uint32
	Status         acidv1.PostgresStatus
	Spec           acidv1.PostgresSpec
	Error          error
}
