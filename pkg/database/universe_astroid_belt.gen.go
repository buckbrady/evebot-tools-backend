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

func newUniverseAstroidBelt(db *gorm.DB, opts ...gen.DOOption) universeAstroidBelt {
	_universeAstroidBelt := universeAstroidBelt{}

	_universeAstroidBelt.universeAstroidBeltDo.UseDB(db, opts...)
	_universeAstroidBelt.universeAstroidBeltDo.UseModel(&model.UniverseAstroidBelt{})

	tableName := _universeAstroidBelt.universeAstroidBeltDo.TableName()
	_universeAstroidBelt.ALL = field.NewAsterisk(tableName)
	_universeAstroidBelt.ID = field.NewInt32(tableName, "id")
	_universeAstroidBelt.SystemID = field.NewInt32(tableName, "system_id")
	_universeAstroidBelt.PositionX = field.NewFloat64(tableName, "position_x")
	_universeAstroidBelt.PositionY = field.NewFloat64(tableName, "position_y")
	_universeAstroidBelt.PositionZ = field.NewFloat64(tableName, "position_z")
	_universeAstroidBelt.Name = field.NewString(tableName, "name")

	_universeAstroidBelt.fillFieldMap()

	return _universeAstroidBelt
}

type universeAstroidBelt struct {
	universeAstroidBeltDo universeAstroidBeltDo

	ALL       field.Asterisk
	ID        field.Int32
	SystemID  field.Int32
	PositionX field.Float64
	PositionY field.Float64
	PositionZ field.Float64
	Name      field.String

	fieldMap map[string]field.Expr
}

func (u universeAstroidBelt) Table(newTableName string) *universeAstroidBelt {
	u.universeAstroidBeltDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u universeAstroidBelt) As(alias string) *universeAstroidBelt {
	u.universeAstroidBeltDo.DO = *(u.universeAstroidBeltDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *universeAstroidBelt) updateTableName(table string) *universeAstroidBelt {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt32(table, "id")
	u.SystemID = field.NewInt32(table, "system_id")
	u.PositionX = field.NewFloat64(table, "position_x")
	u.PositionY = field.NewFloat64(table, "position_y")
	u.PositionZ = field.NewFloat64(table, "position_z")
	u.Name = field.NewString(table, "name")

	u.fillFieldMap()

	return u
}

func (u *universeAstroidBelt) WithContext(ctx context.Context) *universeAstroidBeltDo {
	return u.universeAstroidBeltDo.WithContext(ctx)
}

func (u universeAstroidBelt) TableName() string { return u.universeAstroidBeltDo.TableName() }

func (u universeAstroidBelt) Alias() string { return u.universeAstroidBeltDo.Alias() }

func (u universeAstroidBelt) Columns(cols ...field.Expr) gen.Columns {
	return u.universeAstroidBeltDo.Columns(cols...)
}

func (u *universeAstroidBelt) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *universeAstroidBelt) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 6)
	u.fieldMap["id"] = u.ID
	u.fieldMap["system_id"] = u.SystemID
	u.fieldMap["position_x"] = u.PositionX
	u.fieldMap["position_y"] = u.PositionY
	u.fieldMap["position_z"] = u.PositionZ
	u.fieldMap["name"] = u.Name
}

func (u universeAstroidBelt) clone(db *gorm.DB) universeAstroidBelt {
	u.universeAstroidBeltDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u universeAstroidBelt) replaceDB(db *gorm.DB) universeAstroidBelt {
	u.universeAstroidBeltDo.ReplaceDB(db)
	return u
}

type universeAstroidBeltDo struct{ gen.DO }

func (u universeAstroidBeltDo) Debug() *universeAstroidBeltDo {
	return u.withDO(u.DO.Debug())
}

func (u universeAstroidBeltDo) WithContext(ctx context.Context) *universeAstroidBeltDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u universeAstroidBeltDo) ReadDB() *universeAstroidBeltDo {
	return u.Clauses(dbresolver.Read)
}

func (u universeAstroidBeltDo) WriteDB() *universeAstroidBeltDo {
	return u.Clauses(dbresolver.Write)
}

func (u universeAstroidBeltDo) Session(config *gorm.Session) *universeAstroidBeltDo {
	return u.withDO(u.DO.Session(config))
}

func (u universeAstroidBeltDo) Clauses(conds ...clause.Expression) *universeAstroidBeltDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u universeAstroidBeltDo) Returning(value interface{}, columns ...string) *universeAstroidBeltDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u universeAstroidBeltDo) Not(conds ...gen.Condition) *universeAstroidBeltDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u universeAstroidBeltDo) Or(conds ...gen.Condition) *universeAstroidBeltDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u universeAstroidBeltDo) Select(conds ...field.Expr) *universeAstroidBeltDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u universeAstroidBeltDo) Where(conds ...gen.Condition) *universeAstroidBeltDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u universeAstroidBeltDo) Order(conds ...field.Expr) *universeAstroidBeltDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u universeAstroidBeltDo) Distinct(cols ...field.Expr) *universeAstroidBeltDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u universeAstroidBeltDo) Omit(cols ...field.Expr) *universeAstroidBeltDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u universeAstroidBeltDo) Join(table schema.Tabler, on ...field.Expr) *universeAstroidBeltDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u universeAstroidBeltDo) LeftJoin(table schema.Tabler, on ...field.Expr) *universeAstroidBeltDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u universeAstroidBeltDo) RightJoin(table schema.Tabler, on ...field.Expr) *universeAstroidBeltDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u universeAstroidBeltDo) Group(cols ...field.Expr) *universeAstroidBeltDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u universeAstroidBeltDo) Having(conds ...gen.Condition) *universeAstroidBeltDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u universeAstroidBeltDo) Limit(limit int) *universeAstroidBeltDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u universeAstroidBeltDo) Offset(offset int) *universeAstroidBeltDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u universeAstroidBeltDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *universeAstroidBeltDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u universeAstroidBeltDo) Unscoped() *universeAstroidBeltDo {
	return u.withDO(u.DO.Unscoped())
}

func (u universeAstroidBeltDo) Create(values ...*model.UniverseAstroidBelt) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u universeAstroidBeltDo) CreateInBatches(values []*model.UniverseAstroidBelt, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u universeAstroidBeltDo) Save(values ...*model.UniverseAstroidBelt) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u universeAstroidBeltDo) First() (*model.UniverseAstroidBelt, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseAstroidBelt), nil
	}
}

func (u universeAstroidBeltDo) Take() (*model.UniverseAstroidBelt, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseAstroidBelt), nil
	}
}

func (u universeAstroidBeltDo) Last() (*model.UniverseAstroidBelt, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseAstroidBelt), nil
	}
}

func (u universeAstroidBeltDo) Find() ([]*model.UniverseAstroidBelt, error) {
	result, err := u.DO.Find()
	return result.([]*model.UniverseAstroidBelt), err
}

func (u universeAstroidBeltDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UniverseAstroidBelt, err error) {
	buf := make([]*model.UniverseAstroidBelt, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u universeAstroidBeltDo) FindInBatches(result *[]*model.UniverseAstroidBelt, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u universeAstroidBeltDo) Attrs(attrs ...field.AssignExpr) *universeAstroidBeltDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u universeAstroidBeltDo) Assign(attrs ...field.AssignExpr) *universeAstroidBeltDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u universeAstroidBeltDo) Joins(fields ...field.RelationField) *universeAstroidBeltDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u universeAstroidBeltDo) Preload(fields ...field.RelationField) *universeAstroidBeltDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u universeAstroidBeltDo) FirstOrInit() (*model.UniverseAstroidBelt, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseAstroidBelt), nil
	}
}

func (u universeAstroidBeltDo) FirstOrCreate() (*model.UniverseAstroidBelt, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseAstroidBelt), nil
	}
}

func (u universeAstroidBeltDo) FindByPage(offset int, limit int) (result []*model.UniverseAstroidBelt, count int64, err error) {
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

func (u universeAstroidBeltDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u universeAstroidBeltDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u universeAstroidBeltDo) Delete(models ...*model.UniverseAstroidBelt) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *universeAstroidBeltDo) withDO(do gen.Dao) *universeAstroidBeltDo {
	u.DO = *do.(*gen.DO)
	return u
}
