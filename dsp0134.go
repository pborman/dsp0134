// Package dsp0134 is a wrapper around github.com/google/uuid.  It supports the
// non-standard UUID format described in the DMTF System Management BIOS
// (SMBIOS) Reference Specification document DSP0134.
// https://www.dmtf.org/sites/default/files/standards/documents/DSP0134_3.4.0.pdf
// The DSP0134 standard reorders the first 8 bytes of the UUID.  The standard
// RFC4122 encoding for the UUID "00112233-4455-6677-8899-AABBCCDDEEFF" is
//	00 11 22 33 44 55 66 77 88 99 AA BB CC DD EE FF
// The encoding specified in DSP0134 is:
//	33 22 11 00 55 44 77 66 88 99 AA BB CC DD EE FF
package dsp0134

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

// A UUID is a 128 bit (16 byte) Universal Unique IDentifier as defined by
// DSP0134.
//
// Most methods from github.com/google/uuid can be used by invoking the
// UUID method:
//
//	var u UUID
//	return u.UUID().Version()
type UUID [16]byte

// swapUUID returns u after converting it to/from RFC4122 from/to DSP0134
// ordering.
func swapUUID(u [16]byte) [16]byte {
	u[0], u[1], u[2], u[3] = u[3], u[2], u[1], u[0]
	u[4], u[5] = u[5], u[4]
	u[6], u[7] = u[7], u[6]
	return u
}

// swapInPlace is like swapUUID but swaps u in place.
func swapInPlace(u []byte) {
	u[0], u[1], u[2], u[3] = u[3], u[2], u[1], u[0]
	u[4], u[5] = u[5], u[4]
	u[6], u[7] = u[7], u[6]
}

// ToUUID return u as a uuid.UUID.
func ToUUID(u UUID) uuid.UUID {
	return swapUUID(u)
}

// FromUUID return u as a UUID.
func FromUUID(u uuid.UUID) UUID {
	return swapUUID(u)
}

// Parse is analgous to github.com/google/uuid.Parse.
func Parse(s string) (UUID, error) {
	u, err := uuid.Parse(s)
	return swapUUID(u), err
}

// FromBytes is analgous to github.com/google/uuid.FromBytes.
func FromBytes(b []byte) (UUID, error) {
	u, err := uuid.FromBytes(b)
	return UUID(u), err
}

// ParseBytes is analgous to github.com/google/uuid.ParseBytes.
func ParseBytes(b []byte) (UUID, error) {
	u, err := uuid.ParseBytes(b)
	return FromUUID(u), err
}

// String returns the string form of uuid, xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx.
func (u UUID) String() string {
	return u.UUID().String()
}

// UUID returns u as a github.com/google/uuid.UUID.
func (u UUID) UUID() uuid.UUID {
	return ToUUID(u)
}

// MarshalText implements encoding.TextMarshaler.
func (u UUID) MarshalText() ([]byte, error) {
	return uuid.UUID(swapUUID(u)).MarshalText()
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (u *UUID) UnmarshalText(data []byte) error {
	if err := (*uuid.UUID)(u).UnmarshalText(data); err != nil {
		return err
	}
	swapInPlace(u[:])
	return nil
}

// MarshalBinary implements encoding.BinaryMarshaler
func (u UUID) MarshalBinary() ([]byte, error) {
	return uuid.UUID(u).MarshalBinary()
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (u *UUID) UnmarshalBinary(data []byte) error {
	return (*uuid.UUID)(u).UnmarshalBinary(data)
}

// Scan is analgous to github.com/google/uuid.Scan.
func (u *UUID) Scan(src interface{}) error {
	if err := (*uuid.UUID)(u).Scan(src); err != nil {
		return err
	}
	swapInPlace(u[:])
	return nil
}

// Value is analgous to github.com/google/uuid.Value.
func (u UUID) Value() (driver.Value, error) {
	return u.UUID().Value()
}
