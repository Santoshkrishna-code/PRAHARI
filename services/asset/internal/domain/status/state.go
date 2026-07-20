package status

import (
	"fmt"
)

// Code defines state tags in walkthrough audits.
type Code string

const (
	CodeRegistered     Code = "REGISTERED"
	CodeCommissioned   Code = "COMMISSIONED"
	CodeOperational    Code = "OPERATIONAL"
	CodeMaintenance    Code = "MAINTENANCE"
	CodeOutOfService   Code = "OUT_OF_SERVICE"
	CodeDecommissioned Code = "DECOMMISSIONED"
	CodeDisposed       Code = "DISPOSED"
)

// ValidCodes lists all valid state codes.
var ValidCodes = []Code{
	CodeRegistered,
	CodeCommissioned,
	CodeOperational,
	CodeMaintenance,
	CodeOutOfService,
	CodeDecommissioned,
	CodeDisposed,
}

// Validate checks if code exists.
func (c Code) Validate() error {
	for _, valid := range ValidCodes {
		if c == valid {
			return nil
		}
	}
	return fmt.Errorf("invalid lifecycle code: %s", c)
}
