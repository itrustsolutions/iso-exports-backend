package utils

import (
	"context"

	db "github.com/itrustsolutions/iso-exports-backend/core/identity/internal/db/gen"
	customcontext "github.com/itrustsolutions/iso-exports-backend/utils/context"
)

func GetQueriesWithPossibleTx(ctx context.Context, baseQueries *db.Queries) *db.Queries {
	tx := customcontext.ExtractTx(ctx)
	if tx != nil {
		return baseQueries.WithTx(tx)
	}
	return baseQueries
}
