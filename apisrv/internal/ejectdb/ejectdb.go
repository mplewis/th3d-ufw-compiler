package ejectdb

import (
	"../structs"
	"github.com/golang-collections/go-datastructures/queue"
)

// EjectingDB acts as a Map that ejects the oldest items above a specified size cap.
type EjectingDB interface {
	Put(key string, result structs.CompileResult)
}

// EjectDB is the internal structure that powers an EjectingDB.
type EjectDB struct {
	Max     int64
	Results map[string]structs.CompileResult
	log     *queue.Queue
}

// Put inserts a result for a key.
func (db EjectDB) Put(key string, result structs.CompileResult) {
	db.log.Put(key)
	for db.log.Len() > db.Max {
		delKeys, err := db.log.Get(1)
		delKey := delKeys[0].(string)
		if err != nil {
			panic(err)
		}
		delete(db.Results, delKey)
	}
	db.Results[key] = result
}

// NewEjectDB creates an empty EjectDB with a fixed size.
func NewEjectDB(max int64) EjectDB {
	return EjectDB{
		Max:     max,
		Results: map[string]structs.CompileResult{},
		log:     queue.New(max),
	}
}
