package orchestrator

import (
	"github.com/rancher/longhorn-manager/types"
)

type Request struct {
	NodeID   string
	Instance string

	VolumeName string
	VolumeSize int64

	ReplicaURLs []string
	RestoreFrom string
	RestoreName string
}

type Orchestrator interface {
	StartController(request *Request) (*Instance, error)
	StartReplica(request *Request) (*Instance, error)
	CleanupReplica(request *Request) error

	StopInstance(request *Request) (*Instance, error)
	InspectInstance(request *Request) (*Instance, error)

	GetCurrentNode() *types.NodeInfo
}

type Instance struct {
	Name    string
	NodeID  string
	Running bool
	IP      string
}
