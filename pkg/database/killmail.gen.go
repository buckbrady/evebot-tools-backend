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

func newKillmail(db *gorm.DB, opts ...gen.DOOption) killmail {
	_killmail := killmail{}

	_killmail.killmailDo.UseDB(db, opts...)
	_killmail.killmailDo.UseModel(&model.Killmail{})

	tableName := _killmail.killmailDo.TableName()
	_killmail.ALL = field.NewAsterisk(tableName)
	_killmail.ID = field.NewInt64(tableName, "id")
	_killmail.KillmailTime = field.NewTime(tableName, "killmail_time")
	_killmail.SolarSystemID = field.NewInt32(tableName, "solar_system_id")
	_killmail.WarID = field.NewInt64(tableName, "war_id")
	_killmail.DamageTaken = field.NewInt32(tableName, "damage_taken")
	_killmail.TotalValue = field.NewFloat64(tableName, "total_value")

	_killmail.fillFieldMap()

	return _killmail
}

type killmail struct {
	killmailDo killmailDo

	ALL           field.Asterisk
	ID            field.Int64
	KillmailTime  field.Time
	SolarSystemID field.Int32
	WarID         field.Int64
	DamageTaken   field.Int32
	TotalValue    field.Float64

	fieldMap map[string]field.Expr
}

func (k killmail) Table(newTableName string) *killmail {
	k.killmailDo.UseTable(newTableName)
	return k.updateTableName(newTableName)
}

func (k killmail) As(alias string) *killmail {
	k.killmailDo.DO = *(k.killmailDo.As(alias).(*gen.DO))
	return k.updateTableName(alias)
}

func (k *killmail) updateTableName(table string) *killmail {
	k.ALL = field.NewAsterisk(table)
	k.ID = field.NewInt64(table, "id")
	k.KillmailTime = field.NewTime(table, "killmail_time")
	k.SolarSystemID = field.NewInt32(table, "solar_system_id")
	k.WarID = field.NewInt64(table, "war_id")
	k.DamageTaken = field.NewInt32(table, "damage_taken")
	k.TotalValue = field.NewFloat64(table, "total_value")

	k.fillFieldMap()

	return k
}

func (k *killmail) WithContext(ctx context.Context) *killmailDo { return k.killmailDo.WithContext(ctx) }

func (k killmail) TableName() string { return k.killmailDo.TableName() }

func (k killmail) Alias() string { return k.killmailDo.Alias() }

func (k killmail) Columns(cols ...field.Expr) gen.Columns { return k.killmailDo.Columns(cols...) }

func (k *killmail) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := k.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (k *killmail) fillFieldMap() {
	k.fieldMap = make(map[string]field.Expr, 6)
	k.fieldMap["id"] = k.ID
	k.fieldMap["killmail_time"] = k.KillmailTime
	k.fieldMap["solar_system_id"] = k.SolarSystemID
	k.fieldMap["war_id"] = k.WarID
	k.fieldMap["damage_taken"] = k.DamageTaken
	k.fieldMap["total_value"] = k.TotalValue
}

func (k killmail) clone(db *gorm.DB) killmail {
	k.killmailDo.ReplaceConnPool(db.Statement.ConnPool)
	return k
}

func (k killmail) replaceDB(db *gorm.DB) killmail {
	k.killmailDo.ReplaceDB(db)
	return k
}

type killmailDo struct{ gen.DO }

func (k killmailDo) Debug() *killmailDo {
	return k.withDO(k.DO.Debug())
}

func (k killmailDo) WithContext(ctx context.Context) *killmailDo {
	return k.withDO(k.DO.WithContext(ctx))
}

func (k killmailDo) ReadDB() *killmailDo {
	return k.Clauses(dbresolver.Read)
}

func (k killmailDo) WriteDB() *killmailDo {
	return k.Clauses(dbresolver.Write)
}

func (k killmailDo) Session(config *gorm.Session) *killmailDo {
	return k.withDO(k.DO.Session(config))
}

func (k killmailDo) Clauses(conds ...clause.Expression) *killmailDo {
	return k.withDO(k.DO.Clauses(conds...))
}

func (k killmailDo) Returning(value interface{}, columns ...string) *killmailDo {
	return k.withDO(k.DO.Returning(value, columns...))
}

func (k killmailDo) Not(conds ...gen.Condition) *killmailDo {
	return k.withDO(k.DO.Not(conds...))
}

func (k killmailDo) Or(conds ...gen.Condition) *killmailDo {
	return k.withDO(k.DO.Or(conds...))
}

func (k killmailDo) Select(conds ...field.Expr) *killmailDo {
	return k.withDO(k.DO.Select(conds...))
}

func (k killmailDo) Where(conds ...gen.Condition) *killmailDo {
	return k.withDO(k.DO.Where(conds...))
}

func (k killmailDo) Order(conds ...field.Expr) *killmailDo {
	return k.withDO(k.DO.Order(conds...))
}

func (k killmailDo) Distinct(cols ...field.Expr) *killmailDo {
	return k.withDO(k.DO.Distinct(cols...))
}

func (k killmailDo) Omit(cols ...field.Expr) *killmailDo {
	return k.withDO(k.DO.Omit(cols...))
}

func (k killmailDo) Join(table schema.Tabler, on ...field.Expr) *killmailDo {
	return k.withDO(k.DO.Join(table, on...))
}

func (k killmailDo) LeftJoin(table schema.Tabler, on ...field.Expr) *killmailDo {
	return k.withDO(k.DO.LeftJoin(table, on...))
}

func (k killmailDo) RightJoin(table schema.Tabler, on ...field.Expr) *killmailDo {
	return k.withDO(k.DO.RightJoin(table, on...))
}

func (k killmailDo) Group(cols ...field.Expr) *killmailDo {
	return k.withDO(k.DO.Group(cols...))
}

func (k killmailDo) Having(conds ...gen.Condition) *killmailDo {
	return k.withDO(k.DO.Having(conds...))
}

func (k killmailDo) Limit(limit int) *killmailDo {
	return k.withDO(k.DO.Limit(limit))
}

func (k killmailDo) Offset(offset int) *killmailDo {
	return k.withDO(k.DO.Offset(offset))
}

func (k killmailDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *killmailDo {
	return k.withDO(k.DO.Scopes(funcs...))
}

func (k killmailDo) Unscoped() *killmailDo {
	return k.withDO(k.DO.Unscoped())
}

func (k killmailDo) Create(values ...*model.Killmail) error {
	if len(values) == 0 {
		return nil
	}
	return k.DO.Create(values)
}

func (k killmailDo) CreateInBatches(values []*model.Killmail, batchSize int) error {
	return k.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (k killmailDo) Save(values ...*model.Killmail) error {
	if len(values) == 0 {
		return nil
	}
	return k.DO.Save(values)
}

func (k killmailDo) First() (*model.Killmail, error) {
	if result, err := k.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Killmail), nil
	}
}

func (k killmailDo) Take() (*model.Killmail, error) {
	if result, err := k.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Killmail), nil
	}
}

func (k killmailDo) Last() (*model.Killmail, error) {
	if result, err := k.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Killmail), nil
	}
}

func (k killmailDo) Find() ([]*model.Killmail, error) {
	result, err := k.DO.Find()
	return result.([]*model.Killmail), err
}

func (k killmailDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Killmail, err error) {
	buf := make([]*model.Killmail, 0, batchSize)
	err = k.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (k killmailDo) FindInBatches(result *[]*model.Killmail, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return k.DO.FindInBatches(result, batchSize, fc)
}

func (k killmailDo) Attrs(attrs ...field.AssignExpr) *killmailDo {
	return k.withDO(k.DO.Attrs(attrs...))
}

func (k killmailDo) Assign(attrs ...field.AssignExpr) *killmailDo {
	return k.withDO(k.DO.Assign(attrs...))
}

func (k killmailDo) Joins(fields ...field.RelationField) *killmailDo {
	for _, _f := range fields {
		k = *k.withDO(k.DO.Joins(_f))
	}
	return &k
}

func (k killmailDo) Preload(fields ...field.RelationField) *killmailDo {
	for _, _f := range fields {
		k = *k.withDO(k.DO.Preload(_f))
	}
	return &k
}

func (k killmailDo) FirstOrInit() (*model.Killmail, error) {
	if result, err := k.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Killmail), nil
	}
}

func (k killmailDo) FirstOrCreate() (*model.Killmail, error) {
	if result, err := k.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Killmail), nil
	}
}

func (k killmailDo) FindByPage(offset int, limit int) (result []*model.Killmail, count int64, err error) {
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

func (k killmailDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = k.Count()
	if err != nil {
		return
	}

	err = k.Offset(offset).Limit(limit).Scan(result)
	return
}

func (k killmailDo) Scan(result interface{}) (err error) {
	return k.DO.Scan(result)
}

func (k killmailDo) Delete(models ...*model.Killmail) (result gen.ResultInfo, err error) {
	return k.DO.Delete(models)
}

func (k *killmailDo) withDO(do gen.Dao) *killmailDo {
	k.DO = *do.(*gen.DO)
	return k
}
