//go:build !linux

package device

import (
	"github.com/kshangx/wireguard-go/conn"
	"github.com/kshangx/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
