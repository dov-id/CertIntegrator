package postgres

import sq "github.com/Masterminds/squirrel"

func BuildSelect(table string, columns ...string) sq.SelectBuilder {
	return sq.Select().Columns(columns...).From(table)
}

func ColumnWithAlias(query sq.SelectBuilder, expression sq.Sqlizer, alias string) sq.SelectBuilder {
	return query.Column(sq.Alias(expression, alias))
}

func GroupBy(query sq.SelectBuilder, columns ...string) sq.SelectBuilder {
	return query.GroupBy(columns...)
}

func BuildExpression(sql string, args ...interface{}) sq.Sqlizer {
	return sq.Expr(sql, args...)
}
