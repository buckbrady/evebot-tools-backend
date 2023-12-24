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
	err := db.AutoMigrate(&model.Character{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.Character{}) fail: %s", err)
	}
}

func Test_characterQuery(t *testing.T) {
	character := newCharacter(db)
	character = *character.As(character.TableName())
	_do := character.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(character.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <character> fail:", err)
		return
	}

	_, ok := character.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from character success")
	}

	err = _do.Create(&model.Character{})
	if err != nil {
		t.Error("create item in table <character> fail:", err)
	}

	err = _do.Save(&model.Character{})
	if err != nil {
		t.Error("create item in table <character> fail:", err)
	}

	err = _do.CreateInBatches([]*model.Character{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <character> fail:", err)
	}

	_, err = _do.Select(character.ALL).Take()
	if err != nil {
		t.Error("Take() on table <character> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <character> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <character> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <character> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.Character{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <character> fail:", err)
	}

	_, err = _do.Select(character.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <character> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <character> fail:", err)
	}

	_, err = _do.Select(character.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <character> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <character> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <character> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <character> fail:", err)
	}

	_, err = _do.ScanByPage(&model.Character{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <character> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <character> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <character> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <character> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <character> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <character> fail:", err)
	}
}
