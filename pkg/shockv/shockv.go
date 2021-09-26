package shockv

import (
	"path/filepath"
	"sync"
	"time"

	"github.com/dgraph-io/badger/v3"
)

const (
	databaseDir   = "_databases"
	controllerDir = "controller"
)

var (
	controller *badger.DB
	db         *ShockV
	gen        sync.Once
)

type ShockV struct {
	dbs map[string]*badger.DB
}

func Get() (*ShockV, error) {
	var err error
	gen.Do(func() {
		db, err = _new()
		if err != nil {
			panic(err)
		}
	})
	return db, err
}

func (db *ShockV) NewDiskDB(databaseName string) error {
	dir := filepath.FromSlash(databaseDir + "/" + databaseName)
	b, err := badger.Open(badger.DefaultOptions(dir))
	if err != nil {
		return err
	}
	db.dbs[databaseName] = b
	return set(controller, databaseName, "")
}

func (db *ShockV) NewDisklessDB(databaseName string) error {
	b, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		return err
	}
	db.dbs[databaseName] = b
	return nil
}

func (db *ShockV) ListDB() []string {
	var names []string
	for k := range db.dbs {
		names = append(names, k)
	}
	return names
}

func (db *ShockV) Get(databaseName string, key string) (string, error) {
	return get(db.dbs[databaseName], key)
}

func (db *ShockV) Set(databaseName string, key string, value string) error {
	return set(db.dbs[databaseName], key, value)
}

func (db *ShockV) List(databaseName string) (map[string]string, error) {
	return list(db.dbs[databaseName])
}

func (db *ShockV) Delete(databaseName string, key string) error {
	return _delete(db.dbs[databaseName], key)
}

func _new() (*ShockV, error) {
	var err error
	sdb := &ShockV{
		dbs: map[string]*badger.DB{},
	}
	controller, err = badger.Open(badger.DefaultOptions(filepath.FromSlash(databaseDir + "/" + controllerDir)))
	if err != nil {
		return nil, err
	}
	kv, err := list(controller)
	if err != nil {
		return nil, err
	}
	for k := range kv {
		err = sdb.NewDiskDB(k)
		if err != nil {
			return nil, err
		}
	}
	return sdb, nil
}

func set(db *badger.DB, key string, value string) error {
	bKey := []byte(key)
	bValue := []byte(value)
	return db.Update(func(txn *badger.Txn) error {
		return txn.Set(bKey, bValue)
	})
}

func setWithTTL(db *badger.DB, key string, value string, ttl time.Duration) error {
	bKey := []byte(key)
	bValue := []byte(value)
	return db.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry(bKey, bValue).WithTTL(ttl)
		return txn.SetEntry(entry)
	})
}

func list(db *badger.DB) (map[string]string, error) {
	kv := map[string]string{}
	err := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			err := item.Value(func(val []byte) error {
				kv[string(item.Key())] = string(val)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return kv, nil
}

func get(db *badger.DB, key string) (string, error) {
	bID := []byte(key)
	var value string
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(bID)
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			value = string(val)
			return nil
		})
	})
	if err != nil {
		return "", err
	}
	return value, nil
}

func _delete(db *badger.DB, key string) error {
	bID := []byte(key)
	return db.Update(func(txn *badger.Txn) error {
		return txn.Delete(bID)
	})
}
