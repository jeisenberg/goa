package goaeplus

import (
	"appengine"
	"appengine/memcache"
	"errors"
	"reflect"
)

// use Gob encoder for memcache

func getMemcache(m interface{}, key string, c appengine.Context) (interface{}, error) {
	_, err := memcache.Gob.Get(c, key, m)
	if err != nil {
		return m, err
	}
	return m, nil
}

func setMemcache(m interface{}, c appengine.Context) error {

	id := reflect.ValueOf(m).Elem().FieldByName("Id")

	if id.(string) {
		item := &memache.Item{
			Key:    id,
			Object: m,
		}
		err := memcache.Gob.Set(c, item)
		if err != nil {
			return err
		}
	} else {
		return errors.New("No id value present")
	}
	return nil
}
