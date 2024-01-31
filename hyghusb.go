package hyghUsb

import (
	"fmt"
	"log"

	"github.com/google/gousb"
)

type UsbManager struct {
	context *gousb.Context
}

func NewUsbManager() (*UsbManager, error) {
	context := gousb.NewContext()
	return &UsbManager{context: context}, nil
}

func (um *UsbManager) Close() {
	um.context.Close()
}

func (um *UsbManager) ListDevices() ([]*gousb.Device, error) {
	devices, err := um.context.OpenDevices(nil, nil)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func (um *UsbManager) GetDeviceInformation(device *gousb.Device) (string, error) {
	desc := device.Desc
	info := fmt.Sprintf("Vendor ID: %s, Product ID: %s", desc.Vendor.String(), desc.Product.String())
	return info, nil
}

func (um *UsbManager) SendData(device *gousb.Device, endpoint gousb.Endpoint, data []byte) error {
	_, err := endpoint.Write(data)
	return err
}

func (um *UsbManager) ReceiveData(device *gousb.Device, endpoint gousb.Endpoint, size int) ([]byte, error) {
	data := make([]byte, size)
	_, err := endpoint.Read(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (um *UsbManager) DisableDevice(device *gousb.Device) error {
	err := device.Close()
	return err
}
