// Code generated by 'micro gen' command.
// DO NOT EDIT!

package model

import (
	"time"
	"unsafe"

	"github.com/henrylee2cn/erpc/v6"
	"github.com/henrylee2cn/goutil/coarsetime"
	"github.com/xiaoenai/tp-micro/v6/model/mongo"

	"github.com/xiaoenai/tp-micro/v6/examples/project/args"
)

var _ = erpc.Errorf

// Meta comment...
type Meta args.Meta

// ToMeta converts to *Meta type.
func ToMeta(_m *args.Meta) *Meta {
	return (*Meta)(unsafe.Pointer(_m))
}

// ToArgsMeta converts to *args.Meta type.
func ToArgsMeta(_m *Meta) *args.Meta {
	return (*args.Meta)(unsafe.Pointer(_m))
}

// TableName implements 'github.com/xiaoenai/tp-micro/model'.Cacheable
func (*Meta) TableName() string {
	return "meta"
}

func (_m *Meta) isZeroPrimaryKey() bool {
	var __id mongo.ObjectId
	if _m.Id != __id {
		return false
	}
	return true
}

var metaDB, _ = mongoHandler.RegCacheableDB(new(Meta), time.Hour*24)

// GetMetaDB returns the Meta DB handler.
func GetMetaDB() *mongo.CacheableDB {
	return metaDB
}

// UpsertMeta insert or update the Meta data by selector and updater.
// NOTE:
//  With cache layer;
//  Insert data if the primary key is specified;
//  Update data based on _updateFields if no primary key is specified;
func UpsertMeta(selector, updater mongo.M) error {
	updater["updated_at"] = coarsetime.FloorTimeNow().Unix()
	return metaDB.WitchCollection(func(col *mongo.Collection) error {
		_, err := col.Upsert(selector, mongo.M{"$set": updater})
		return err
	})
}

// GetMetaByFields query a Meta data from database by WHERE field.
// NOTE:
//  With cache layer;
//  If @return error!=nil, means the database error.
func GetMetaByFields(_m *Meta, _fields ...string) (bool, error) {
	err := metaDB.CacheGet(_m, _fields...)
	switch err {
	case nil:
		return true, nil
	case mongo.ErrNotFound:
		return false, nil
	default:
		return false, err
	}
}

// GetMetaByWhere query a Meta data from database by WHERE condition.
// NOTE:
//  Without cache layer;
//  If @return error!=nil, means the database error.
func GetMetaByWhere(query mongo.M) (*Meta, bool, error) {
	_m := &Meta{}
	err := metaDB.WitchCollection(func(col *mongo.Collection) error {
		return col.Find(query).One(_m)
	})
	switch err {
	case nil:
		return _m, true, nil
	case mongo.ErrNotFound:
		return nil, false, nil
	default:
		return nil, false, err
	}
}
