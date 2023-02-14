package fetcher

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
)

func getDatasetFrom(ctx context.Context) (string, error) {
	fieldContext := graphql.GetFieldContext(ctx)
	if fieldContext != nil && len(fieldContext.Path()) > 0 {
		dataset := fieldContext.Path().String()
		return dataset, nil
	}
	return "", errors.New("field context does not exist in resolver")
}
