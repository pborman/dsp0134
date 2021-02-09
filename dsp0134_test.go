package dsp0134

import (
	"testing"

	"github.com/google/uuid"
)

const sample = "00112233-4455-6677-8899-aabbccddeeff"

var (
	dsp0134 = [16]byte{
		0x33, 0x22, 0x11, 0x00, 0x55, 0x44, 0x77, 0x66,
		0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff,
	}
	rfc4122 = [16]byte{
		0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77,
		0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff,
	}
)

func TestSwapUUID(t *testing.T) {
	got := swapUUID(rfc4122)
	if got != dsp0134 {
		t.Errorf("Swap got %x, want %x", got[:], dsp0134)
	}
	if rfc4122 == dsp0134 {
		t.Fatalf("swap altered its input")
	}
	got = swapUUID(got)
	if got != rfc4122 {
		t.Errorf("Reswap got %s, want %x", got[:], rfc4122)
	}
	swapInPlace(got[:])
	if got != dsp0134 {
		t.Errorf("swapInPlace got %x, want %x", got[:], dsp0134)
	}
}

func TestFromTo(t *testing.T) {

	var u UUID = dsp0134
	if got := ToUUID(u); got != rfc4122 {
		t.Errorf("ToUUID got %x, want %x", got[:], rfc4122)
	}
	if got := u.UUID(); got != rfc4122 {
		t.Errorf("u.UUID got %x, want %x", got[:], rfc4122)
	}
	var gu uuid.UUID = rfc4122
	if got := FromUUID(gu); got != dsp0134 {
		t.Errorf("ToUUID got %x, want %x", got[:], dsp0134)
	}
	if got, err := FromBytes(dsp0134[:]); err != nil {
                t.Errorf("FromBytes got %v", err)
	} else if got != dsp0134 {
                t.Errorf("FromBytes %x, want %x", got[:], dsp0134)
        }
}

func TestParse(t *testing.T) {
	got, err := Parse(sample)
	if err != nil {
		t.Fatal(err)
	}
	want := UUID{
		0x33, 0x22, 0x11, 0x00, 0x55, 0x44, 0x77, 0x66,
		0x88, 0x99, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF}
	if got != want {
		t.Errorf("Got %x, want %x\n", got[:], want[:])
	}
}

func TestParseBytes(t *testing.T) {
	got, err := ParseBytes(([]byte)(sample))
	if err != nil {
		t.Fatal(err)
	}
	want := UUID{
		0x33, 0x22, 0x11, 0x00, 0x55, 0x44, 0x77, 0x66,
		0x88, 0x99, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF}
	if got != want {
		t.Errorf("Got %x, want %x\n", got[:], want[:])
	}
}

func TestMarshalText(t *testing.T) {
	var u UUID = dsp0134
	data, err := u.MarshalText()
	if err != nil {
		t.Fatal(err)
	}
	got := string(data)
	if got != sample {
		t.Errorf("got %q, want %q\n", got, sample)
	}
}

func TestUnmarshalText(t *testing.T) {
	var u UUID
	err := u.UnmarshalText([]byte(sample))
	if err != nil {
		t.Fatal(err)
	}
	if u != dsp0134 {
		t.Errorf("Got %x, want %x\n", u[:], dsp0134)
	}
	if err := u.UnmarshalText([]byte("wrong input")); err == nil {
		t.Errorf("Did not get expected error")
	}
}

func TestMarshalBinary(t *testing.T) {
	var u UUID = dsp0134
	data, err := u.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	if len(data) != 16 {
		t.Fatalf("MarshalBinary returned %d bytes", len(data))
	}
	var got UUID
	copy(got[:], data)
	if got != dsp0134 {
		t.Errorf("got %x, want %x", got[:], dsp0134)
	}
}

func TestUnmarshalBinary(t *testing.T) {
	var u UUID
	if err := u.UnmarshalBinary(dsp0134[:]); err != nil {
		t.Fatal(err)
	}
	if u != dsp0134 {
		t.Errorf("Got %x, want %x\n", u[:], dsp0134)
	}
	if err := u.UnmarshalBinary(dsp0134[:3]); err == nil {
		t.Errorf("Did not get expected error")
	}
}

func TestScan(t *testing.T) {
	// We only need to test an error and a success as the
	// underlying Scan is already tested.
	var u UUID
	if err := u.Scan("invalid-UUID"); err == nil {
		t.Errorf("Did not get error on invalid input")
	}
	if err := u.Scan(sample); err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if u != dsp0134 {
		t.Errorf("Got %x, want %x", u[:], dsp0134)
	}
}

func TestValue(t *testing.T) {
	var u UUID = dsp0134
	got, _ := u.Value()
	if got != sample {
		t.Errorf("Got %q, want %q", got, sample)
	}
}

func TestString(t *testing.T) {
	var u UUID = dsp0134
	if got := u.String(); got != sample {
		t.Errorf("Got %q, want %q", got, sample)
	}
	u = rfc4122
	if got := u.String(); got == sample {
		t.Errorf("Invalid string result.")
	}
}
