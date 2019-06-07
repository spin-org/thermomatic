// Package imei implements an IMEI decoder.
package imei

// NOTE: for more information about IMEI codes and their structure you may
// consult with:
//
// https://en.wikipedia.org/wiki/International_Mobile_Equipment_Identity.

import (
	"errors"

	"github.com/spin-org/thermomatic/internal/common"
)

var (
	ErrInvalid  = errors.New("imei: invalid ")
	ErrChecksum = errors.New("imei: invalid checksum")
)

// Decode returns the IMEI code contained in the first 15 bytes of b.
//
// In case b isn't strictly composed of digits, the returned error will be
// ErrInvalid.
//
// In case b's checksum is wrong, the returned error will be ErrChecksum.
//
// Decode does NOT allocate under any condition. Additionally, it panics if b
// isn't at least 15 bytes long.
func Decode(b []byte) (code uint64, err error) {
	panic(common.ErrNotImplemented)
}
