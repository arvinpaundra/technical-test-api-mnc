package dbutil

type QueryOption func(opt *QueryOptions)

type QueryOptions struct {
	Select   []string
	Preloads []preload
	Joins    []joinClause
	Where    []whereClause
	Group    string
	Order    string
	Omit     string
	Limit    int
	Offset   int
}

type preload struct {
	Query string
	Args  []any
}

type joinClause struct {
	Query string
	Args  []any
}

type whereClause struct {
	Query any
	Args  []any
}

func Select(query ...string) QueryOption {
	return func(opt *QueryOptions) {
		opt.Select = append(opt.Select, query...)
	}
}

func Preload(query string, args ...any) QueryOption {
	return func(opt *QueryOptions) {
		opt.Preloads = append(opt.Preloads, preload{query, args})
	}
}

func Joins(query string, args ...any) QueryOption {
	return func(opt *QueryOptions) {
		opt.Joins = append(opt.Joins, joinClause{query, args})
	}
}

func Where(query any, args ...any) QueryOption {
	return func(opt *QueryOptions) {
		opt.Where = append(opt.Where, whereClause{query, args})
	}
}

func Group(cols string) QueryOption {
	return func(opt *QueryOptions) {
		opt.Group = cols
	}
}

func Order(query string) QueryOption {
	return func(opt *QueryOptions) {
		opt.Order = query
	}
}

func Omit(cols string) QueryOption {
	return func(opt *QueryOptions) {
		opt.Omit = cols
	}
}

func Limit(limit int) QueryOption {
	return func(opt *QueryOptions) {
		opt.Limit = limit
	}
}

func Offset(offset int) QueryOption {
	return func(opt *QueryOptions) {
		opt.Offset = offset
	}
}
