package types

import "github.com/jackc/pgx/v5"

type ListQueryParams struct {
	Limit  int `query:"limit" validate:"min=0"`
	Offset int `query:"offset" validate:"min=0"`
}

func (params *ListQueryParams) WrapQuery(query string) string {
	if params.Limit > 0 {
		return query + " limit @limit offset @offset"
	}

	return query
}

func (params *ListQueryParams) WrapNamedArgs(args pgx.NamedArgs) pgx.NamedArgs {
	args["limit"] = params.Limit
	args["offset"] = params.Offset

	return args
}

func NewDefaultListQueryParams() ListQueryParams {
	return ListQueryParams{
		Limit:  0, // ignored
		Offset: 0,
	}
}
