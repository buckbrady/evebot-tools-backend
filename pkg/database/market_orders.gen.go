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

func newMarketOrder(db *gorm.DB, opts ...gen.DOOption) marketOrder {
	_marketOrder := marketOrder{}

	_marketOrder.marketOrderDo.UseDB(db, opts...)
	_marketOrder.marketOrderDo.UseModel(&model.MarketOrder{})

	tableName := _marketOrder.marketOrderDo.TableName()
	_marketOrder.ALL = field.NewAsterisk(tableName)
	_marketOrder.ID = field.NewInt64(tableName, "id")
	_marketOrder.TypeID = field.NewInt32(tableName, "type_id")
	_marketOrder.RegionID = field.NewInt32(tableName, "region_id")
	_marketOrder.LocationID = field.NewInt64(tableName, "location_id")
	_marketOrder.Range = field.NewString(tableName, "range")
	_marketOrder.IsBuyOrder = field.NewBool(tableName, "is_buy_order")
	_marketOrder.Price = field.NewFloat64(tableName, "price")
	_marketOrder.VolumeTotal = field.NewInt32(tableName, "volume_total")
	_marketOrder.VolumeRemain = field.NewInt32(tableName, "volume_remain")
	_marketOrder.Issued = field.NewTime(tableName, "issued")
	_marketOrder.Duration = field.NewInt32(tableName, "duration")
	_marketOrder.MinVolume = field.NewInt32(tableName, "min_volume")
	_marketOrder.LastUpdated = field.NewTime(tableName, "last_updated")
	_marketOrder.IsActive = field.NewBool(tableName, "is_active")

	_marketOrder.fillFieldMap()

	return _marketOrder
}

type marketOrder struct {
	marketOrderDo marketOrderDo

	ALL          field.Asterisk
	ID           field.Int64
	TypeID       field.Int32
	RegionID     field.Int32
	LocationID   field.Int64
	Range        field.String
	IsBuyOrder   field.Bool
	Price        field.Float64
	VolumeTotal  field.Int32
	VolumeRemain field.Int32
	Issued       field.Time
	Duration     field.Int32
	MinVolume    field.Int32
	LastUpdated  field.Time
	IsActive     field.Bool

	fieldMap map[string]field.Expr
}

func (m marketOrder) Table(newTableName string) *marketOrder {
	m.marketOrderDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m marketOrder) As(alias string) *marketOrder {
	m.marketOrderDo.DO = *(m.marketOrderDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *marketOrder) updateTableName(table string) *marketOrder {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt64(table, "id")
	m.TypeID = field.NewInt32(table, "type_id")
	m.RegionID = field.NewInt32(table, "region_id")
	m.LocationID = field.NewInt64(table, "location_id")
	m.Range = field.NewString(table, "range")
	m.IsBuyOrder = field.NewBool(table, "is_buy_order")
	m.Price = field.NewFloat64(table, "price")
	m.VolumeTotal = field.NewInt32(table, "volume_total")
	m.VolumeRemain = field.NewInt32(table, "volume_remain")
	m.Issued = field.NewTime(table, "issued")
	m.Duration = field.NewInt32(table, "duration")
	m.MinVolume = field.NewInt32(table, "min_volume")
	m.LastUpdated = field.NewTime(table, "last_updated")
	m.IsActive = field.NewBool(table, "is_active")

	m.fillFieldMap()

	return m
}

func (m *marketOrder) WithContext(ctx context.Context) *marketOrderDo {
	return m.marketOrderDo.WithContext(ctx)
}

func (m marketOrder) TableName() string { return m.marketOrderDo.TableName() }

func (m marketOrder) Alias() string { return m.marketOrderDo.Alias() }

func (m marketOrder) Columns(cols ...field.Expr) gen.Columns { return m.marketOrderDo.Columns(cols...) }

func (m *marketOrder) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *marketOrder) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 14)
	m.fieldMap["id"] = m.ID
	m.fieldMap["type_id"] = m.TypeID
	m.fieldMap["region_id"] = m.RegionID
	m.fieldMap["location_id"] = m.LocationID
	m.fieldMap["range"] = m.Range
	m.fieldMap["is_buy_order"] = m.IsBuyOrder
	m.fieldMap["price"] = m.Price
	m.fieldMap["volume_total"] = m.VolumeTotal
	m.fieldMap["volume_remain"] = m.VolumeRemain
	m.fieldMap["issued"] = m.Issued
	m.fieldMap["duration"] = m.Duration
	m.fieldMap["min_volume"] = m.MinVolume
	m.fieldMap["last_updated"] = m.LastUpdated
	m.fieldMap["is_active"] = m.IsActive
}

func (m marketOrder) clone(db *gorm.DB) marketOrder {
	m.marketOrderDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m marketOrder) replaceDB(db *gorm.DB) marketOrder {
	m.marketOrderDo.ReplaceDB(db)
	return m
}

type marketOrderDo struct{ gen.DO }

func (m marketOrderDo) Debug() *marketOrderDo {
	return m.withDO(m.DO.Debug())
}

func (m marketOrderDo) WithContext(ctx context.Context) *marketOrderDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m marketOrderDo) ReadDB() *marketOrderDo {
	return m.Clauses(dbresolver.Read)
}

func (m marketOrderDo) WriteDB() *marketOrderDo {
	return m.Clauses(dbresolver.Write)
}

func (m marketOrderDo) Session(config *gorm.Session) *marketOrderDo {
	return m.withDO(m.DO.Session(config))
}

func (m marketOrderDo) Clauses(conds ...clause.Expression) *marketOrderDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m marketOrderDo) Returning(value interface{}, columns ...string) *marketOrderDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m marketOrderDo) Not(conds ...gen.Condition) *marketOrderDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m marketOrderDo) Or(conds ...gen.Condition) *marketOrderDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m marketOrderDo) Select(conds ...field.Expr) *marketOrderDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m marketOrderDo) Where(conds ...gen.Condition) *marketOrderDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m marketOrderDo) Order(conds ...field.Expr) *marketOrderDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m marketOrderDo) Distinct(cols ...field.Expr) *marketOrderDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m marketOrderDo) Omit(cols ...field.Expr) *marketOrderDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m marketOrderDo) Join(table schema.Tabler, on ...field.Expr) *marketOrderDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m marketOrderDo) LeftJoin(table schema.Tabler, on ...field.Expr) *marketOrderDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m marketOrderDo) RightJoin(table schema.Tabler, on ...field.Expr) *marketOrderDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m marketOrderDo) Group(cols ...field.Expr) *marketOrderDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m marketOrderDo) Having(conds ...gen.Condition) *marketOrderDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m marketOrderDo) Limit(limit int) *marketOrderDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m marketOrderDo) Offset(offset int) *marketOrderDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m marketOrderDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *marketOrderDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m marketOrderDo) Unscoped() *marketOrderDo {
	return m.withDO(m.DO.Unscoped())
}

func (m marketOrderDo) Create(values ...*model.MarketOrder) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m marketOrderDo) CreateInBatches(values []*model.MarketOrder, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m marketOrderDo) Save(values ...*model.MarketOrder) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m marketOrderDo) First() (*model.MarketOrder, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MarketOrder), nil
	}
}

func (m marketOrderDo) Take() (*model.MarketOrder, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MarketOrder), nil
	}
}

func (m marketOrderDo) Last() (*model.MarketOrder, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MarketOrder), nil
	}
}

func (m marketOrderDo) Find() ([]*model.MarketOrder, error) {
	result, err := m.DO.Find()
	return result.([]*model.MarketOrder), err
}

func (m marketOrderDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MarketOrder, err error) {
	buf := make([]*model.MarketOrder, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m marketOrderDo) FindInBatches(result *[]*model.MarketOrder, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m marketOrderDo) Attrs(attrs ...field.AssignExpr) *marketOrderDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m marketOrderDo) Assign(attrs ...field.AssignExpr) *marketOrderDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m marketOrderDo) Joins(fields ...field.RelationField) *marketOrderDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m marketOrderDo) Preload(fields ...field.RelationField) *marketOrderDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m marketOrderDo) FirstOrInit() (*model.MarketOrder, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MarketOrder), nil
	}
}

func (m marketOrderDo) FirstOrCreate() (*model.MarketOrder, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MarketOrder), nil
	}
}

func (m marketOrderDo) FindByPage(offset int, limit int) (result []*model.MarketOrder, count int64, err error) {
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

func (m marketOrderDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m marketOrderDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m marketOrderDo) Delete(models ...*model.MarketOrder) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *marketOrderDo) withDO(do gen.Dao) *marketOrderDo {
	m.DO = *do.(*gen.DO)
	return m
}
