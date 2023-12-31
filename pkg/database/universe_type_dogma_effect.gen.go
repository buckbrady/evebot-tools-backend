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

func newUniverseTypeDogmaEffect(db *gorm.DB, opts ...gen.DOOption) universeTypeDogmaEffect {
	_universeTypeDogmaEffect := universeTypeDogmaEffect{}

	_universeTypeDogmaEffect.universeTypeDogmaEffectDo.UseDB(db, opts...)
	_universeTypeDogmaEffect.universeTypeDogmaEffectDo.UseModel(&model.UniverseTypeDogmaEffect{})

	tableName := _universeTypeDogmaEffect.universeTypeDogmaEffectDo.TableName()
	_universeTypeDogmaEffect.ALL = field.NewAsterisk(tableName)
	_universeTypeDogmaEffect.TypeID = field.NewInt32(tableName, "type_id")
	_universeTypeDogmaEffect.EffectID = field.NewInt32(tableName, "effect_id")
	_universeTypeDogmaEffect.IsDefault = field.NewBool(tableName, "is_default")

	_universeTypeDogmaEffect.fillFieldMap()

	return _universeTypeDogmaEffect
}

type universeTypeDogmaEffect struct {
	universeTypeDogmaEffectDo universeTypeDogmaEffectDo

	ALL       field.Asterisk
	TypeID    field.Int32
	EffectID  field.Int32
	IsDefault field.Bool

	fieldMap map[string]field.Expr
}

func (u universeTypeDogmaEffect) Table(newTableName string) *universeTypeDogmaEffect {
	u.universeTypeDogmaEffectDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u universeTypeDogmaEffect) As(alias string) *universeTypeDogmaEffect {
	u.universeTypeDogmaEffectDo.DO = *(u.universeTypeDogmaEffectDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *universeTypeDogmaEffect) updateTableName(table string) *universeTypeDogmaEffect {
	u.ALL = field.NewAsterisk(table)
	u.TypeID = field.NewInt32(table, "type_id")
	u.EffectID = field.NewInt32(table, "effect_id")
	u.IsDefault = field.NewBool(table, "is_default")

	u.fillFieldMap()

	return u
}

func (u *universeTypeDogmaEffect) WithContext(ctx context.Context) *universeTypeDogmaEffectDo {
	return u.universeTypeDogmaEffectDo.WithContext(ctx)
}

func (u universeTypeDogmaEffect) TableName() string { return u.universeTypeDogmaEffectDo.TableName() }

func (u universeTypeDogmaEffect) Alias() string { return u.universeTypeDogmaEffectDo.Alias() }

func (u universeTypeDogmaEffect) Columns(cols ...field.Expr) gen.Columns {
	return u.universeTypeDogmaEffectDo.Columns(cols...)
}

func (u *universeTypeDogmaEffect) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *universeTypeDogmaEffect) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 3)
	u.fieldMap["type_id"] = u.TypeID
	u.fieldMap["effect_id"] = u.EffectID
	u.fieldMap["is_default"] = u.IsDefault
}

func (u universeTypeDogmaEffect) clone(db *gorm.DB) universeTypeDogmaEffect {
	u.universeTypeDogmaEffectDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u universeTypeDogmaEffect) replaceDB(db *gorm.DB) universeTypeDogmaEffect {
	u.universeTypeDogmaEffectDo.ReplaceDB(db)
	return u
}

type universeTypeDogmaEffectDo struct{ gen.DO }

func (u universeTypeDogmaEffectDo) Debug() *universeTypeDogmaEffectDo {
	return u.withDO(u.DO.Debug())
}

func (u universeTypeDogmaEffectDo) WithContext(ctx context.Context) *universeTypeDogmaEffectDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u universeTypeDogmaEffectDo) ReadDB() *universeTypeDogmaEffectDo {
	return u.Clauses(dbresolver.Read)
}

func (u universeTypeDogmaEffectDo) WriteDB() *universeTypeDogmaEffectDo {
	return u.Clauses(dbresolver.Write)
}

func (u universeTypeDogmaEffectDo) Session(config *gorm.Session) *universeTypeDogmaEffectDo {
	return u.withDO(u.DO.Session(config))
}

func (u universeTypeDogmaEffectDo) Clauses(conds ...clause.Expression) *universeTypeDogmaEffectDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u universeTypeDogmaEffectDo) Returning(value interface{}, columns ...string) *universeTypeDogmaEffectDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u universeTypeDogmaEffectDo) Not(conds ...gen.Condition) *universeTypeDogmaEffectDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u universeTypeDogmaEffectDo) Or(conds ...gen.Condition) *universeTypeDogmaEffectDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u universeTypeDogmaEffectDo) Select(conds ...field.Expr) *universeTypeDogmaEffectDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u universeTypeDogmaEffectDo) Where(conds ...gen.Condition) *universeTypeDogmaEffectDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u universeTypeDogmaEffectDo) Order(conds ...field.Expr) *universeTypeDogmaEffectDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u universeTypeDogmaEffectDo) Distinct(cols ...field.Expr) *universeTypeDogmaEffectDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u universeTypeDogmaEffectDo) Omit(cols ...field.Expr) *universeTypeDogmaEffectDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u universeTypeDogmaEffectDo) Join(table schema.Tabler, on ...field.Expr) *universeTypeDogmaEffectDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u universeTypeDogmaEffectDo) LeftJoin(table schema.Tabler, on ...field.Expr) *universeTypeDogmaEffectDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u universeTypeDogmaEffectDo) RightJoin(table schema.Tabler, on ...field.Expr) *universeTypeDogmaEffectDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u universeTypeDogmaEffectDo) Group(cols ...field.Expr) *universeTypeDogmaEffectDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u universeTypeDogmaEffectDo) Having(conds ...gen.Condition) *universeTypeDogmaEffectDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u universeTypeDogmaEffectDo) Limit(limit int) *universeTypeDogmaEffectDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u universeTypeDogmaEffectDo) Offset(offset int) *universeTypeDogmaEffectDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u universeTypeDogmaEffectDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *universeTypeDogmaEffectDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u universeTypeDogmaEffectDo) Unscoped() *universeTypeDogmaEffectDo {
	return u.withDO(u.DO.Unscoped())
}

func (u universeTypeDogmaEffectDo) Create(values ...*model.UniverseTypeDogmaEffect) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u universeTypeDogmaEffectDo) CreateInBatches(values []*model.UniverseTypeDogmaEffect, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u universeTypeDogmaEffectDo) Save(values ...*model.UniverseTypeDogmaEffect) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u universeTypeDogmaEffectDo) First() (*model.UniverseTypeDogmaEffect, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseTypeDogmaEffect), nil
	}
}

func (u universeTypeDogmaEffectDo) Take() (*model.UniverseTypeDogmaEffect, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseTypeDogmaEffect), nil
	}
}

func (u universeTypeDogmaEffectDo) Last() (*model.UniverseTypeDogmaEffect, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseTypeDogmaEffect), nil
	}
}

func (u universeTypeDogmaEffectDo) Find() ([]*model.UniverseTypeDogmaEffect, error) {
	result, err := u.DO.Find()
	return result.([]*model.UniverseTypeDogmaEffect), err
}

func (u universeTypeDogmaEffectDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UniverseTypeDogmaEffect, err error) {
	buf := make([]*model.UniverseTypeDogmaEffect, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u universeTypeDogmaEffectDo) FindInBatches(result *[]*model.UniverseTypeDogmaEffect, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u universeTypeDogmaEffectDo) Attrs(attrs ...field.AssignExpr) *universeTypeDogmaEffectDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u universeTypeDogmaEffectDo) Assign(attrs ...field.AssignExpr) *universeTypeDogmaEffectDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u universeTypeDogmaEffectDo) Joins(fields ...field.RelationField) *universeTypeDogmaEffectDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u universeTypeDogmaEffectDo) Preload(fields ...field.RelationField) *universeTypeDogmaEffectDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u universeTypeDogmaEffectDo) FirstOrInit() (*model.UniverseTypeDogmaEffect, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseTypeDogmaEffect), nil
	}
}

func (u universeTypeDogmaEffectDo) FirstOrCreate() (*model.UniverseTypeDogmaEffect, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseTypeDogmaEffect), nil
	}
}

func (u universeTypeDogmaEffectDo) FindByPage(offset int, limit int) (result []*model.UniverseTypeDogmaEffect, count int64, err error) {
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

func (u universeTypeDogmaEffectDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u universeTypeDogmaEffectDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u universeTypeDogmaEffectDo) Delete(models ...*model.UniverseTypeDogmaEffect) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *universeTypeDogmaEffectDo) withDO(do gen.Dao) *universeTypeDogmaEffectDo {
	u.DO = *do.(*gen.DO)
	return u
}
