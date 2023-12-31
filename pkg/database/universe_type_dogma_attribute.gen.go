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

func newUniverseTypeDogmaAttribute(db *gorm.DB, opts ...gen.DOOption) universeTypeDogmaAttribute {
	_universeTypeDogmaAttribute := universeTypeDogmaAttribute{}

	_universeTypeDogmaAttribute.universeTypeDogmaAttributeDo.UseDB(db, opts...)
	_universeTypeDogmaAttribute.universeTypeDogmaAttributeDo.UseModel(&model.UniverseTypeDogmaAttribute{})

	tableName := _universeTypeDogmaAttribute.universeTypeDogmaAttributeDo.TableName()
	_universeTypeDogmaAttribute.ALL = field.NewAsterisk(tableName)
	_universeTypeDogmaAttribute.TypeID = field.NewInt32(tableName, "type_id")
	_universeTypeDogmaAttribute.AttributeID = field.NewInt32(tableName, "attribute_id")
	_universeTypeDogmaAttribute.Value = field.NewFloat64(tableName, "value")

	_universeTypeDogmaAttribute.fillFieldMap()

	return _universeTypeDogmaAttribute
}

type universeTypeDogmaAttribute struct {
	universeTypeDogmaAttributeDo universeTypeDogmaAttributeDo

	ALL         field.Asterisk
	TypeID      field.Int32
	AttributeID field.Int32
	Value       field.Float64

	fieldMap map[string]field.Expr
}

func (u universeTypeDogmaAttribute) Table(newTableName string) *universeTypeDogmaAttribute {
	u.universeTypeDogmaAttributeDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u universeTypeDogmaAttribute) As(alias string) *universeTypeDogmaAttribute {
	u.universeTypeDogmaAttributeDo.DO = *(u.universeTypeDogmaAttributeDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *universeTypeDogmaAttribute) updateTableName(table string) *universeTypeDogmaAttribute {
	u.ALL = field.NewAsterisk(table)
	u.TypeID = field.NewInt32(table, "type_id")
	u.AttributeID = field.NewInt32(table, "attribute_id")
	u.Value = field.NewFloat64(table, "value")

	u.fillFieldMap()

	return u
}

func (u *universeTypeDogmaAttribute) WithContext(ctx context.Context) *universeTypeDogmaAttributeDo {
	return u.universeTypeDogmaAttributeDo.WithContext(ctx)
}

func (u universeTypeDogmaAttribute) TableName() string {
	return u.universeTypeDogmaAttributeDo.TableName()
}

func (u universeTypeDogmaAttribute) Alias() string { return u.universeTypeDogmaAttributeDo.Alias() }

func (u universeTypeDogmaAttribute) Columns(cols ...field.Expr) gen.Columns {
	return u.universeTypeDogmaAttributeDo.Columns(cols...)
}

func (u *universeTypeDogmaAttribute) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *universeTypeDogmaAttribute) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 3)
	u.fieldMap["type_id"] = u.TypeID
	u.fieldMap["attribute_id"] = u.AttributeID
	u.fieldMap["value"] = u.Value
}

func (u universeTypeDogmaAttribute) clone(db *gorm.DB) universeTypeDogmaAttribute {
	u.universeTypeDogmaAttributeDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u universeTypeDogmaAttribute) replaceDB(db *gorm.DB) universeTypeDogmaAttribute {
	u.universeTypeDogmaAttributeDo.ReplaceDB(db)
	return u
}

type universeTypeDogmaAttributeDo struct{ gen.DO }

func (u universeTypeDogmaAttributeDo) Debug() *universeTypeDogmaAttributeDo {
	return u.withDO(u.DO.Debug())
}

func (u universeTypeDogmaAttributeDo) WithContext(ctx context.Context) *universeTypeDogmaAttributeDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u universeTypeDogmaAttributeDo) ReadDB() *universeTypeDogmaAttributeDo {
	return u.Clauses(dbresolver.Read)
}

func (u universeTypeDogmaAttributeDo) WriteDB() *universeTypeDogmaAttributeDo {
	return u.Clauses(dbresolver.Write)
}

func (u universeTypeDogmaAttributeDo) Session(config *gorm.Session) *universeTypeDogmaAttributeDo {
	return u.withDO(u.DO.Session(config))
}

func (u universeTypeDogmaAttributeDo) Clauses(conds ...clause.Expression) *universeTypeDogmaAttributeDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u universeTypeDogmaAttributeDo) Returning(value interface{}, columns ...string) *universeTypeDogmaAttributeDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u universeTypeDogmaAttributeDo) Not(conds ...gen.Condition) *universeTypeDogmaAttributeDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u universeTypeDogmaAttributeDo) Or(conds ...gen.Condition) *universeTypeDogmaAttributeDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u universeTypeDogmaAttributeDo) Select(conds ...field.Expr) *universeTypeDogmaAttributeDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u universeTypeDogmaAttributeDo) Where(conds ...gen.Condition) *universeTypeDogmaAttributeDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u universeTypeDogmaAttributeDo) Order(conds ...field.Expr) *universeTypeDogmaAttributeDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u universeTypeDogmaAttributeDo) Distinct(cols ...field.Expr) *universeTypeDogmaAttributeDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u universeTypeDogmaAttributeDo) Omit(cols ...field.Expr) *universeTypeDogmaAttributeDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u universeTypeDogmaAttributeDo) Join(table schema.Tabler, on ...field.Expr) *universeTypeDogmaAttributeDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u universeTypeDogmaAttributeDo) LeftJoin(table schema.Tabler, on ...field.Expr) *universeTypeDogmaAttributeDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u universeTypeDogmaAttributeDo) RightJoin(table schema.Tabler, on ...field.Expr) *universeTypeDogmaAttributeDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u universeTypeDogmaAttributeDo) Group(cols ...field.Expr) *universeTypeDogmaAttributeDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u universeTypeDogmaAttributeDo) Having(conds ...gen.Condition) *universeTypeDogmaAttributeDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u universeTypeDogmaAttributeDo) Limit(limit int) *universeTypeDogmaAttributeDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u universeTypeDogmaAttributeDo) Offset(offset int) *universeTypeDogmaAttributeDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u universeTypeDogmaAttributeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *universeTypeDogmaAttributeDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u universeTypeDogmaAttributeDo) Unscoped() *universeTypeDogmaAttributeDo {
	return u.withDO(u.DO.Unscoped())
}

func (u universeTypeDogmaAttributeDo) Create(values ...*model.UniverseTypeDogmaAttribute) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u universeTypeDogmaAttributeDo) CreateInBatches(values []*model.UniverseTypeDogmaAttribute, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u universeTypeDogmaAttributeDo) Save(values ...*model.UniverseTypeDogmaAttribute) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u universeTypeDogmaAttributeDo) First() (*model.UniverseTypeDogmaAttribute, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseTypeDogmaAttribute), nil
	}
}

func (u universeTypeDogmaAttributeDo) Take() (*model.UniverseTypeDogmaAttribute, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseTypeDogmaAttribute), nil
	}
}

func (u universeTypeDogmaAttributeDo) Last() (*model.UniverseTypeDogmaAttribute, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseTypeDogmaAttribute), nil
	}
}

func (u universeTypeDogmaAttributeDo) Find() ([]*model.UniverseTypeDogmaAttribute, error) {
	result, err := u.DO.Find()
	return result.([]*model.UniverseTypeDogmaAttribute), err
}

func (u universeTypeDogmaAttributeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UniverseTypeDogmaAttribute, err error) {
	buf := make([]*model.UniverseTypeDogmaAttribute, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u universeTypeDogmaAttributeDo) FindInBatches(result *[]*model.UniverseTypeDogmaAttribute, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u universeTypeDogmaAttributeDo) Attrs(attrs ...field.AssignExpr) *universeTypeDogmaAttributeDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u universeTypeDogmaAttributeDo) Assign(attrs ...field.AssignExpr) *universeTypeDogmaAttributeDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u universeTypeDogmaAttributeDo) Joins(fields ...field.RelationField) *universeTypeDogmaAttributeDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u universeTypeDogmaAttributeDo) Preload(fields ...field.RelationField) *universeTypeDogmaAttributeDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u universeTypeDogmaAttributeDo) FirstOrInit() (*model.UniverseTypeDogmaAttribute, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseTypeDogmaAttribute), nil
	}
}

func (u universeTypeDogmaAttributeDo) FirstOrCreate() (*model.UniverseTypeDogmaAttribute, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UniverseTypeDogmaAttribute), nil
	}
}

func (u universeTypeDogmaAttributeDo) FindByPage(offset int, limit int) (result []*model.UniverseTypeDogmaAttribute, count int64, err error) {
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

func (u universeTypeDogmaAttributeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u universeTypeDogmaAttributeDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u universeTypeDogmaAttributeDo) Delete(models ...*model.UniverseTypeDogmaAttribute) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *universeTypeDogmaAttributeDo) withDO(do gen.Dao) *universeTypeDogmaAttributeDo {
	u.DO = *do.(*gen.DO)
	return u
}
