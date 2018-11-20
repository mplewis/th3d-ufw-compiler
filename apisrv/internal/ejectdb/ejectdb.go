package ejectdb

import (
	"../structs"
	"github.com/golang-collections/go-datastructures/queue"
)

// EjectingDB acts as a Map that ejects the oldest items above a specified size cap.
type EjectingDB interface {
	Get(key string) (structs.CompileResult, bool)
	Put(result structs.CompileResult)
}

// EjectDB is the internal structure that powers an EjectingDB.
type EjectDB struct {
	Max     int64
	Results map[string]structs.CompileResult
	log     *queue.Queue
}

// Get retrieves a result for a key, returning the optional result and false if the result was not found.
func (db EjectDB) Get(key string) (structs.CompileResult, bool) {
	val, ok := db.Results[key]
	return val, ok
}

// Put inserts a result for a key.
func (db EjectDB) Put(result structs.CompileResult) {
	key := result.Request.ID
	db.log.Put(key)
	for db.log.Len() > db.Max {
		delKeys, err := db.log.Get(1)
		if err != nil {
			panic(err)
		}
		delKey := delKeys[0].(string)
		delete(db.Results, delKey)
	}
	db.Results[key] = result
}

// NewEjectDB creates an empty EjectDB with a fixed size.
func NewEjectDB(max int64) EjectingDB {
	return EjectDB{
		Max:     max,
		Results: map[string]structs.CompileResult{},
		log:     queue.New(max),
	}
}
