// Package itrdsn is the small wrapper to get the hard disk serial number.
package itrdsn

import (
	"github.com/jaypipes/ghw"
)

// GetDiskSerialNo gets the hard disk serial number.
func GetDiskSerialNo() (string, error) {
	block, err := ghw.Block()
	if err != nil {
		return "", err
	}
	serialNo := ""
	for _, disk := range block.Disks {
		serialNo = disk.SerialNumber
	}
	return serialNo, nil
}
