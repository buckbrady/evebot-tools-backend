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

func newCharacter(db *gorm.DB, opts ...gen.DOOption) character {
	_character := character{}

	_character.characterDo.UseDB(db, opts...)
	_character.characterDo.UseModel(&model.Character{})

	tableName := _character.characterDo.TableName()
	_character.ALL = field.NewAsterisk(tableName)
	_character.ID = field.NewInt32(tableName, "id")
	_character.AllianceID = field.NewInt32(tableName, "alliance_id")
	_character.Birthday = field.NewTime(tableName, "birthday")
	_character.BloodlineID = field.NewInt32(tableName, "bloodline_id")
	_character.CorporationID = field.NewInt32(tableName, "corporation_id")
	_character.Description = field.NewString(tableName, "description")
	_character.FactionID = field.NewInt32(tableName, "faction_id")
	_character.Gender = field.NewString(tableName, "gender")
	_character.Name = field.NewString(tableName, "name")
	_character.RaceID = field.NewInt32(tableName, "race_id")
	_character.SecurityStatus = field.NewFloat64(tableName, "security_status")
	_character.Title = field.NewString(tableName, "title")

	_character.fillFieldMap()

	return _character
}

type character struct {
	characterDo characterDo

	ALL            field.Asterisk
	ID             field.Int32
	AllianceID     field.Int32
	Birthday       field.Time
	BloodlineID    field.Int32
	CorporationID  field.Int32
	Description    field.String
	FactionID      field.Int32
	Gender         field.String
	Name           field.String
	RaceID         field.Int32
	SecurityStatus field.Float64
	Title          field.String

	fieldMap map[string]field.Expr
}

func (c character) Table(newTableName string) *character {
	c.characterDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c character) As(alias string) *character {
	c.characterDo.DO = *(c.characterDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *character) updateTableName(table string) *character {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt32(table, "id")
	c.AllianceID = field.NewInt32(table, "alliance_id")
	c.Birthday = field.NewTime(table, "birthday")
	c.BloodlineID = field.NewInt32(table, "bloodline_id")
	c.CorporationID = field.NewInt32(table, "corporation_id")
	c.Description = field.NewString(table, "description")
	c.FactionID = field.NewInt32(table, "faction_id")
	c.Gender = field.NewString(table, "gender")
	c.Name = field.NewString(table, "name")
	c.RaceID = field.NewInt32(table, "race_id")
	c.SecurityStatus = field.NewFloat64(table, "security_status")
	c.Title = field.NewString(table, "title")

	c.fillFieldMap()

	return c
}

func (c *character) WithContext(ctx context.Context) *characterDo {
	return c.characterDo.WithContext(ctx)
}

func (c character) TableName() string { return c.characterDo.TableName() }

func (c character) Alias() string { return c.characterDo.Alias() }

func (c character) Columns(cols ...field.Expr) gen.Columns { return c.characterDo.Columns(cols...) }

func (c *character) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *character) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 12)
	c.fieldMap["id"] = c.ID
	c.fieldMap["alliance_id"] = c.AllianceID
	c.fieldMap["birthday"] = c.Birthday
	c.fieldMap["bloodline_id"] = c.BloodlineID
	c.fieldMap["corporation_id"] = c.CorporationID
	c.fieldMap["description"] = c.Description
	c.fieldMap["faction_id"] = c.FactionID
	c.fieldMap["gender"] = c.Gender
	c.fieldMap["name"] = c.Name
	c.fieldMap["race_id"] = c.RaceID
	c.fieldMap["security_status"] = c.SecurityStatus
	c.fieldMap["title"] = c.Title
}

func (c character) clone(db *gorm.DB) character {
	c.characterDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c character) replaceDB(db *gorm.DB) character {
	c.characterDo.ReplaceDB(db)
	return c
}

type characterDo struct{ gen.DO }

func (c characterDo) Debug() *characterDo {
	return c.withDO(c.DO.Debug())
}

func (c characterDo) WithContext(ctx context.Context) *characterDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c characterDo) ReadDB() *characterDo {
	return c.Clauses(dbresolver.Read)
}

func (c characterDo) WriteDB() *characterDo {
	return c.Clauses(dbresolver.Write)
}

func (c characterDo) Session(config *gorm.Session) *characterDo {
	return c.withDO(c.DO.Session(config))
}

func (c characterDo) Clauses(conds ...clause.Expression) *characterDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c characterDo) Returning(value interface{}, columns ...string) *characterDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c characterDo) Not(conds ...gen.Condition) *characterDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c characterDo) Or(conds ...gen.Condition) *characterDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c characterDo) Select(conds ...field.Expr) *characterDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c characterDo) Where(conds ...gen.Condition) *characterDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c characterDo) Order(conds ...field.Expr) *characterDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c characterDo) Distinct(cols ...field.Expr) *characterDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c characterDo) Omit(cols ...field.Expr) *characterDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c characterDo) Join(table schema.Tabler, on ...field.Expr) *characterDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c characterDo) LeftJoin(table schema.Tabler, on ...field.Expr) *characterDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c characterDo) RightJoin(table schema.Tabler, on ...field.Expr) *characterDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c characterDo) Group(cols ...field.Expr) *characterDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c characterDo) Having(conds ...gen.Condition) *characterDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c characterDo) Limit(limit int) *characterDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c characterDo) Offset(offset int) *characterDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c characterDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *characterDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c characterDo) Unscoped() *characterDo {
	return c.withDO(c.DO.Unscoped())
}

func (c characterDo) Create(values ...*model.Character) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c characterDo) CreateInBatches(values []*model.Character, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c characterDo) Save(values ...*model.Character) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c characterDo) First() (*model.Character, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Character), nil
	}
}

func (c characterDo) Take() (*model.Character, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Character), nil
	}
}

func (c characterDo) Last() (*model.Character, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Character), nil
	}
}

func (c characterDo) Find() ([]*model.Character, error) {
	result, err := c.DO.Find()
	return result.([]*model.Character), err
}

func (c characterDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Character, err error) {
	buf := make([]*model.Character, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c characterDo) FindInBatches(result *[]*model.Character, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c characterDo) Attrs(attrs ...field.AssignExpr) *characterDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c characterDo) Assign(attrs ...field.AssignExpr) *characterDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c characterDo) Joins(fields ...field.RelationField) *characterDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c characterDo) Preload(fields ...field.RelationField) *characterDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c characterDo) FirstOrInit() (*model.Character, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Character), nil
	}
}

func (c characterDo) FirstOrCreate() (*model.Character, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Character), nil
	}
}

func (c characterDo) FindByPage(offset int, limit int) (result []*model.Character, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c characterDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c characterDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c characterDo) Delete(models ...*model.Character) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *characterDo) withDO(do gen.Dao) *characterDo {
	c.DO = *do.(*gen.DO)
	return c
}
