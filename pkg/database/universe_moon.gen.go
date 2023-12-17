// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package database

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/buckbrady/evebot-tools-backend/pkg/model"
)

func newUniverseMoon(db *gorm.DB, opts ...gen.DOOption) universeMoon {
	_universeMoon := universeMoon{}

	_universeMoon.universeMoonDo.UseDB(db, opts...)
	_universeMoon.universeMoonDo.UseModel(&model.UniverseMoon{})

	tableName := _universeMoon.universeMoonDo.TableName()
	_universeMoon.ALL = field.NewAsterisk(tableName)
	_universeMoon.ID = field.NewInt32(tableName, "id")
	_universeMoon.SystemID = field.NewInt32(tableName, "system_id")

	_universeMoon.fillFieldMap()

	return _universeMoon
}

type universeMoon struct {
	universeMoonDo universeMoonDo

	ALL      field.Asterisk
	ID       field.Int32
	SystemID field.Int32

	fieldMap map[string]field.Expr
}

func (u universeMoon) Table(newTableName string) *universeMoon {
	u.universeMoonDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u universeMoon) As(alias string) *universeMoon {
	u.universeMoonDo.DO = *(u.universeMoonDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *universeMoon) updateTableName(table string) *universeMoon {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt32(table, "id")
	u.SystemID = field.NewInt32(table, "system_id")

	u.fillFieldMap()

	return u
}

func (u *universeMoon) WithContext(ctx context.Context) *universeMoonDo {
	return u.universeMoonDo.WithContext(ctx)
}

func (u universeMoon) TableName() string { return u.universeMoonDo.TableName() }

func (u universeMoon) Alias() string { return u.universeMoonDo.Alias() }

func (u universeMoon) Columns(cols ...field.Expr) gen.Columns {
	return u.universeMoonDo.Columns(cols...)
}

func (u *universeMoon) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *universeMoon) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 2)
	u.fieldMap["id"] = u.ID
	u.fieldMap["system_id"] = u.SystemID
}

func (u universeMoon) clone(db *gorm.DB) universeMoon {
	u.universeMoonDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u universeMoon) replaceDB(db *gorm.DB) universeMoon {
	u.universeMoonDo.ReplaceDB(db)
	return u
}

type universeMoonDo struct{ gen.DO }

func (u universeMoonDo) Debug() *universeMoonDo {
	return u.withDO(u.DO.Debug())
}

func (u universeMoonDo) WithContext(ctx context.Context) *universeMoonDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u universeMoonDo) ReadDB() *universeMoonDo {
	return u.Clauses(dbresolver.Read)
}

func (u universeMoonDo) WriteDB() *universeMoonDo {
	return u.Clauses(dbresolver.Write)
}

func (u universeMoonDo) Session(config *gorm.Session) *universeMoonDo {
	return u.withDO(u.DO.Session(config))
}

func (u universeMoonDo) Clauses(conds ...clause.Expression) *universeMoonDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u universeMoonDo) Returning(value interface{}, columns ...string) *universeMoonDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u universeMoonDo) Not(conds ...gen.Condition) *universeMoonDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u universeMoonDo) Or(conds ...gen.Condition) *universeMoonDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u universeMoonDo) Select(conds ...field.Expr) *universeMoonDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u universeMoonDo) Where(conds ...gen.Condition) *universeMoonDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u universeMoonDo) Order(conds ...field.Expr) *universeMoonDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u universeMoonDo) Distinct(cols ...field.Expr) *universeMoonDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u universeMoonDo) Omit(cols ...field.Expr) *universeMoonDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u universeMoonDo) Join(table schema.Tabler, on ...field.Expr) *universeMoonDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u universeMoonDo) LeftJoin(table schema.Tabler, on ...field.Expr) *universeMoonDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u universeMoonDo) RightJoin(table schema.Tabler, on ...field.Expr) *universeMoonDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u universeMoonDo) Group(cols ...field.Expr) *universeMoonDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u universeMoonDo) Having(conds ...gen.Condition) *universeMoonDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u universeMoonDo) Limit(limit int) *universeMoonDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u universeMoonDo) Offset(offset int) *universeMoonDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u universeMoonDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *universeMoonDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u universeMoonDo) Unscoped() *universeMoonDo {
	return u.withDO(u.DO.Unscoped())
}

func (u universeMoonDo) Create(values ...*model.UniverseMoon) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u universeMoonDo) CreateInBatches(values []*model.UniverseMoon, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u universeMoonDo) Save(values ...*model.UniverseMoon) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u universeMoonDo) First() (*model.UniverseMoon, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseMoon), nil
	}
}

func (u universeMoonDo) Take() (*model.UniverseMoon, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseMoon), nil
	}
}

func (u universeMoonDo) Last() (*model.UniverseMoon, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseMoon), nil
	}
}

func (u universeMoonDo) Find() ([]*model.UniverseMoon, error) {
	result, err := u.DO.Find()
	return result.([]*model.UniverseMoon), err
}

func (u universeMoonDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UniverseMoon, err error) {
	buf := make([]*model.UniverseMoon, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u universeMoonDo) FindInBatches(result *[]*model.UniverseMoon, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u universeMoonDo) Attrs(attrs ...field.AssignExpr) *universeMoonDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u universeMoonDo) Assign(attrs ...field.AssignExpr) *universeMoonDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u universeMoonDo) Joins(fields ...field.RelationField) *universeMoonDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u universeMoonDo) Preload(fields ...field.RelationField) *universeMoonDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u universeMoonDo) FirstOrInit() (*model.UniverseMoon, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseMoon), nil
	}
}

func (u universeMoonDo) FirstOrCreate() (*model.UniverseMoon, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseMoon), nil
	}
}

func (u universeMoonDo) FindByPage(offset int, limit int) (result []*model.UniverseMoon, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u universeMoonDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u universeMoonDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u universeMoonDo) Delete(models ...*model.UniverseMoon) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *universeMoonDo) withDO(do gen.Dao) *universeMoonDo {
	u.DO = *do.(*gen.DO)
	return u
}
