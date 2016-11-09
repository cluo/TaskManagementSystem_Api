package common

import (
	"time"

	"github.com/tidwall/buntdb"
)

var Bunt *BuntDBSessionStruct = &BuntDBSessionStruct{}

type BuntDBSessionStruct struct {
	Db *buntdb.DB
}

func (session *BuntDBSessionStruct) InitDB() (err error) {
	// Open the data.db file. It will be created if it doesn't exist.
	session.Db, err = buntdb.Open(":memory:")
	return
}

func (session *BuntDBSessionStruct) CloseDB() (err error) {
	// Open the data.db file. It will be created if it doesn't exist.
	err = session.Db.Close()
	return
}

func (session *BuntDBSessionStruct) Set(key, value string) (err error) {
	err = session.Db.Update(func(tx *buntdb.Tx) error {
		_, _, err1 := tx.Set(key, value, &buntdb.SetOptions{Expires: true, TTL: time.Second * 3600})
		return err1
	})
	return
}

func (session *BuntDBSessionStruct) Get(key string) (value string, err error) {
	err = session.Db.Update(func(tx *buntdb.Tx) error {
		val, err1 := tx.Get(key)
		if err1 != nil {
			return err1
		}
		_, _, _ = tx.Set(key, val, &buntdb.SetOptions{Expires: true, TTL: time.Second * 3600})
		value = val
		return nil
	})
	return
}
