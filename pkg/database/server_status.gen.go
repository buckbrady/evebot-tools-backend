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

func newServerStatus(db *gorm.DB, opts ...gen.DOOption) serverStatus {
	_serverStatus := serverStatus{}

	_serverStatus.serverStatusDo.UseDB(db, opts...)
	_serverStatus.serverStatusDo.UseModel(&model.ServerStatus{})

	tableName := _serverStatus.serverStatusDo.TableName()
	_serverStatus.ALL = field.NewAsterisk(tableName)
	_serverStatus.ID = field.NewInt64(tableName, "id")
	_serverStatus.Players = field.NewInt32(tableName, "players")
	_serverStatus.ServerVersion = field.NewString(tableName, "server_version")
	_serverStatus.StartTime = field.NewTime(tableName, "start_time")
	_serverStatus.Vip = field.NewBool(tableName, "vip")

	_serverStatus.fillFieldMap()

	return _serverStatus
}

type serverStatus struct {
	serverStatusDo serverStatusDo

	ALL           field.Asterisk
	ID            field.Int64
	Players       field.Int32
	ServerVersion field.String
	StartTime     field.Time
	Vip           field.Bool

	fieldMap map[string]field.Expr
}

func (s serverStatus) Table(newTableName string) *serverStatus {
	s.serverStatusDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s serverStatus) As(alias string) *serverStatus {
	s.serverStatusDo.DO = *(s.serverStatusDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *serverStatus) updateTableName(table string) *serverStatus {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.Players = field.NewInt32(table, "players")
	s.ServerVersion = field.NewString(table, "server_version")
	s.StartTime = field.NewTime(table, "start_time")
	s.Vip = field.NewBool(table, "vip")

	s.fillFieldMap()

	return s
}

func (s *serverStatus) WithContext(ctx context.Context) *serverStatusDo {
	return s.serverStatusDo.WithContext(ctx)
}

func (s serverStatus) TableName() string { return s.serverStatusDo.TableName() }

func (s serverStatus) Alias() string { return s.serverStatusDo.Alias() }

func (s serverStatus) Columns(cols ...field.Expr) gen.Columns {
	return s.serverStatusDo.Columns(cols...)
}

func (s *serverStatus) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *serverStatus) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 5)
	s.fieldMap["id"] = s.ID
	s.fieldMap["players"] = s.Players
	s.fieldMap["server_version"] = s.ServerVersion
	s.fieldMap["start_time"] = s.StartTime
	s.fieldMap["vip"] = s.Vip
}

func (s serverStatus) clone(db *gorm.DB) serverStatus {
	s.serverStatusDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s serverStatus) replaceDB(db *gorm.DB) serverStatus {
	s.serverStatusDo.ReplaceDB(db)
	return s
}

type serverStatusDo struct{ gen.DO }

func (s serverStatusDo) Debug() *serverStatusDo {
	return s.withDO(s.DO.Debug())
}

func (s serverStatusDo) WithContext(ctx context.Context) *serverStatusDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s serverStatusDo) ReadDB() *serverStatusDo {
	return s.Clauses(dbresolver.Read)
}

func (s serverStatusDo) WriteDB() *serverStatusDo {
	return s.Clauses(dbresolver.Write)
}

func (s serverStatusDo) Session(config *gorm.Session) *serverStatusDo {
	return s.withDO(s.DO.Session(config))
}

func (s serverStatusDo) Clauses(conds ...clause.Expression) *serverStatusDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s serverStatusDo) Returning(value interface{}, columns ...string) *serverStatusDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s serverStatusDo) Not(conds ...gen.Condition) *serverStatusDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s serverStatusDo) Or(conds ...gen.Condition) *serverStatusDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s serverStatusDo) Select(conds ...field.Expr) *serverStatusDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s serverStatusDo) Where(conds ...gen.Condition) *serverStatusDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s serverStatusDo) Order(conds ...field.Expr) *serverStatusDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s serverStatusDo) Distinct(cols ...field.Expr) *serverStatusDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s serverStatusDo) Omit(cols ...field.Expr) *serverStatusDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s serverStatusDo) Join(table schema.Tabler, on ...field.Expr) *serverStatusDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s serverStatusDo) LeftJoin(table schema.Tabler, on ...field.Expr) *serverStatusDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s serverStatusDo) RightJoin(table schema.Tabler, on ...field.Expr) *serverStatusDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s serverStatusDo) Group(cols ...field.Expr) *serverStatusDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s serverStatusDo) Having(conds ...gen.Condition) *serverStatusDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s serverStatusDo) Limit(limit int) *serverStatusDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s serverStatusDo) Offset(offset int) *serverStatusDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s serverStatusDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *serverStatusDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s serverStatusDo) Unscoped() *serverStatusDo {
	return s.withDO(s.DO.Unscoped())
}

func (s serverStatusDo) Create(values ...*model.ServerStatus) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s serverStatusDo) CreateInBatches(values []*model.ServerStatus, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s serverStatusDo) Save(values ...*model.ServerStatus) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s serverStatusDo) First() (*model.ServerStatus, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerStatus), nil
	}
}

func (s serverStatusDo) Take() (*model.ServerStatus, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerStatus), nil
	}
}

func (s serverStatusDo) Last() (*model.ServerStatus, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerStatus), nil
	}
}

func (s serverStatusDo) Find() ([]*model.ServerStatus, error) {
	result, err := s.DO.Find()
	return result.([]*model.ServerStatus), err
}

func (s serverStatusDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ServerStatus, err error) {
	buf := make([]*model.ServerStatus, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s serverStatusDo) FindInBatches(result *[]*model.ServerStatus, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s serverStatusDo) Attrs(attrs ...field.AssignExpr) *serverStatusDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s serverStatusDo) Assign(attrs ...field.AssignExpr) *serverStatusDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s serverStatusDo) Joins(fields ...field.RelationField) *serverStatusDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s serverStatusDo) Preload(fields ...field.RelationField) *serverStatusDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s serverStatusDo) FirstOrInit() (*model.ServerStatus, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerStatus), nil
	}
}

func (s serverStatusDo) FirstOrCreate() (*model.ServerStatus, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerStatus), nil
	}
}

func (s serverStatusDo) FindByPage(offset int, limit int) (result []*model.ServerStatus, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s serverStatusDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s serverStatusDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s serverStatusDo) Delete(models ...*model.ServerStatus) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *serverStatusDo) withDO(do gen.Dao) *serverStatusDo {
	s.DO = *do.(*gen.DO)
	return s
}
