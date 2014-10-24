package goaeplus

import (
	"appengine"
	"appengine/memcache"
	"reflect"
)

// use Gob encoder for memcache

// fetch item from memcache using the
// item's id to get it

func getMemcache(m interface{}, key string, c appengine.Context) (interface{}, error) {
	_, err := memcache.Gob.Get(c, key, m)
	if err != nil {
		return m, err
	}
	return m, nil
}

// set the item in memcache
// the key will be the encoded
// datastore key as a string

func setMemcache(m interface{}, c appengine.Context) error {

	id := reflect.ValueOf(m).Elem().FieldByName("Id")
	idToString := id.String()

	item := &memcache.Item{
		Key:    idToString,
		Object: m,
	}
	err := memcache.Gob.Set(c, item)
	if err != nil {
		return err
	}

	return nil
}
