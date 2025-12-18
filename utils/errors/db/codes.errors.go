package customdberrors

import "errors"

// Custom DB errors (generic, not business-specific or db-specific)
var (
	ErrDBUniqueViolation           = errors.New("unique constraint violation")
	ErrDBForeignKeyViolation       = errors.New("foreign key constraint violation")
	ErrDBNotNullViolation          = errors.New("not null constraint violation")
	ErrDBCheckViolation            = errors.New("check constraint violation")
	ErrDBExclusionViolation        = errors.New("exclusion constraint violation")
	ErrDBSerializationFailure      = errors.New("serialization failure")
	ErrDBDeadlockDetected          = errors.New("deadlock detected")
	ErrDBSyntaxError               = errors.New("syntax error")
	ErrDBUndefinedTable            = errors.New("undefined table")
	ErrDBUndefinedColumn           = errors.New("undefined column")
	ErrDBInvalidTextRepresentation = errors.New("invalid text representation")
	ErrDBInvalidDatetimeFormat     = errors.New("invalid datetime format")
	ErrDBTooManyConnections        = errors.New("too many connections")
	ErrDBDivisionByZero            = errors.New("division by zero")
	ErrDBUnknown                   = errors.New("unknown database error")
)
