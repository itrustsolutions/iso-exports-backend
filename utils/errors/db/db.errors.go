package dberrors

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

// Map of PostgreSQL error codes (SQLSTATE) to custom DB errors
var pgCodeToDBError = map[string]error{
	// Constraint violations
	"23505": ErrUniqueViolation,     // unique_violation
	"23503": ErrForeignKeyViolation, // foreign_key_violation
	"23502": ErrNotNullViolation,    // not_null_violation
	"23514": ErrCheckViolation,      // check_violation
	"23P01": ErrExclusionViolation,  // exclusion_violation

	// Serialization / concurrency
	"40001": ErrSerializationFailure, // serialization_failure
	"40P01": ErrDeadlockDetected,     // deadlock_detected

	// Syntax / schema errors
	"42601": ErrSyntaxError,               // syntax_error
	"42P01": ErrUndefinedTable,            // undefined_table
	"42703": ErrUndefinedColumn,           // undefined_column
	"22P02": ErrInvalidTextRepresentation, // invalid_text_representation
	"22007": ErrInvalidDatetimeFormat,     // invalid_datetime_format
	"53000": ErrTooManyConnections,        // too_many_connections
	"22012": ErrDivisionByZero,            // division_by_zero
}

// MapPostgresError maps a pgx/pgconn error to your custom DB error.
func MapPostgresError(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if mappedErr, ok := pgCodeToDBError[pgErr.Code]; ok {
			return mappedErr
		}
		return ErrUnknown // fallback for unknown codes
	}
	return err // not a pg error, return as-is
}

// GetPgError returns the underlying *pgconn.PgError if the error is a PostgreSQL error, or nil otherwise.
func GetPgError(err error) *pgconn.PgError {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr
	}
	return nil
}
