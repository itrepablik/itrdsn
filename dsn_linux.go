// +build !linux

// Package itrdsn is the small wrapper to get the hard disk serial number.
package itrdsn

import (
	"github.com/shirou/gopsutil/disk"
)

// GetDiskSerialNo gets the hard disk serial number.
func GetDiskSerialNo() (string, error) {
	return disk.GetDiskSerialNumber("/dev/sda"), nil
}
