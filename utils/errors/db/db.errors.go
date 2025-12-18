package customdberrors

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

// Map of PostgreSQL error codes (SQLSTATE) to custom DB errors
var pgCodeToDBError = map[string]error{
	// Constraint violations
	"23505": ErrDBUniqueViolation,     // unique_violation
	"23503": ErrDBForeignKeyViolation, // foreign_key_violation
	"23502": ErrDBNotNullViolation,    // not_null_violation
	"23514": ErrDBCheckViolation,      // check_violation
	"23P01": ErrDBExclusionViolation,  // exclusion_violation

	// Serialization / concurrency
	"40001": ErrDBSerializationFailure, // serialization_failure
	"40P01": ErrDBDeadlockDetected,     // deadlock_detected

	// Syntax / schema errors
	"42601": ErrDBSyntaxError,               // syntax_error
	"42P01": ErrDBUndefinedTable,            // undefined_table
	"42703": ErrDBUndefinedColumn,           // undefined_column
	"22P02": ErrDBInvalidTextRepresentation, // invalid_text_representation
	"22007": ErrDBInvalidDatetimeFormat,     // invalid_datetime_format
	"53000": ErrDBTooManyConnections,        // too_many_connections
	"22012": ErrDBDivisionByZero,            // division_by_zero
}

// MapPostgresError maps a pgx/pgconn error to your custom DB error.
func MapPostgresError(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if mappedErr, ok := pgCodeToDBError[pgErr.Code]; ok {
			return mappedErr
		}
		return ErrDBUnknown // fallback for unknown codes
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
