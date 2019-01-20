package driver

import "fmt"

// ErrUnsupportedAllocation is returned when a request is asking to be allocated an IP.
type ErrUnsupportedAllocation struct{}

func (e ErrUnsupportedAllocation) Error() string {
	return "allocation of IP pools is not supported"
}

// BadRequest denotes the type of this error.
func (e ErrUnsupportedAllocation) BadRequest() {}

// ErrInvalidAddrSpace is returned when an unrecognized address space is requested.
type ErrInvalidAddrSpace string

func (e ErrInvalidAddrSpace) Error() string {
	return fmt.Sprintf("invalid address space: %s", e)
}

// BadRequest denotes the type of this error.
func (e ErrInvalidAddrSpace) BadRequest() {}

// ErrParsePool is returned when a pool ID failed to parse as a CIDR subnet address.
type ErrParsePool string

func (e ErrParsePool) Error() string {
	return fmt.Sprintf("failed to parse pool as CIDR: %s", e)
}

// BadRequest denotes the type of this error.
func (e ErrParsePool) BadRequest() {}

// ErrParseIP is returned when a given address fails to parse as IP.
type ErrParseIP string

func (e ErrParseIP) Error() string {
	return fmt.Sprintf("failed to parse as IP: %s", e)
}

// BadRequest denotes the type of this error.
func (e ErrParseIP) BadRequest() {}
