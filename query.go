package morm

import (
	"context"
	"fmt"
	"strings"

	"github.com/mingmxren/morm/condition"
	"github.com/mingmxren/morm/dbtag"
	"github.com/mingmxren/morm/errs"
	"github.com/mingmxren/morm/limit"
	"github.com/mingmxren/morm/options/query"
	"github.com/mingmxren/morm/structreflect"
)

func QuestionQuerySqlArgs(p PersistentObject, opts ...query.Option) (sql string, args []interface{}, err error) {
	qo := &query.Options{}
	for _, opt := range opts {
		opt(qo)
	}
	where := ""
	args = make([]interface{}, 0)
	if qo.Condition != nil {
		where = fmt.Sprintf("WHERE %s", qo.Condition.Sql())
		args = qo.Condition.Args()
	}
	forUpdate := ""
	if qo.ForUpdate {
		forUpdate = "FOR UPDATE"
	}

	fields, _, err := dbtag.NameValuesWithTargetOrIgnore(p, qo.Fields, qo.IgnoreFields)
	if err != nil {
		return "", nil, err
	}

	sql = fmt.Sprintf("SELECT %s FROM %s %s %s %s %s",
		strings.Join(fields, ","), p.TableName(), where, qo.OrderBy, qo.Limit, forUpdate)

	return sql, args, nil
}

func QueryByPrimaryKey(ctx context.Context, db QueryxClient, p PersistentObjectWithPrimaryKey,
	opts ...query.Option) (found bool, err error) {
	pkNames, pkValues, err := dbtag.NameValuesByFieldPtr(p, p.PrimaryKey()...)
	if err != nil {
		return false, err
	}

	cond := make([]condition.Condition, 0, len(pkNames))
	for i, name := range pkNames {
		cond = append(cond, condition.Equal(name, pkValues[i]))
	}
	newOpts := make([]query.Option, 0)
	newOpts = append(newOpts,
		query.WithCondition(condition.And(cond...)),
		query.WithLimit(limit.LimitBy(2)),
	)
	newOpts = append(newOpts, opts...)

	ps, err := Query(ctx, db, p, newOpts...)
	if err != nil {
		return false, err
	}

	if len(ps) == 0 {
		return false, nil
	}

	if len(ps) > 1 {
		return true, errs.QueryOneGotMultiRow
	}

	structreflect.Copy(p, ps[0])

	return true, nil
}

func Query(ctx context.Context, db QueryxClient, p PersistentObject, opts ...query.Option) ([]PersistentObject, error) {
	sql, args, err := QuestionQuerySqlArgs(p, opts...)
	if err != nil {
		return nil, err
	}

	rows, err := db.QueryxContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ps := make([]PersistentObject, 0)
	for rows.Next() {
		np := structreflect.NewSameStruct(p).(PersistentObject)
		err := rows.StructScan(np)
		if err != nil {
			return nil, err
		}
		ps = append(ps, np)
	}
	return ps, nil
}
