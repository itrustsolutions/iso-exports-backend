package dberrors

import "errors"

// Custom DB errors (generic, not business-specific or db-specific)
var (
	ErrUniqueViolation           = errors.New("unique constraint violation")
	ErrForeignKeyViolation       = errors.New("foreign key constraint violation")
	ErrNotNullViolation          = errors.New("not null constraint violation")
	ErrCheckViolation            = errors.New("check constraint violation")
	ErrExclusionViolation        = errors.New("exclusion constraint violation")
	ErrSerializationFailure      = errors.New("serialization failure")
	ErrDeadlockDetected          = errors.New("deadlock detected")
	ErrSyntaxError               = errors.New("syntax error")
	ErrUndefinedTable            = errors.New("undefined table")
	ErrUndefinedColumn           = errors.New("undefined column")
	ErrInvalidTextRepresentation = errors.New("invalid text representation")
	ErrInvalidDatetimeFormat     = errors.New("invalid datetime format")
	ErrTooManyConnections        = errors.New("too many connections")
	ErrDivisionByZero            = errors.New("division by zero")
	ErrUnknown                   = errors.New("unknown database error")
)
