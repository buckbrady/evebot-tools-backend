// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package database

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                         db,
		Killmail:                   newKillmail(db, opts...),
		KillmailAttacker:           newKillmailAttacker(db, opts...),
		KillmailItem:               newKillmailItem(db, opts...),
		KillmailStat:               newKillmailStat(db, opts...),
		KillmailVictim:             newKillmailVictim(db, opts...),
		ServerStatus:               newServerStatus(db, opts...),
		UniverseAstroidBelt:        newUniverseAstroidBelt(db, opts...),
		UniverseMoon:               newUniverseMoon(db, opts...),
		UniversePlanet:             newUniversePlanet(db, opts...),
		UniverseSystem:             newUniverseSystem(db, opts...),
		UniverseSystemPosition:     newUniverseSystemPosition(db, opts...),
		UniverseType:               newUniverseType(db, opts...),
		UniverseTypeDogmaAttribute: newUniverseTypeDogmaAttribute(db, opts...),
		UniverseTypeDogmaEffect:    newUniverseTypeDogmaEffect(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Killmail                   killmail
	KillmailAttacker           killmailAttacker
	KillmailItem               killmailItem
	KillmailStat               killmailStat
	KillmailVictim             killmailVictim
	ServerStatus               serverStatus
	UniverseAstroidBelt        universeAstroidBelt
	UniverseMoon               universeMoon
	UniversePlanet             universePlanet
	UniverseSystem             universeSystem
	UniverseSystemPosition     universeSystemPosition
	UniverseType               universeType
	UniverseTypeDogmaAttribute universeTypeDogmaAttribute
	UniverseTypeDogmaEffect    universeTypeDogmaEffect
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                         db,
		Killmail:                   q.Killmail.clone(db),
		KillmailAttacker:           q.KillmailAttacker.clone(db),
		KillmailItem:               q.KillmailItem.clone(db),
		KillmailStat:               q.KillmailStat.clone(db),
		KillmailVictim:             q.KillmailVictim.clone(db),
		ServerStatus:               q.ServerStatus.clone(db),
		UniverseAstroidBelt:        q.UniverseAstroidBelt.clone(db),
		UniverseMoon:               q.UniverseMoon.clone(db),
		UniversePlanet:             q.UniversePlanet.clone(db),
		UniverseSystem:             q.UniverseSystem.clone(db),
		UniverseSystemPosition:     q.UniverseSystemPosition.clone(db),
		UniverseType:               q.UniverseType.clone(db),
		UniverseTypeDogmaAttribute: q.UniverseTypeDogmaAttribute.clone(db),
		UniverseTypeDogmaEffect:    q.UniverseTypeDogmaEffect.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                         db,
		Killmail:                   q.Killmail.replaceDB(db),
		KillmailAttacker:           q.KillmailAttacker.replaceDB(db),
		KillmailItem:               q.KillmailItem.replaceDB(db),
		KillmailStat:               q.KillmailStat.replaceDB(db),
		KillmailVictim:             q.KillmailVictim.replaceDB(db),
		ServerStatus:               q.ServerStatus.replaceDB(db),
		UniverseAstroidBelt:        q.UniverseAstroidBelt.replaceDB(db),
		UniverseMoon:               q.UniverseMoon.replaceDB(db),
		UniversePlanet:             q.UniversePlanet.replaceDB(db),
		UniverseSystem:             q.UniverseSystem.replaceDB(db),
		UniverseSystemPosition:     q.UniverseSystemPosition.replaceDB(db),
		UniverseType:               q.UniverseType.replaceDB(db),
		UniverseTypeDogmaAttribute: q.UniverseTypeDogmaAttribute.replaceDB(db),
		UniverseTypeDogmaEffect:    q.UniverseTypeDogmaEffect.replaceDB(db),
	}
}

type queryCtx struct {
	Killmail                   *killmailDo
	KillmailAttacker           *killmailAttackerDo
	KillmailItem               *killmailItemDo
	KillmailStat               *killmailStatDo
	KillmailVictim             *killmailVictimDo
	ServerStatus               *serverStatusDo
	UniverseAstroidBelt        *universeAstroidBeltDo
	UniverseMoon               *universeMoonDo
	UniversePlanet             *universePlanetDo
	UniverseSystem             *universeSystemDo
	UniverseSystemPosition     *universeSystemPositionDo
	UniverseType               *universeTypeDo
	UniverseTypeDogmaAttribute *universeTypeDogmaAttributeDo
	UniverseTypeDogmaEffect    *universeTypeDogmaEffectDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Killmail:                   q.Killmail.WithContext(ctx),
		KillmailAttacker:           q.KillmailAttacker.WithContext(ctx),
		KillmailItem:               q.KillmailItem.WithContext(ctx),
		KillmailStat:               q.KillmailStat.WithContext(ctx),
		KillmailVictim:             q.KillmailVictim.WithContext(ctx),
		ServerStatus:               q.ServerStatus.WithContext(ctx),
		UniverseAstroidBelt:        q.UniverseAstroidBelt.WithContext(ctx),
		UniverseMoon:               q.UniverseMoon.WithContext(ctx),
		UniversePlanet:             q.UniversePlanet.WithContext(ctx),
		UniverseSystem:             q.UniverseSystem.WithContext(ctx),
		UniverseSystemPosition:     q.UniverseSystemPosition.WithContext(ctx),
		UniverseType:               q.UniverseType.WithContext(ctx),
		UniverseTypeDogmaAttribute: q.UniverseTypeDogmaAttribute.WithContext(ctx),
		UniverseTypeDogmaEffect:    q.UniverseTypeDogmaEffect.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
