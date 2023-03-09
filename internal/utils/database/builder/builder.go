package builder

import (
	"fmt"
	"strings"

	// "gitlab.com/hmcorp/wallet-houston/internal/utils/filtering"
)

const (
	LeftJoin   string = "LEFT JOIN"
	RightJoin  string = "RIGHT JOIN"
	InnerJoin  string = "INNER JOIN"
	LikeBefore string = "BEFORE"
	LikeAfter  string = "AFTER"
	LikeBoth   string = "BOTH"
)

type QueryBuilder struct {
	table                    string
	primaryKey               string
	excludeSoftDelete        bool
	groupStatement           string
	orderByStatement         string
	limitStatement           string
	limitWithOffsetStatement string
	offsetStatement					 string
	selectColumns            string
	rawJoinStatements        *RawJoinStatement
	whereExpression          []*Expression
	joinStatements           []*JoinStatement
	setStatements            []*SetStatement
	havingExpression         []*Expression
}

type Expression struct {
	Operation       string
	QueryExpression string
	Params          []interface{}
}

type WhereStatement struct {
	Statement string
	Params    []interface{}
}

type SetStatement struct {
	Field string
	Value interface{}
}

type JoinStatement struct {
	JoinTable string
	Condition string
	JoinType  string
}

type RawJoinStatement struct {
	JoinQuery string
	Params    []interface{}
}

func NewBuilder(table string) *QueryBuilder {
	return &QueryBuilder{
		table: table,
	}
}

func NewExpression(column, expresion string, param interface{}) *Expression {
	return &Expression{QueryExpression: fmt.Sprintf("%s %s %s", column, expresion, "?"), Params: []interface{}{param}}
}

func BuildPlaceholder(length int) string {
	if length <= 0 {
		return ""
	}

	placeholders := []string{}
	for i := 0; i < length; i++ {
		placeholders = append(placeholders, "?")
	}

	return strings.Join(placeholders, ",")
}

func GenerateStatement(groupedOperation, operation string, expressions ...*Expression) *Expression {
	var (
		query   = ""
		queries = []string{}
		params  = []interface{}{}
	)

	if len(expressions) == 0 {
		return nil
	}

	for _, stmt := range expressions {
		queries = append(queries, stmt.QueryExpression)
		params = append(params, stmt.Params...)
	}

	query = strings.Join(queries, fmt.Sprintf(" %s ", operation))
	if len(expressions) > 1 {
		query = fmt.Sprintf("( %s )", query)
	}

	return &Expression{
		Operation:       groupedOperation,
		QueryExpression: query,
		Params:          params,
	}
}

func Where(expressions ...*Expression) *Expression {
	return GenerateStatement("AND", "AND", expressions...)
}

func OrWhere(expressions ...*Expression) *Expression {
	return GenerateStatement("OR", "AND", expressions...)
}

func WhereAndOr(expressions ...*Expression) *Expression {
	return GenerateStatement("AND", "OR", expressions...)
}

func Like(column, expresion string, param interface{}) *Expression {
	var (
		exp = "CONCAT('%', ?, '%')"
	)

	switch expresion {
	case LikeAfter:
		exp = "CONCAT(?, '%')"
	case LikeBefore:
		exp = "CONCAT('%', ?)"
	}

	return &Expression{
		Operation:       "AND",
		QueryExpression: fmt.Sprintf("%s LIKE %s", column, exp),
		Params:          []interface{}{param},
	}
}

func OrLike(column, expresion string, param interface{}) *Expression {
	var (
		exp = "CONCAT('%', ?, '%')"
	)

	switch expresion {
	case LikeAfter:
		exp = "CONCAT(?, '%')"
	case LikeBefore:
		exp = "CONCAT('%', ?)"
	}

	return &Expression{
		Operation:       "OR",
		QueryExpression: fmt.Sprintf("%s LIKE %s", column, exp),
		Params:          []interface{}{param},
	}
}

func WhereIn(column string, params []interface{}) *Expression {
	placeholders := BuildPlaceholder(len(params))
	return &Expression{
		Operation:       "AND",
		QueryExpression: fmt.Sprintf("%s IN (%s)", column, placeholders),
		Params:          params,
	}
}

func WhereBetween(column string, params []interface{}) *Expression {
	return &Expression{
		Operation:       "AND",
		QueryExpression: fmt.Sprintf("%s BETWEEN (?) AND (?)", column),
		Params:          params,
	}
}

func RawWhereIn(column string, params string) *Expression {
	return &Expression{
		Operation:       "AND",
		QueryExpression: fmt.Sprintf("%s IN (%s)", column, params),
		Params:          []interface{}{},
	}
}

func WhereNotIn(column string, params []interface{}) *Expression {
	placeholders := BuildPlaceholder(len(params))
	return &Expression{
		Operation:       "AND",
		QueryExpression: fmt.Sprintf("%s NOT IN (%s)", column, placeholders),
		Params:          params,
	}
}

func RawWhereNotIn(column string, params string) *Expression {
	return &Expression{
		Operation:       "AND",
		QueryExpression: fmt.Sprintf("%s NOT IN (%s)", column, params),
		Params:          []interface{}{},
	}
}

func OrWhereIn(column string, params []interface{}) *Expression {
	placeholders := BuildPlaceholder(len(params))
	return &Expression{
		Operation:       "OR",
		QueryExpression: fmt.Sprintf("%s IN (%s)", column, placeholders),
		Params:          params,
	}
}

func RawOrWhereIn(column string, params string) *Expression {
	return &Expression{
		Operation:       "OR",
		QueryExpression: fmt.Sprintf("%s IN (%s)", column, params),
		Params:          []interface{}{},
	}
}

func OrWhereNotIn(column string, params []interface{}) *Expression {
	placeholders := BuildPlaceholder(len(params))
	return &Expression{
		Operation:       "OR",
		QueryExpression: fmt.Sprintf("%s NOT IN (%s)", column, placeholders),
		Params:          params,
	}
}

func RawOrWhereNotIn(column string, params string) *Expression {
	return &Expression{
		Operation:       "OR",
		QueryExpression: fmt.Sprintf("%s NOT IN (%s)", column, params),
		Params:          []interface{}{},
	}
}

func WhereIsNull(column string) *Expression {
	return &Expression{
		Operation:       "AND",
		QueryExpression: fmt.Sprintf("%s IS NULL", column),
		Params:          []interface{}{},
	}
}

func WhereIsNotNull(column string) *Expression {
	return &Expression{
		Operation:       "AND",
		QueryExpression: fmt.Sprintf("%s IS NOT NULL", column),
		Params:          []interface{}{},
	}
}

func OrWhereIsNull(column string) *Expression {
	return &Expression{
		Operation:       "OR",
		QueryExpression: fmt.Sprintf("%s IS NULL", column),
		Params:          []interface{}{},
	}
}

func OrWhereIsNotNull(column string) *Expression {
	return &Expression{
		Operation:       "OR",
		QueryExpression: fmt.Sprintf("%s IS NOT NULL", column),
		Params:          []interface{}{},
	}
}

func (qb *QueryBuilder) Page(currentPage, perPage int) *QueryBuilder {
	currentPage = (currentPage * perPage) - perPage
	qb.limitStatement = fmt.Sprintf("LIMIT %d, %d", currentPage, perPage)
	return qb
}

func (qb *QueryBuilder) Limit(limit int) *QueryBuilder {
	if limit > 0 {
		qb.limitStatement = fmt.Sprintf("LIMIT %d", limit)
	}
	return qb
}

func (qb *QueryBuilder) Offset(offset int) *QueryBuilder {
	if offset > 0 {
		qb.offsetStatement = fmt.Sprintf("OFFSET %d", offset)
	}
	return qb
}

func (qb *QueryBuilder) LimitWithOffset(limit, offset int) *QueryBuilder {
	qb.limitWithOffsetStatement = fmt.Sprintf("LIMIT %d OFFSET %d", limit, offset)
	return qb
}

func (qb *QueryBuilder) OrderBy(fields ...string) *QueryBuilder {
	qb.orderByStatement = fmt.Sprintf("ORDER BY %s ASC", strings.Join(fields, ","))
	return qb
}

func (qb *QueryBuilder) RawOrderBy(query string) *QueryBuilder {
	qb.orderByStatement = query
	return qb
}

func (qb *QueryBuilder) OrderByDesc(fields ...string) *QueryBuilder {
	qb.orderByStatement = fmt.Sprintf("ORDER BY %s DESC", strings.Join(fields, ","))
	return qb
}

func (qb *QueryBuilder) GroupBy(fields ...string) *QueryBuilder {
	qb.groupStatement = fmt.Sprintf("GROUP BY %s", strings.Join(fields, ","))
	return qb
}

func (qb *QueryBuilder) Condition(expressions ...*Expression) *QueryBuilder {
	qb.whereExpression = append(qb.whereExpression, expressions...)
	return qb
}

func (qb *QueryBuilder) BuildWhereStatement() *WhereStatement {
	var (
		whereQuery = ""
		params     = []interface{}{}
	)

	if len(qb.whereExpression) == 0 {
		return nil
	}

	for idx, exp := range qb.whereExpression {
		if idx == 0 {
			whereQuery = fmt.Sprintf("%s %s", whereQuery, exp.QueryExpression)
		} else {
			whereQuery = fmt.Sprintf("%s %s %s", whereQuery, exp.Operation, exp.QueryExpression)
		}

		params = append(params, exp.Params...)
	}

	result := &WhereStatement{
		Statement: whereQuery,
		Params:    params,
	}

	return result
}

func (qb *QueryBuilder) Having(expressions ...*Expression) *QueryBuilder {
	qb.havingExpression = append(qb.havingExpression, expressions...)
	return qb
}

func (qb *QueryBuilder) BuildHavingStatement() *WhereStatement {
	var (
		whereQuery = ""
		params     = []interface{}{}
	)

	if len(qb.havingExpression) == 0 {
		return nil
	}

	for idx, exp := range qb.havingExpression {
		if idx == 0 {
			whereQuery = fmt.Sprintf("%s %s", whereQuery, exp.QueryExpression)
		} else {
			whereQuery = fmt.Sprintf("%s %s %s", whereQuery, exp.Operation, exp.QueryExpression)
		}

		params = append(params, exp.Params...)
	}

	result := &WhereStatement{
		Statement: whereQuery,
		Params:    params,
	}

	return result
}

func (qb *QueryBuilder) PrimaryKey(pk string) *QueryBuilder {
	qb.primaryKey = pk
	return qb
}

func (qb *QueryBuilder) Select(columns ...string) *QueryBuilder {
	qb.selectColumns = strings.Join(columns, ",")
	return qb
}

func (qb *QueryBuilder) SelectDistinct(columns ...string) *QueryBuilder {
	qb.selectColumns = "DISTINCT " + strings.Join(columns, ",")
	return qb
}

func (qb *QueryBuilder) BuildJoinStatement() (string, []interface{}) {
	query := ""
	params := []interface{}{}

	if len(qb.joinStatements) > 0 {
		for _, join := range qb.joinStatements {
			joinStmt := fmt.Sprintf("%s %s ON %s", join.JoinType, join.JoinTable, join.Condition)
			query = fmt.Sprintf(" %s %s ", query, joinStmt)
		}
	}

	if qb.rawJoinStatements != nil {
		query = fmt.Sprintf(" %s %s ", query, qb.rawJoinStatements.JoinQuery)
		params = append(params, qb.rawJoinStatements.Params...)
	}

	return query, params
}

func (qb *QueryBuilder) Join(joinStatements ...*JoinStatement) *QueryBuilder {
	qb.joinStatements = joinStatements
	return qb
}

func (qb *QueryBuilder) RawJoinWithParams(joinStatement string, params []interface{}) *QueryBuilder {
	if qb.rawJoinStatements == nil {
		qb.rawJoinStatements = &RawJoinStatement{
			JoinQuery: joinStatement,
			Params:    params,
		}
	} else {
		qb.rawJoinStatements.JoinQuery = fmt.Sprintf("%s %s", qb.rawJoinStatements.JoinQuery, joinStatement)
		qb.rawJoinStatements.Params = append(qb.rawJoinStatements.Params, params...)
	}

	return qb
}

func (qb *QueryBuilder) SetData(setStatement ...*SetStatement) *QueryBuilder {
	qb.setStatements = setStatement
	return qb
}

func (qb *QueryBuilder) ExcludeDeleted() *QueryBuilder {
	qb.excludeSoftDelete = true
	return qb
}

func (qb *QueryBuilder) Find(ID int64) (populateQuery string, params []interface{}) {
	var (
		query                = "SELECT"
		selectColumns        = "*"
		primaryKey           = "id"
		joinStmt, joinParams = qb.BuildJoinStatement()
	)

	if qb.selectColumns != "" {
		selectColumns = qb.selectColumns
	}

	if qb.primaryKey != "" {
		primaryKey = qb.primaryKey
	}

	query = fmt.Sprintf("%s %s FROM %s %s WHERE %s=?", query, selectColumns, qb.table, joinStmt, primaryKey)
	params = append(params, joinParams...)
	params = append(params, ID)
	return query, params
}

func (qb *QueryBuilder) Count() (populateQuery string, params []interface{}) {
	var (
		query                = "SELECT COUNT(*) as total"
		joinStmt, joinParams = qb.BuildJoinStatement()
		whereStmt            = qb.BuildWhereStatement()
		havingStmt           = qb.BuildHavingStatement()
	)

	query = fmt.Sprintf("%s FROM %s %s", query, qb.table, joinStmt)
	params = append(params, joinParams...)

	// where statement
	if whereStmt != nil && whereStmt.Statement != "" {
		query = fmt.Sprintf("%s WHERE %s", query, whereStmt.Statement)
	}

	if whereStmt != nil && len(whereStmt.Params) > 0 {
		params = append(params, whereStmt.Params...)
	}
	// end where statement

	// group by statement
	if qb.groupStatement != "" {
		query = fmt.Sprintf("%s %s", query, qb.groupStatement)
	}

	// having statement
	if havingStmt != nil && havingStmt.Statement != "" {
		query = fmt.Sprintf("%s HAVING %s", query, havingStmt.Statement)
	}

	if havingStmt != nil && len(havingStmt.Params) > 0 {
		params = append(params, havingStmt.Params...)
	}
	// end having statement

	return query, params
}

func (qb *QueryBuilder) First() (populateQuery string, params []interface{}) {
	var (
		query                = "SELECT"
		selectColumns        = "*"
		joinStmt, joinParams = qb.BuildJoinStatement()
		whereStmt            = qb.BuildWhereStatement()
		havingStmt           = qb.BuildHavingStatement()
	)

	if qb.selectColumns != "" {
		selectColumns = qb.selectColumns
	}

	// base query statement
	query = fmt.Sprintf("%s %s FROM %s %s", query, selectColumns, qb.table, joinStmt)
	params = append(params, joinParams...)

	// where statement
	if whereStmt != nil && whereStmt.Statement != "" {
		query = fmt.Sprintf("%s WHERE %s", query, whereStmt.Statement)
	}

	if whereStmt != nil && len(whereStmt.Params) > 0 {
		params = append(params, whereStmt.Params...)
	}
	// end where statement

	// group by statement
	if qb.groupStatement != "" {
		query = fmt.Sprintf("%s %s", query, qb.groupStatement)
	}

	// having statement
	if havingStmt != nil && havingStmt.Statement != "" {
		query = fmt.Sprintf("%s HAVING %s", query, havingStmt.Statement)
	}

	if havingStmt != nil && len(havingStmt.Params) > 0 {
		params = append(params, havingStmt.Params...)
	}
	// end having statement

	// order by statement
	if qb.orderByStatement != "" {
		query = fmt.Sprintf("%s %s", query, qb.orderByStatement)
	}

	// limit statement
	query = fmt.Sprintf("%s LIMIT 1", query)
	return query, params
}

func (qb *QueryBuilder) Get() (populateQuery string, params []interface{}) {
	var (
		query                = "SELECT"
		selectColumns        = "*"
		joinStmt, joinParams = qb.BuildJoinStatement()
		whereStmt            = qb.BuildWhereStatement()
		havingStmt           = qb.BuildHavingStatement()
	)

	if qb.selectColumns != "" {
		selectColumns = qb.selectColumns
	}

	// base query statement
	query = fmt.Sprintf("%s %s FROM %s %s", query, selectColumns, qb.table, joinStmt)
	params = append(params, joinParams...)

	// where statement
	if whereStmt != nil && whereStmt.Statement != "" {
		query = fmt.Sprintf("%s WHERE %s", query, whereStmt.Statement)
	}

	if whereStmt != nil && len(whereStmt.Params) > 0 {
		params = append(params, whereStmt.Params...)
	}
	// end where statement

	// group by statement
	if qb.groupStatement != "" {
		query = fmt.Sprintf("%s %s", query, qb.groupStatement)
	}

	// having statement
	if havingStmt != nil && havingStmt.Statement != "" {
		query = fmt.Sprintf("%s HAVING %s", query, havingStmt.Statement)
	}

	if havingStmt != nil && len(havingStmt.Params) > 0 {
		params = append(params, havingStmt.Params...)
	}
	// end having statement

	// order by statement
	if qb.orderByStatement != "" {
		query = fmt.Sprintf("%s %s", query, qb.orderByStatement)
	}

	// limit/pagging statement
	if qb.limitStatement != "" {
		query = fmt.Sprintf("%s %s", query, qb.limitStatement)
	}

	// limit with offset statement
	if qb.limitWithOffsetStatement != "" {
		query = fmt.Sprintf("%s %s", query, qb.limitWithOffsetStatement)
	}

	return query, params
}

func (qb *QueryBuilder) Insert() (populateQuery string, params []interface{}) {
	var (
		query    = "INSERT INTO"
		setStmts = []string{}
	)

	if len(qb.setStatements) == 0 {
		return "", nil
	}

	for _, set := range qb.setStatements {
		setStmts = append(setStmts, fmt.Sprintf("%s=?", set.Field))
		params = append(params, set.Value)
	}

	

	query = fmt.Sprintf("%s %s SET %s", query, qb.table, strings.Join(setStmts, ","))
	return query, params
}

func (qb *QueryBuilder) BatchInsert(data []map[string]interface{}, columns []string) (populateQuery string, params []interface{}) {
	var (
		query      = "INSERT INTO"
		insertStmt = []string{}
	)

	for _, dt := range data {
		placeholders := BuildPlaceholder(len(columns))
		insertStmt = append(insertStmt, fmt.Sprintf("(%s)", placeholders))
		paramsPool := make([]interface{}, len(columns))

		for idx, cl := range columns {
			paramsPool[idx] = dt[cl]
		}

		params = append(params, paramsPool...)
	}

	query = fmt.Sprintf("%s %s (%s) VALUES %s", query, qb.table, strings.Join(columns, ","), strings.Join(insertStmt, ","))
	return query, params
}

func (qb *QueryBuilder) Update() (populateQuery string, params []interface{}) {
	var (
		query     = "UPDATE"
		setStmts  = []string{}
		whereStmt = qb.BuildWhereStatement()
	)

	if len(qb.setStatements) == 0 {
		return "", nil
	}

	for _, set := range qb.setStatements {
		setStmts = append(setStmts, fmt.Sprintf("%s=?", set.Field))
		params = append(params, set.Value)
	}

	// base query statement
	query = fmt.Sprintf("%s %s SET %s", query, qb.table, strings.Join(setStmts, ","))

	// where statement
	if whereStmt != nil && whereStmt.Statement != "" {
		query = fmt.Sprintf("%s WHERE %s", query, whereStmt.Statement)
	}

	if whereStmt != nil && len(whereStmt.Params) > 0 {
		params = append(params, whereStmt.Params...)
	}
	// end where statement

	return query, params
}

// permanent delete
func (qb *QueryBuilder) Delete() (populateQuery string, params []interface{}) {
	var (
		query     = "DELETE FROM"
		whereStmt = qb.BuildWhereStatement()
	)

	// prevent unexpected delete
	if whereStmt == nil || (whereStmt != nil && whereStmt.Statement == "") {
		return "", params
	}

	// base query statement
	query = fmt.Sprintf("%s %s WHERE %s", query, qb.table, whereStmt.Statement)
	params = append(params, whereStmt.Params...)

	return query, params
}
