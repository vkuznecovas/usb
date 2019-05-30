// usb - Self contained USB and HID library for Go
// Copyright 2019 The library Authors
//
// This library is free software: you can redistribute it and/or modify it under
// the terms of the GNU Lesser General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option) any
// later version.
//
// The library is distributed in the hope that it will be useful, but WITHOUT ANY
// WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
// A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License along
// with the library. If not, see <http://www.gnu.org/licenses/>.

package usb

import (
	"C"
)

type GenericEndpointDirection uint8

// List of endpoint direction types
const (
	GenericEndpointDirectionOut = 0x00
	GenericEndpointDirectionIn  = 0x80
)

// List of endpoint attributes
const (
	GenericEndpointAttributeInterrupt = 3
)

// GenericEndpoint represents a USB endpoint
type GenericEndpoint struct {
	Address    uint8
	Direction  GenericEndpointDirection
	Attributes uint8
}

type GenericDeviceInfo struct {
	Path      string // Platform-specific device path
	VendorID  uint16 // Device Vendor ID
	ProductID uint16 // Device Product ID

	device *GenericDevice

	Interface int

	Endpoints []GenericEndpoint
}

func (gdi *GenericDeviceInfo) Type() DeviceType {
	return DeviceTypeGeneric
}

// Platform-specific device path
func (gdi *GenericDeviceInfo) GetPath() string {
	return gdi.Path
}

// IDs returns the vendor and product IDs for the device
func (gdi *GenericDeviceInfo) IDs() (uint16, uint16, int, uint16) {
	return gdi.VendorID, gdi.ProductID, gdi.Interface, 0
}