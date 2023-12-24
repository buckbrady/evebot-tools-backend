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

func newKillmailItem(db *gorm.DB, opts ...gen.DOOption) killmailItem {
	_killmailItem := killmailItem{}

	_killmailItem.killmailItemDo.UseDB(db, opts...)
	_killmailItem.killmailItemDo.UseModel(&model.KillmailItem{})

	tableName := _killmailItem.killmailItemDo.TableName()
	_killmailItem.ALL = field.NewAsterisk(tableName)
	_killmailItem.ID = field.NewInt64(tableName, "id")
	_killmailItem.KillmailID = field.NewInt64(tableName, "killmail_id")
	_killmailItem.TypeID = field.NewInt32(tableName, "type_id")
	_killmailItem.Flag = field.NewInt32(tableName, "flag")
	_killmailItem.QuantityDestroyed = field.NewInt32(tableName, "quantity_destroyed")
	_killmailItem.QuantityDropped = field.NewInt32(tableName, "quantity_dropped")
	_killmailItem.Singleton = field.NewBool(tableName, "singleton")

	_killmailItem.fillFieldMap()

	return _killmailItem
}

type killmailItem struct {
	killmailItemDo killmailItemDo

	ALL               field.Asterisk
	ID                field.Int64
	KillmailID        field.Int64
	TypeID            field.Int32
	Flag              field.Int32
	QuantityDestroyed field.Int32
	QuantityDropped   field.Int32
	Singleton         field.Bool

	fieldMap map[string]field.Expr
}

func (k killmailItem) Table(newTableName string) *killmailItem {
	k.killmailItemDo.UseTable(newTableName)
	return k.updateTableName(newTableName)
}

func (k killmailItem) As(alias string) *killmailItem {
	k.killmailItemDo.DO = *(k.killmailItemDo.As(alias).(*gen.DO))
	return k.updateTableName(alias)
}

func (k *killmailItem) updateTableName(table string) *killmailItem {
	k.ALL = field.NewAsterisk(table)
	k.ID = field.NewInt64(table, "id")
	k.KillmailID = field.NewInt64(table, "killmail_id")
	k.TypeID = field.NewInt32(table, "type_id")
	k.Flag = field.NewInt32(table, "flag")
	k.QuantityDestroyed = field.NewInt32(table, "quantity_destroyed")
	k.QuantityDropped = field.NewInt32(table, "quantity_dropped")
	k.Singleton = field.NewBool(table, "singleton")

	k.fillFieldMap()

	return k
}

func (k *killmailItem) WithContext(ctx context.Context) *killmailItemDo {
	return k.killmailItemDo.WithContext(ctx)
}

func (k killmailItem) TableName() string { return k.killmailItemDo.TableName() }

func (k killmailItem) Alias() string { return k.killmailItemDo.Alias() }

func (k killmailItem) Columns(cols ...field.Expr) gen.Columns {
	return k.killmailItemDo.Columns(cols...)
}

func (k *killmailItem) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := k.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (k *killmailItem) fillFieldMap() {
	k.fieldMap = make(map[string]field.Expr, 7)
	k.fieldMap["id"] = k.ID
	k.fieldMap["killmail_id"] = k.KillmailID
	k.fieldMap["type_id"] = k.TypeID
	k.fieldMap["flag"] = k.Flag
	k.fieldMap["quantity_destroyed"] = k.QuantityDestroyed
	k.fieldMap["quantity_dropped"] = k.QuantityDropped
	k.fieldMap["singleton"] = k.Singleton
}

func (k killmailItem) clone(db *gorm.DB) killmailItem {
	k.killmailItemDo.ReplaceConnPool(db.Statement.ConnPool)
	return k
}

func (k killmailItem) replaceDB(db *gorm.DB) killmailItem {
	k.killmailItemDo.ReplaceDB(db)
	return k
}

type killmailItemDo struct{ gen.DO }

func (k killmailItemDo) Debug() *killmailItemDo {
	return k.withDO(k.DO.Debug())
}

func (k killmailItemDo) WithContext(ctx context.Context) *killmailItemDo {
	return k.withDO(k.DO.WithContext(ctx))
}

func (k killmailItemDo) ReadDB() *killmailItemDo {
	return k.Clauses(dbresolver.Read)
}

func (k killmailItemDo) WriteDB() *killmailItemDo {
	return k.Clauses(dbresolver.Write)
}

func (k killmailItemDo) Session(config *gorm.Session) *killmailItemDo {
	return k.withDO(k.DO.Session(config))
}

func (k killmailItemDo) Clauses(conds ...clause.Expression) *killmailItemDo {
	return k.withDO(k.DO.Clauses(conds...))
}

func (k killmailItemDo) Returning(value interface{}, columns ...string) *killmailItemDo {
	return k.withDO(k.DO.Returning(value, columns...))
}

func (k killmailItemDo) Not(conds ...gen.Condition) *killmailItemDo {
	return k.withDO(k.DO.Not(conds...))
}

func (k killmailItemDo) Or(conds ...gen.Condition) *killmailItemDo {
	return k.withDO(k.DO.Or(conds...))
}

func (k killmailItemDo) Select(conds ...field.Expr) *killmailItemDo {
	return k.withDO(k.DO.Select(conds...))
}

func (k killmailItemDo) Where(conds ...gen.Condition) *killmailItemDo {
	return k.withDO(k.DO.Where(conds...))
}

func (k killmailItemDo) Order(conds ...field.Expr) *killmailItemDo {
	return k.withDO(k.DO.Order(conds...))
}

func (k killmailItemDo) Distinct(cols ...field.Expr) *killmailItemDo {
	return k.withDO(k.DO.Distinct(cols...))
}

func (k killmailItemDo) Omit(cols ...field.Expr) *killmailItemDo {
	return k.withDO(k.DO.Omit(cols...))
}

func (k killmailItemDo) Join(table schema.Tabler, on ...field.Expr) *killmailItemDo {
	return k.withDO(k.DO.Join(table, on...))
}

func (k killmailItemDo) LeftJoin(table schema.Tabler, on ...field.Expr) *killmailItemDo {
	return k.withDO(k.DO.LeftJoin(table, on...))
}

func (k killmailItemDo) RightJoin(table schema.Tabler, on ...field.Expr) *killmailItemDo {
	return k.withDO(k.DO.RightJoin(table, on...))
}

func (k killmailItemDo) Group(cols ...field.Expr) *killmailItemDo {
	return k.withDO(k.DO.Group(cols...))
}

func (k killmailItemDo) Having(conds ...gen.Condition) *killmailItemDo {
	return k.withDO(k.DO.Having(conds...))
}

func (k killmailItemDo) Limit(limit int) *killmailItemDo {
	return k.withDO(k.DO.Limit(limit))
}

func (k killmailItemDo) Offset(offset int) *killmailItemDo {
	return k.withDO(k.DO.Offset(offset))
}

func (k killmailItemDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *killmailItemDo {
	return k.withDO(k.DO.Scopes(funcs...))
}

func (k killmailItemDo) Unscoped() *killmailItemDo {
	return k.withDO(k.DO.Unscoped())
}

func (k killmailItemDo) Create(values ...*model.KillmailItem) error {
	if len(values) == 0 {
		return nil
	}
	return k.DO.Create(values)
}

func (k killmailItemDo) CreateInBatches(values []*model.KillmailItem, batchSize int) error {
	return k.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (k killmailItemDo) Save(values ...*model.KillmailItem) error {
	if len(values) == 0 {
		return nil
	}
	return k.DO.Save(values)
}

func (k killmailItemDo) First() (*model.KillmailItem, error) {
	if result, err := k.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.KillmailItem), nil
	}
}

func (k killmailItemDo) Take() (*model.KillmailItem, error) {
	if result, err := k.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.KillmailItem), nil
	}
}

func (k killmailItemDo) Last() (*model.KillmailItem, error) {
	if result, err := k.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.KillmailItem), nil
	}
}

func (k killmailItemDo) Find() ([]*model.KillmailItem, error) {
	result, err := k.DO.Find()
	return result.([]*model.KillmailItem), err
}

func (k killmailItemDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.KillmailItem, err error) {
	buf := make([]*model.KillmailItem, 0, batchSize)
	err = k.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (k killmailItemDo) FindInBatches(result *[]*model.KillmailItem, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return k.DO.FindInBatches(result, batchSize, fc)
}

func (k killmailItemDo) Attrs(attrs ...field.AssignExpr) *killmailItemDo {
	return k.withDO(k.DO.Attrs(attrs...))
}

func (k killmailItemDo) Assign(attrs ...field.AssignExpr) *killmailItemDo {
	return k.withDO(k.DO.Assign(attrs...))
}

func (k killmailItemDo) Joins(fields ...field.RelationField) *killmailItemDo {
	for _, _f := range fields {
		k = *k.withDO(k.DO.Joins(_f))
	}
	return &k
}

func (k killmailItemDo) Preload(fields ...field.RelationField) *killmailItemDo {
	for _, _f := range fields {
		k = *k.withDO(k.DO.Preload(_f))
	}
	return &k
}

func (k killmailItemDo) FirstOrInit() (*model.KillmailItem, error) {
	if result, err := k.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.KillmailItem), nil
	}
}

func (k killmailItemDo) FirstOrCreate() (*model.KillmailItem, error) {
	if result, err := k.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.KillmailItem), nil
	}
}

func (k killmailItemDo) FindByPage(offset int, limit int) (result []*model.KillmailItem, count int64, err error) {
	result, err = k.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = k.Offset(-1).Limit(-1).Count()
	return
}

func (k killmailItemDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = k.Count()
	if err != nil {
		return
	}

	err = k.Offset(offset).Limit(limit).Scan(result)
	return
}

func (k killmailItemDo) Scan(result interface{}) (err error) {
	return k.DO.Scan(result)
}

func (k killmailItemDo) Delete(models ...*model.KillmailItem) (result gen.ResultInfo, err error) {
	return k.DO.Delete(models)
}

func (k *killmailItemDo) withDO(do gen.Dao) *killmailItemDo {
	k.DO = *do.(*gen.DO)
	return k
}