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

func newUniverseRegion(db *gorm.DB, opts ...gen.DOOption) universeRegion {
	_universeRegion := universeRegion{}

	_universeRegion.universeRegionDo.UseDB(db, opts...)
	_universeRegion.universeRegionDo.UseModel(&model.UniverseRegion{})

	tableName := _universeRegion.universeRegionDo.TableName()
	_universeRegion.ALL = field.NewAsterisk(tableName)
	_universeRegion.ID = field.NewInt32(tableName, "id")
	_universeRegion.Name = field.NewString(tableName, "name")
	_universeRegion.Description = field.NewString(tableName, "description")

	_universeRegion.fillFieldMap()

	return _universeRegion
}

type universeRegion struct {
	universeRegionDo universeRegionDo

	ALL         field.Asterisk
	ID          field.Int32
	Name        field.String
	Description field.String

	fieldMap map[string]field.Expr
}

func (u universeRegion) Table(newTableName string) *universeRegion {
	u.universeRegionDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u universeRegion) As(alias string) *universeRegion {
	u.universeRegionDo.DO = *(u.universeRegionDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *universeRegion) updateTableName(table string) *universeRegion {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt32(table, "id")
	u.Name = field.NewString(table, "name")
	u.Description = field.NewString(table, "description")

	u.fillFieldMap()

	return u
}

func (u *universeRegion) WithContext(ctx context.Context) *universeRegionDo {
	return u.universeRegionDo.WithContext(ctx)
}

func (u universeRegion) TableName() string { return u.universeRegionDo.TableName() }

func (u universeRegion) Alias() string { return u.universeRegionDo.Alias() }

func (u universeRegion) Columns(cols ...field.Expr) gen.Columns {
	return u.universeRegionDo.Columns(cols...)
}

func (u *universeRegion) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *universeRegion) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 3)
	u.fieldMap["id"] = u.ID
	u.fieldMap["name"] = u.Name
	u.fieldMap["description"] = u.Description
}

func (u universeRegion) clone(db *gorm.DB) universeRegion {
	u.universeRegionDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u universeRegion) replaceDB(db *gorm.DB) universeRegion {
	u.universeRegionDo.ReplaceDB(db)
	return u
}

type universeRegionDo struct{ gen.DO }

func (u universeRegionDo) Debug() *universeRegionDo {
	return u.withDO(u.DO.Debug())
}

func (u universeRegionDo) WithContext(ctx context.Context) *universeRegionDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u universeRegionDo) ReadDB() *universeRegionDo {
	return u.Clauses(dbresolver.Read)
}

func (u universeRegionDo) WriteDB() *universeRegionDo {
	return u.Clauses(dbresolver.Write)
}

func (u universeRegionDo) Session(config *gorm.Session) *universeRegionDo {
	return u.withDO(u.DO.Session(config))
}

func (u universeRegionDo) Clauses(conds ...clause.Expression) *universeRegionDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u universeRegionDo) Returning(value interface{}, columns ...string) *universeRegionDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u universeRegionDo) Not(conds ...gen.Condition) *universeRegionDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u universeRegionDo) Or(conds ...gen.Condition) *universeRegionDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u universeRegionDo) Select(conds ...field.Expr) *universeRegionDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u universeRegionDo) Where(conds ...gen.Condition) *universeRegionDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u universeRegionDo) Order(conds ...field.Expr) *universeRegionDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u universeRegionDo) Distinct(cols ...field.Expr) *universeRegionDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u universeRegionDo) Omit(cols ...field.Expr) *universeRegionDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u universeRegionDo) Join(table schema.Tabler, on ...field.Expr) *universeRegionDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u universeRegionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *universeRegionDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u universeRegionDo) RightJoin(table schema.Tabler, on ...field.Expr) *universeRegionDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u universeRegionDo) Group(cols ...field.Expr) *universeRegionDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u universeRegionDo) Having(conds ...gen.Condition) *universeRegionDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u universeRegionDo) Limit(limit int) *universeRegionDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u universeRegionDo) Offset(offset int) *universeRegionDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u universeRegionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *universeRegionDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u universeRegionDo) Unscoped() *universeRegionDo {
	return u.withDO(u.DO.Unscoped())
}

func (u universeRegionDo) Create(values ...*model.UniverseRegion) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u universeRegionDo) CreateInBatches(values []*model.UniverseRegion, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u universeRegionDo) Save(values ...*model.UniverseRegion) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u universeRegionDo) First() (*model.UniverseRegion, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseRegion), nil
	}
}

func (u universeRegionDo) Take() (*model.UniverseRegion, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseRegion), nil
	}
}

func (u universeRegionDo) Last() (*model.UniverseRegion, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseRegion), nil
	}
}

func (u universeRegionDo) Find() ([]*model.UniverseRegion, error) {
	result, err := u.DO.Find()
	return result.([]*model.UniverseRegion), err
}

func (u universeRegionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UniverseRegion, err error) {
	buf := make([]*model.UniverseRegion, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u universeRegionDo) FindInBatches(result *[]*model.UniverseRegion, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u universeRegionDo) Attrs(attrs ...field.AssignExpr) *universeRegionDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u universeRegionDo) Assign(attrs ...field.AssignExpr) *universeRegionDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u universeRegionDo) Joins(fields ...field.RelationField) *universeRegionDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u universeRegionDo) Preload(fields ...field.RelationField) *universeRegionDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u universeRegionDo) FirstOrInit() (*model.UniverseRegion, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseRegion), nil
	}
}

func (u universeRegionDo) FirstOrCreate() (*model.UniverseRegion, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseRegion), nil
	}
}

func (u universeRegionDo) FindByPage(offset int, limit int) (result []*model.UniverseRegion, count int64, err error) {
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

func (u universeRegionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u universeRegionDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u universeRegionDo) Delete(models ...*model.UniverseRegion) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *universeRegionDo) withDO(do gen.Dao) *universeRegionDo {
	u.DO = *do.(*gen.DO)
	return u
}
