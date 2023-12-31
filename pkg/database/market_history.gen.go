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

func newMarketHistory(db *gorm.DB, opts ...gen.DOOption) marketHistory {
	_marketHistory := marketHistory{}

	_marketHistory.marketHistoryDo.UseDB(db, opts...)
	_marketHistory.marketHistoryDo.UseModel(&model.MarketHistory{})

	tableName := _marketHistory.marketHistoryDo.TableName()
	_marketHistory.ALL = field.NewAsterisk(tableName)
	_marketHistory.TypeID = field.NewInt32(tableName, "type_id")
	_marketHistory.RegionID = field.NewInt32(tableName, "region_id")
	_marketHistory.Date = field.NewTime(tableName, "date")
	_marketHistory.Average = field.NewFloat64(tableName, "average")
	_marketHistory.Highest = field.NewFloat64(tableName, "highest")
	_marketHistory.Lowest = field.NewFloat64(tableName, "lowest")
	_marketHistory.OrderCount = field.NewInt64(tableName, "order_count")
	_marketHistory.Volume = field.NewInt64(tableName, "volume")

	_marketHistory.fillFieldMap()

	return _marketHistory
}

type marketHistory struct {
	marketHistoryDo marketHistoryDo

	ALL        field.Asterisk
	TypeID     field.Int32
	RegionID   field.Int32
	Date       field.Time
	Average    field.Float64
	Highest    field.Float64
	Lowest     field.Float64
	OrderCount field.Int64
	Volume     field.Int64

	fieldMap map[string]field.Expr
}

func (m marketHistory) Table(newTableName string) *marketHistory {
	m.marketHistoryDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m marketHistory) As(alias string) *marketHistory {
	m.marketHistoryDo.DO = *(m.marketHistoryDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *marketHistory) updateTableName(table string) *marketHistory {
	m.ALL = field.NewAsterisk(table)
	m.TypeID = field.NewInt32(table, "type_id")
	m.RegionID = field.NewInt32(table, "region_id")
	m.Date = field.NewTime(table, "date")
	m.Average = field.NewFloat64(table, "average")
	m.Highest = field.NewFloat64(table, "highest")
	m.Lowest = field.NewFloat64(table, "lowest")
	m.OrderCount = field.NewInt64(table, "order_count")
	m.Volume = field.NewInt64(table, "volume")

	m.fillFieldMap()

	return m
}

func (m *marketHistory) WithContext(ctx context.Context) *marketHistoryDo {
	return m.marketHistoryDo.WithContext(ctx)
}

func (m marketHistory) TableName() string { return m.marketHistoryDo.TableName() }

func (m marketHistory) Alias() string { return m.marketHistoryDo.Alias() }

func (m marketHistory) Columns(cols ...field.Expr) gen.Columns {
	return m.marketHistoryDo.Columns(cols...)
}

func (m *marketHistory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *marketHistory) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 8)
	m.fieldMap["type_id"] = m.TypeID
	m.fieldMap["region_id"] = m.RegionID
	m.fieldMap["date"] = m.Date
	m.fieldMap["average"] = m.Average
	m.fieldMap["highest"] = m.Highest
	m.fieldMap["lowest"] = m.Lowest
	m.fieldMap["order_count"] = m.OrderCount
	m.fieldMap["volume"] = m.Volume
}

func (m marketHistory) clone(db *gorm.DB) marketHistory {
	m.marketHistoryDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m marketHistory) replaceDB(db *gorm.DB) marketHistory {
	m.marketHistoryDo.ReplaceDB(db)
	return m
}

type marketHistoryDo struct{ gen.DO }

func (m marketHistoryDo) Debug() *marketHistoryDo {
	return m.withDO(m.DO.Debug())
}

func (m marketHistoryDo) WithContext(ctx context.Context) *marketHistoryDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m marketHistoryDo) ReadDB() *marketHistoryDo {
	return m.Clauses(dbresolver.Read)
}

func (m marketHistoryDo) WriteDB() *marketHistoryDo {
	return m.Clauses(dbresolver.Write)
}

func (m marketHistoryDo) Session(config *gorm.Session) *marketHistoryDo {
	return m.withDO(m.DO.Session(config))
}

func (m marketHistoryDo) Clauses(conds ...clause.Expression) *marketHistoryDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m marketHistoryDo) Returning(value interface{}, columns ...string) *marketHistoryDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m marketHistoryDo) Not(conds ...gen.Condition) *marketHistoryDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m marketHistoryDo) Or(conds ...gen.Condition) *marketHistoryDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m marketHistoryDo) Select(conds ...field.Expr) *marketHistoryDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m marketHistoryDo) Where(conds ...gen.Condition) *marketHistoryDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m marketHistoryDo) Order(conds ...field.Expr) *marketHistoryDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m marketHistoryDo) Distinct(cols ...field.Expr) *marketHistoryDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m marketHistoryDo) Omit(cols ...field.Expr) *marketHistoryDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m marketHistoryDo) Join(table schema.Tabler, on ...field.Expr) *marketHistoryDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m marketHistoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) *marketHistoryDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m marketHistoryDo) RightJoin(table schema.Tabler, on ...field.Expr) *marketHistoryDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m marketHistoryDo) Group(cols ...field.Expr) *marketHistoryDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m marketHistoryDo) Having(conds ...gen.Condition) *marketHistoryDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m marketHistoryDo) Limit(limit int) *marketHistoryDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m marketHistoryDo) Offset(offset int) *marketHistoryDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m marketHistoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *marketHistoryDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m marketHistoryDo) Unscoped() *marketHistoryDo {
	return m.withDO(m.DO.Unscoped())
}

func (m marketHistoryDo) Create(values ...*model.MarketHistory) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m marketHistoryDo) CreateInBatches(values []*model.MarketHistory, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m marketHistoryDo) Save(values ...*model.MarketHistory) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m marketHistoryDo) First() (*model.MarketHistory, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MarketHistory), nil
	}
}

func (m marketHistoryDo) Take() (*model.MarketHistory, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MarketHistory), nil
	}
}

func (m marketHistoryDo) Last() (*model.MarketHistory, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MarketHistory), nil
	}
}

func (m marketHistoryDo) Find() ([]*model.MarketHistory, error) {
	result, err := m.DO.Find()
	return result.([]*model.MarketHistory), err
}

func (m marketHistoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MarketHistory, err error) {
	buf := make([]*model.MarketHistory, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m marketHistoryDo) FindInBatches(result *[]*model.MarketHistory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m marketHistoryDo) Attrs(attrs ...field.AssignExpr) *marketHistoryDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m marketHistoryDo) Assign(attrs ...field.AssignExpr) *marketHistoryDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m marketHistoryDo) Joins(fields ...field.RelationField) *marketHistoryDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m marketHistoryDo) Preload(fields ...field.RelationField) *marketHistoryDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m marketHistoryDo) FirstOrInit() (*model.MarketHistory, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MarketHistory), nil
	}
}

func (m marketHistoryDo) FirstOrCreate() (*model.MarketHistory, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MarketHistory), nil
	}
}

func (m marketHistoryDo) FindByPage(offset int, limit int) (result []*model.MarketHistory, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m marketHistoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m marketHistoryDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m marketHistoryDo) Delete(models ...*model.MarketHistory) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *marketHistoryDo) withDO(do gen.Dao) *marketHistoryDo {
	m.DO = *do.(*gen.DO)
	return m
}
