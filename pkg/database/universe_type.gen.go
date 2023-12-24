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

func newUniverseType(db *gorm.DB, opts ...gen.DOOption) universeType {
	_universeType := universeType{}

	_universeType.universeTypeDo.UseDB(db, opts...)
	_universeType.universeTypeDo.UseModel(&model.UniverseType{})

	tableName := _universeType.universeTypeDo.TableName()
	_universeType.ALL = field.NewAsterisk(tableName)
	_universeType.ID = field.NewInt32(tableName, "id")
	_universeType.Name = field.NewString(tableName, "name")
	_universeType.Description = field.NewString(tableName, "description")
	_universeType.Capacity = field.NewFloat64(tableName, "capacity")
	_universeType.GraphicID = field.NewInt32(tableName, "graphic_id")
	_universeType.GroupID = field.NewInt32(tableName, "group_id")
	_universeType.IconID = field.NewInt32(tableName, "icon_id")
	_universeType.MarketGroupID = field.NewInt32(tableName, "market_group_id")
	_universeType.Mass = field.NewFloat64(tableName, "mass")
	_universeType.PackagedVolume = field.NewFloat64(tableName, "packaged_volume")
	_universeType.PortionSize = field.NewInt32(tableName, "portion_size")
	_universeType.Published = field.NewBool(tableName, "published")
	_universeType.Radius = field.NewFloat64(tableName, "radius")
	_universeType.Volume = field.NewFloat64(tableName, "volume")

	_universeType.fillFieldMap()

	return _universeType
}

type universeType struct {
	universeTypeDo universeTypeDo

	ALL            field.Asterisk
	ID             field.Int32
	Name           field.String
	Description    field.String
	Capacity       field.Float64
	GraphicID      field.Int32
	GroupID        field.Int32
	IconID         field.Int32
	MarketGroupID  field.Int32
	Mass           field.Float64
	PackagedVolume field.Float64
	PortionSize    field.Int32
	Published      field.Bool
	Radius         field.Float64
	Volume         field.Float64

	fieldMap map[string]field.Expr
}

func (u universeType) Table(newTableName string) *universeType {
	u.universeTypeDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u universeType) As(alias string) *universeType {
	u.universeTypeDo.DO = *(u.universeTypeDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *universeType) updateTableName(table string) *universeType {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt32(table, "id")
	u.Name = field.NewString(table, "name")
	u.Description = field.NewString(table, "description")
	u.Capacity = field.NewFloat64(table, "capacity")
	u.GraphicID = field.NewInt32(table, "graphic_id")
	u.GroupID = field.NewInt32(table, "group_id")
	u.IconID = field.NewInt32(table, "icon_id")
	u.MarketGroupID = field.NewInt32(table, "market_group_id")
	u.Mass = field.NewFloat64(table, "mass")
	u.PackagedVolume = field.NewFloat64(table, "packaged_volume")
	u.PortionSize = field.NewInt32(table, "portion_size")
	u.Published = field.NewBool(table, "published")
	u.Radius = field.NewFloat64(table, "radius")
	u.Volume = field.NewFloat64(table, "volume")

	u.fillFieldMap()

	return u
}

func (u *universeType) WithContext(ctx context.Context) *universeTypeDo {
	return u.universeTypeDo.WithContext(ctx)
}

func (u universeType) TableName() string { return u.universeTypeDo.TableName() }

func (u universeType) Alias() string { return u.universeTypeDo.Alias() }

func (u universeType) Columns(cols ...field.Expr) gen.Columns {
	return u.universeTypeDo.Columns(cols...)
}

func (u *universeType) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *universeType) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 14)
	u.fieldMap["id"] = u.ID
	u.fieldMap["name"] = u.Name
	u.fieldMap["description"] = u.Description
	u.fieldMap["capacity"] = u.Capacity
	u.fieldMap["graphic_id"] = u.GraphicID
	u.fieldMap["group_id"] = u.GroupID
	u.fieldMap["icon_id"] = u.IconID
	u.fieldMap["market_group_id"] = u.MarketGroupID
	u.fieldMap["mass"] = u.Mass
	u.fieldMap["packaged_volume"] = u.PackagedVolume
	u.fieldMap["portion_size"] = u.PortionSize
	u.fieldMap["published"] = u.Published
	u.fieldMap["radius"] = u.Radius
	u.fieldMap["volume"] = u.Volume
}

func (u universeType) clone(db *gorm.DB) universeType {
	u.universeTypeDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u universeType) replaceDB(db *gorm.DB) universeType {
	u.universeTypeDo.ReplaceDB(db)
	return u
}

type universeTypeDo struct{ gen.DO }

func (u universeTypeDo) Debug() *universeTypeDo {
	return u.withDO(u.DO.Debug())
}

func (u universeTypeDo) WithContext(ctx context.Context) *universeTypeDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u universeTypeDo) ReadDB() *universeTypeDo {
	return u.Clauses(dbresolver.Read)
}

func (u universeTypeDo) WriteDB() *universeTypeDo {
	return u.Clauses(dbresolver.Write)
}

func (u universeTypeDo) Session(config *gorm.Session) *universeTypeDo {
	return u.withDO(u.DO.Session(config))
}

func (u universeTypeDo) Clauses(conds ...clause.Expression) *universeTypeDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u universeTypeDo) Returning(value interface{}, columns ...string) *universeTypeDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u universeTypeDo) Not(conds ...gen.Condition) *universeTypeDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u universeTypeDo) Or(conds ...gen.Condition) *universeTypeDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u universeTypeDo) Select(conds ...field.Expr) *universeTypeDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u universeTypeDo) Where(conds ...gen.Condition) *universeTypeDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u universeTypeDo) Order(conds ...field.Expr) *universeTypeDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u universeTypeDo) Distinct(cols ...field.Expr) *universeTypeDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u universeTypeDo) Omit(cols ...field.Expr) *universeTypeDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u universeTypeDo) Join(table schema.Tabler, on ...field.Expr) *universeTypeDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u universeTypeDo) LeftJoin(table schema.Tabler, on ...field.Expr) *universeTypeDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u universeTypeDo) RightJoin(table schema.Tabler, on ...field.Expr) *universeTypeDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u universeTypeDo) Group(cols ...field.Expr) *universeTypeDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u universeTypeDo) Having(conds ...gen.Condition) *universeTypeDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u universeTypeDo) Limit(limit int) *universeTypeDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u universeTypeDo) Offset(offset int) *universeTypeDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u universeTypeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *universeTypeDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u universeTypeDo) Unscoped() *universeTypeDo {
	return u.withDO(u.DO.Unscoped())
}

func (u universeTypeDo) Create(values ...*model.UniverseType) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u universeTypeDo) CreateInBatches(values []*model.UniverseType, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u universeTypeDo) Save(values ...*model.UniverseType) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u universeTypeDo) First() (*model.UniverseType, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseType), nil
	}
}

func (u universeTypeDo) Take() (*model.UniverseType, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseType), nil
	}
}

func (u universeTypeDo) Last() (*model.UniverseType, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseType), nil
	}
}

func (u universeTypeDo) Find() ([]*model.UniverseType, error) {
	result, err := u.DO.Find()
	return result.([]*model.UniverseType), err
}

func (u universeTypeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UniverseType, err error) {
	buf := make([]*model.UniverseType, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u universeTypeDo) FindInBatches(result *[]*model.UniverseType, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u universeTypeDo) Attrs(attrs ...field.AssignExpr) *universeTypeDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u universeTypeDo) Assign(attrs ...field.AssignExpr) *universeTypeDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u universeTypeDo) Joins(fields ...field.RelationField) *universeTypeDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u universeTypeDo) Preload(fields ...field.RelationField) *universeTypeDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u universeTypeDo) FirstOrInit() (*model.UniverseType, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseType), nil
	}
}

func (u universeTypeDo) FirstOrCreate() (*model.UniverseType, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseType), nil
	}
}

func (u universeTypeDo) FindByPage(offset int, limit int) (result []*model.UniverseType, count int64, err error) {
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

func (u universeTypeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u universeTypeDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u universeTypeDo) Delete(models ...*model.UniverseType) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *universeTypeDo) withDO(do gen.Dao) *universeTypeDo {
	u.DO = *do.(*gen.DO)
	return u
}
