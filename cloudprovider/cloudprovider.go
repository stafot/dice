package cloudprovider

import "github.com/dmathieu/dice/kubernetes"

type CloudProvider interface {
	// Name returns name of the cloud provider.
	Name() string

	// Delete finds and deletes the specified node
	Delete(*kubernetes.Node) error

	// Refresh is called before every main loop and can be used to dynamically update cloud provider state.
	// In particular the list of node groups returned by NodeGroups can change as a result of CloudProvider.Refresh().
	Refresh() error
}