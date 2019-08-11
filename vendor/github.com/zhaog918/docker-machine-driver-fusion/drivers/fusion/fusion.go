// +build !darwin

package fusion

import "github.com/docker/machine/libmachine/drivers"

func NewDriver(hostName, storePath string) drivers.Driver {
	return drivers.NewDriverNotSupported("fusion", hostName, storePath)
}
