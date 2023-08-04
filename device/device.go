package device

import (
	"github.com/luckyComet55/some-go-stuff/uuid"
)

type Device struct {
	status Status
	id     uuid.Uuid
}

func (d *Device) String() string {
	return "DEV.INSTANCE " + d.id.String() + " STATUS " + d.status.String()
}

func NewDevice() *Device {
	device := &Device{
		status: CONNECTING,
		id:     uuid.NewUuid(),
	}
	return device
}
