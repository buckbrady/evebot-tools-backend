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

func newUniverseStructure(db *gorm.DB, opts ...gen.DOOption) universeStructure {
	_universeStructure := universeStructure{}

	_universeStructure.universeStructureDo.UseDB(db, opts...)
	_universeStructure.universeStructureDo.UseModel(&model.UniverseStructure{})

	tableName := _universeStructure.universeStructureDo.TableName()
	_universeStructure.ALL = field.NewAsterisk(tableName)
	_universeStructure.ID = field.NewInt64(tableName, "id")
	_universeStructure.Name = field.NewString(tableName, "name")
	_universeStructure.OwnerID = field.NewInt32(tableName, "owner_id")
	_universeStructure.PositionX = field.NewFloat64(tableName, "position_x")
	_universeStructure.PositionY = field.NewFloat64(tableName, "position_y")
	_universeStructure.PositionZ = field.NewFloat64(tableName, "position_z")
	_universeStructure.SolarSystemID = field.NewInt32(tableName, "solar_system_id")
	_universeStructure.TypeID = field.NewInt32(tableName, "type_id")

	_universeStructure.fillFieldMap()

	return _universeStructure
}

type universeStructure struct {
	universeStructureDo universeStructureDo

	ALL           field.Asterisk
	ID            field.Int64
	Name          field.String
	OwnerID       field.Int32
	PositionX     field.Float64
	PositionY     field.Float64
	PositionZ     field.Float64
	SolarSystemID field.Int32
	TypeID        field.Int32

	fieldMap map[string]field.Expr
}

func (u universeStructure) Table(newTableName string) *universeStructure {
	u.universeStructureDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u universeStructure) As(alias string) *universeStructure {
	u.universeStructureDo.DO = *(u.universeStructureDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *universeStructure) updateTableName(table string) *universeStructure {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.Name = field.NewString(table, "name")
	u.OwnerID = field.NewInt32(table, "owner_id")
	u.PositionX = field.NewFloat64(table, "position_x")
	u.PositionY = field.NewFloat64(table, "position_y")
	u.PositionZ = field.NewFloat64(table, "position_z")
	u.SolarSystemID = field.NewInt32(table, "solar_system_id")
	u.TypeID = field.NewInt32(table, "type_id")

	u.fillFieldMap()

	return u
}

func (u *universeStructure) WithContext(ctx context.Context) *universeStructureDo {
	return u.universeStructureDo.WithContext(ctx)
}

func (u universeStructure) TableName() string { return u.universeStructureDo.TableName() }

func (u universeStructure) Alias() string { return u.universeStructureDo.Alias() }

func (u universeStructure) Columns(cols ...field.Expr) gen.Columns {
	return u.universeStructureDo.Columns(cols...)
}

func (u *universeStructure) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *universeStructure) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 8)
	u.fieldMap["id"] = u.ID
	u.fieldMap["name"] = u.Name
	u.fieldMap["owner_id"] = u.OwnerID
	u.fieldMap["position_x"] = u.PositionX
	u.fieldMap["position_y"] = u.PositionY
	u.fieldMap["position_z"] = u.PositionZ
	u.fieldMap["solar_system_id"] = u.SolarSystemID
	u.fieldMap["type_id"] = u.TypeID
}

func (u universeStructure) clone(db *gorm.DB) universeStructure {
	u.universeStructureDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u universeStructure) replaceDB(db *gorm.DB) universeStructure {
	u.universeStructureDo.ReplaceDB(db)
	return u
}

type universeStructureDo struct{ gen.DO }

func (u universeStructureDo) Debug() *universeStructureDo {
	return u.withDO(u.DO.Debug())
}

func (u universeStructureDo) WithContext(ctx context.Context) *universeStructureDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u universeStructureDo) ReadDB() *universeStructureDo {
	return u.Clauses(dbresolver.Read)
}

func (u universeStructureDo) WriteDB() *universeStructureDo {
	return u.Clauses(dbresolver.Write)
}

func (u universeStructureDo) Session(config *gorm.Session) *universeStructureDo {
	return u.withDO(u.DO.Session(config))
}

func (u universeStructureDo) Clauses(conds ...clause.Expression) *universeStructureDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u universeStructureDo) Returning(value interface{}, columns ...string) *universeStructureDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u universeStructureDo) Not(conds ...gen.Condition) *universeStructureDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u universeStructureDo) Or(conds ...gen.Condition) *universeStructureDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u universeStructureDo) Select(conds ...field.Expr) *universeStructureDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u universeStructureDo) Where(conds ...gen.Condition) *universeStructureDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u universeStructureDo) Order(conds ...field.Expr) *universeStructureDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u universeStructureDo) Distinct(cols ...field.Expr) *universeStructureDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u universeStructureDo) Omit(cols ...field.Expr) *universeStructureDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u universeStructureDo) Join(table schema.Tabler, on ...field.Expr) *universeStructureDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u universeStructureDo) LeftJoin(table schema.Tabler, on ...field.Expr) *universeStructureDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u universeStructureDo) RightJoin(table schema.Tabler, on ...field.Expr) *universeStructureDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u universeStructureDo) Group(cols ...field.Expr) *universeStructureDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u universeStructureDo) Having(conds ...gen.Condition) *universeStructureDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u universeStructureDo) Limit(limit int) *universeStructureDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u universeStructureDo) Offset(offset int) *universeStructureDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u universeStructureDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *universeStructureDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u universeStructureDo) Unscoped() *universeStructureDo {
	return u.withDO(u.DO.Unscoped())
}

func (u universeStructureDo) Create(values ...*model.UniverseStructure) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u universeStructureDo) CreateInBatches(values []*model.UniverseStructure, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u universeStructureDo) Save(values ...*model.UniverseStructure) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u universeStructureDo) First() (*model.UniverseStructure, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseStructure), nil
	}
}

func (u universeStructureDo) Take() (*model.UniverseStructure, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseStructure), nil
	}
}

func (u universeStructureDo) Last() (*model.UniverseStructure, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseStructure), nil
	}
}

func (u universeStructureDo) Find() ([]*model.UniverseStructure, error) {
	result, err := u.DO.Find()
	return result.([]*model.UniverseStructure), err
}

func (u universeStructureDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UniverseStructure, err error) {
	buf := make([]*model.UniverseStructure, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u universeStructureDo) FindInBatches(result *[]*model.UniverseStructure, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u universeStructureDo) Attrs(attrs ...field.AssignExpr) *universeStructureDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u universeStructureDo) Assign(attrs ...field.AssignExpr) *universeStructureDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u universeStructureDo) Joins(fields ...field.RelationField) *universeStructureDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u universeStructureDo) Preload(fields ...field.RelationField) *universeStructureDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u universeStructureDo) FirstOrInit() (*model.UniverseStructure, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseStructure), nil
	}
}

func (u universeStructureDo) FirstOrCreate() (*model.UniverseStructure, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseStructure), nil
	}
}

func (u universeStructureDo) FindByPage(offset int, limit int) (result []*model.UniverseStructure, count int64, err error) {
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

func (u universeStructureDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u universeStructureDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u universeStructureDo) Delete(models ...*model.UniverseStructure) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *universeStructureDo) withDO(do gen.Dao) *universeStructureDo {
	u.DO = *do.(*gen.DO)
	return u
}