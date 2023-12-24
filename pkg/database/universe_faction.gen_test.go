// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package database

import (
	"context"
	"fmt"
	"testing"

	"github.com/buckbrady/evebot-tools-backend/pkg/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&model.UniverseFaction{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.UniverseFaction{}) fail: %s", err)
	}
}

func Test_universeFactionQuery(t *testing.T) {
	universeFaction := newUniverseFaction(db)
	universeFaction = *universeFaction.As(universeFaction.TableName())
	_do := universeFaction.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(universeFaction.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <universe_faction> fail:", err)
		return
	}

	_, ok := universeFaction.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from universeFaction success")
	}

	err = _do.Create(&model.UniverseFaction{})
	if err != nil {
		t.Error("create item in table <universe_faction> fail:", err)
	}

	err = _do.Save(&model.UniverseFaction{})
	if err != nil {
		t.Error("create item in table <universe_faction> fail:", err)
	}

	err = _do.CreateInBatches([]*model.UniverseFaction{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <universe_faction> fail:", err)
	}

	_, err = _do.Select(universeFaction.ALL).Take()
	if err != nil {
		t.Error("Take() on table <universe_faction> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <universe_faction> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <universe_faction> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <universe_faction> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.UniverseFaction{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <universe_faction> fail:", err)
	}

	_, err = _do.Select(universeFaction.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <universe_faction> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <universe_faction> fail:", err)
	}

	_, err = _do.Select(universeFaction.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <universe_faction> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <universe_faction> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <universe_faction> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <universe_faction> fail:", err)
	}

	_, err = _do.ScanByPage(&model.UniverseFaction{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <universe_faction> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <universe_faction> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <universe_faction> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <universe_faction> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <universe_faction> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <universe_faction> fail:", err)
	}
}
