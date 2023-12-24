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

func newUniverseStar(db *gorm.DB, opts ...gen.DOOption) universeStar {
	_universeStar := universeStar{}

	_universeStar.universeStarDo.UseDB(db, opts...)
	_universeStar.universeStarDo.UseModel(&model.UniverseStar{})

	tableName := _universeStar.universeStarDo.TableName()
	_universeStar.ALL = field.NewAsterisk(tableName)
	_universeStar.ID = field.NewInt32(tableName, "id")
	_universeStar.Name = field.NewString(tableName, "name")
	_universeStar.Age = field.NewInt64(tableName, "age")
	_universeStar.Luminosity = field.NewFloat64(tableName, "luminosity")
	_universeStar.Radius = field.NewInt64(tableName, "radius")
	_universeStar.SolarSystemID = field.NewInt32(tableName, "solar_system_id")
	_universeStar.SpectralClass = field.NewString(tableName, "spectral_class")
	_universeStar.Temperature = field.NewInt32(tableName, "temperature")
	_universeStar.TypeID = field.NewInt32(tableName, "type_id")

	_universeStar.fillFieldMap()

	return _universeStar
}

type universeStar struct {
	universeStarDo universeStarDo

	ALL           field.Asterisk
	ID            field.Int32
	Name          field.String
	Age           field.Int64
	Luminosity    field.Float64
	Radius        field.Int64
	SolarSystemID field.Int32
	SpectralClass field.String
	Temperature   field.Int32
	TypeID        field.Int32

	fieldMap map[string]field.Expr
}

func (u universeStar) Table(newTableName string) *universeStar {
	u.universeStarDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u universeStar) As(alias string) *universeStar {
	u.universeStarDo.DO = *(u.universeStarDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *universeStar) updateTableName(table string) *universeStar {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt32(table, "id")
	u.Name = field.NewString(table, "name")
	u.Age = field.NewInt64(table, "age")
	u.Luminosity = field.NewFloat64(table, "luminosity")
	u.Radius = field.NewInt64(table, "radius")
	u.SolarSystemID = field.NewInt32(table, "solar_system_id")
	u.SpectralClass = field.NewString(table, "spectral_class")
	u.Temperature = field.NewInt32(table, "temperature")
	u.TypeID = field.NewInt32(table, "type_id")

	u.fillFieldMap()

	return u
}

func (u *universeStar) WithContext(ctx context.Context) *universeStarDo {
	return u.universeStarDo.WithContext(ctx)
}

func (u universeStar) TableName() string { return u.universeStarDo.TableName() }

func (u universeStar) Alias() string { return u.universeStarDo.Alias() }

func (u universeStar) Columns(cols ...field.Expr) gen.Columns {
	return u.universeStarDo.Columns(cols...)
}

func (u *universeStar) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *universeStar) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 9)
	u.fieldMap["id"] = u.ID
	u.fieldMap["name"] = u.Name
	u.fieldMap["age"] = u.Age
	u.fieldMap["luminosity"] = u.Luminosity
	u.fieldMap["radius"] = u.Radius
	u.fieldMap["solar_system_id"] = u.SolarSystemID
	u.fieldMap["spectral_class"] = u.SpectralClass
	u.fieldMap["temperature"] = u.Temperature
	u.fieldMap["type_id"] = u.TypeID
}

func (u universeStar) clone(db *gorm.DB) universeStar {
	u.universeStarDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u universeStar) replaceDB(db *gorm.DB) universeStar {
	u.universeStarDo.ReplaceDB(db)
	return u
}

type universeStarDo struct{ gen.DO }

func (u universeStarDo) Debug() *universeStarDo {
	return u.withDO(u.DO.Debug())
}

func (u universeStarDo) WithContext(ctx context.Context) *universeStarDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u universeStarDo) ReadDB() *universeStarDo {
	return u.Clauses(dbresolver.Read)
}

func (u universeStarDo) WriteDB() *universeStarDo {
	return u.Clauses(dbresolver.Write)
}

func (u universeStarDo) Session(config *gorm.Session) *universeStarDo {
	return u.withDO(u.DO.Session(config))
}

func (u universeStarDo) Clauses(conds ...clause.Expression) *universeStarDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u universeStarDo) Returning(value interface{}, columns ...string) *universeStarDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u universeStarDo) Not(conds ...gen.Condition) *universeStarDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u universeStarDo) Or(conds ...gen.Condition) *universeStarDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u universeStarDo) Select(conds ...field.Expr) *universeStarDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u universeStarDo) Where(conds ...gen.Condition) *universeStarDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u universeStarDo) Order(conds ...field.Expr) *universeStarDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u universeStarDo) Distinct(cols ...field.Expr) *universeStarDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u universeStarDo) Omit(cols ...field.Expr) *universeStarDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u universeStarDo) Join(table schema.Tabler, on ...field.Expr) *universeStarDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u universeStarDo) LeftJoin(table schema.Tabler, on ...field.Expr) *universeStarDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u universeStarDo) RightJoin(table schema.Tabler, on ...field.Expr) *universeStarDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u universeStarDo) Group(cols ...field.Expr) *universeStarDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u universeStarDo) Having(conds ...gen.Condition) *universeStarDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u universeStarDo) Limit(limit int) *universeStarDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u universeStarDo) Offset(offset int) *universeStarDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u universeStarDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *universeStarDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u universeStarDo) Unscoped() *universeStarDo {
	return u.withDO(u.DO.Unscoped())
}

func (u universeStarDo) Create(values ...*model.UniverseStar) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u universeStarDo) CreateInBatches(values []*model.UniverseStar, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u universeStarDo) Save(values ...*model.UniverseStar) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u universeStarDo) First() (*model.UniverseStar, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseStar), nil
	}
}

func (u universeStarDo) Take() (*model.UniverseStar, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseStar), nil
	}
}

func (u universeStarDo) Last() (*model.UniverseStar, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseStar), nil
	}
}

func (u universeStarDo) Find() ([]*model.UniverseStar, error) {
	result, err := u.DO.Find()
	return result.([]*model.UniverseStar), err
}

func (u universeStarDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UniverseStar, err error) {
	buf := make([]*model.UniverseStar, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u universeStarDo) FindInBatches(result *[]*model.UniverseStar, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u universeStarDo) Attrs(attrs ...field.AssignExpr) *universeStarDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u universeStarDo) Assign(attrs ...field.AssignExpr) *universeStarDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u universeStarDo) Joins(fields ...field.RelationField) *universeStarDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u universeStarDo) Preload(fields ...field.RelationField) *universeStarDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u universeStarDo) FirstOrInit() (*model.UniverseStar, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseStar), nil
	}
}

func (u universeStarDo) FirstOrCreate() (*model.UniverseStar, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseStar), nil
	}
}

func (u universeStarDo) FindByPage(offset int, limit int) (result []*model.UniverseStar, count int64, err error) {
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

func (u universeStarDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u universeStarDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u universeStarDo) Delete(models ...*model.UniverseStar) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *universeStarDo) withDO(do gen.Dao) *universeStarDo {
	u.DO = *do.(*gen.DO)
	return u
}
