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
	err := db.AutoMigrate(&model.MarketPrice{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.MarketPrice{}) fail: %s", err)
	}
}

func Test_marketPriceQuery(t *testing.T) {
	marketPrice := newMarketPrice(db)
	marketPrice = *marketPrice.As(marketPrice.TableName())
	_do := marketPrice.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(marketPrice.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <market_price> fail:", err)
		return
	}

	_, ok := marketPrice.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from marketPrice success")
	}

	err = _do.Create(&model.MarketPrice{})
	if err != nil {
		t.Error("create item in table <market_price> fail:", err)
	}

	err = _do.Save(&model.MarketPrice{})
	if err != nil {
		t.Error("create item in table <market_price> fail:", err)
	}

	err = _do.CreateInBatches([]*model.MarketPrice{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <market_price> fail:", err)
	}

	_, err = _do.Select(marketPrice.ALL).Take()
	if err != nil {
		t.Error("Take() on table <market_price> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <market_price> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <market_price> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <market_price> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.MarketPrice{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <market_price> fail:", err)
	}

	_, err = _do.Select(marketPrice.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <market_price> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <market_price> fail:", err)
	}

	_, err = _do.Select(marketPrice.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <market_price> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <market_price> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <market_price> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <market_price> fail:", err)
	}

	_, err = _do.ScanByPage(&model.MarketPrice{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <market_price> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <market_price> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <market_price> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <market_price> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <market_price> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <market_price> fail:", err)
	}
}
