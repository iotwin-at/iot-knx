package knx

import "fmt"

// ------------------------------------------------------------------------------------
// ErrAddressResolution
// ------------------------------------------------------------------------------------
type ErrAddressResolution string

// Error implements error.
func (e ErrAddressResolution) Error() string {
	return fmt.Sprintf("ErrAddressResolution: %s", string(e))
}

func NewErrAddressResolution(format string, args ...any) error {
	return ErrAddressResolution(fmt.Sprintf(format, args...))
}

// ------------------------------------------------------------------------------------
// ErrNetConnection
// ------------------------------------------------------------------------------------
type ErrNetConnection string

// Error implements error.
func (e ErrNetConnection) Error() string {
	return fmt.Sprintf("ErrNetConnection: %s", string(e))
}

func NewErrNetConnection(format string, args ...any) error {
	return ErrNetConnection(fmt.Sprintf(format, args...))
}

// ------------------------------------------------------------------------------------
// ErrInvalidDataframe
// ------------------------------------------------------------------------------------
type ErrInvalidDataframe string

// Error implements error.
func (e ErrInvalidDataframe) Error() string {
	return fmt.Sprintf("ErrInvalidDataframe: %s", string(e))
}

func NewErrInvalidDataframe(format string, args ...any) error {
	return ErrInvalidDataframe(fmt.Sprintf(format, args...))
}

// ------------------------------------------------------------------------------------
// ErrInvalidDataType
// ------------------------------------------------------------------------------------
type ErrInvalidDataType string

// Error implements error.
func (e ErrInvalidDataType) Error() string {
	return fmt.Sprintf("ErrInvalidDataType: %s", string(e))
}

func NewErrInvalidDataType(format string, args ...any) error {
	return ErrInvalidDataType(fmt.Sprintf(format, args...))
}

// ------------------------------------------------------------------------------------
// ErrNotSupported
// ------------------------------------------------------------------------------------
type ErrNotSupported string

// Error implements error.
func (e ErrNotSupported) Error() string {
	return fmt.Sprintf("ErrNotSupported: %s", string(e))
}

func NewErErrNotSupported(format string, args ...any) error {
	return ErrNotSupported(fmt.Sprintf(format, args...))
}

// ------------------------------------------------------------------------------------
// ErrNotImplemented
// ------------------------------------------------------------------------------------
type ErrNotImplemented string

// Error implements error.
func (e ErrNotImplemented) Error() string {
	return fmt.Sprintf("ErrNotImplemented: %s", string(e))
}

func NewErrNotImplemented(format string, args ...any) error {
	return ErrNotImplemented(fmt.Sprintf(format, args...))
}
