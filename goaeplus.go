package goaeplus

import (
	"appengine"
	"appengine/datastore"
	"appengine/memcache"
	_ "log"
	"reflect"
	"strings"
)

// go's appengine datastore service returns
// keys and structs separately
// this method uses a string value on any struct
// called Id that will be returned as an encoded
// datastore key

/// ex:
// type User struct {
// 	name  string
// 	email string
// 	Id    string
// }

// options to use memcache layer
// defaults to true

type Options struct {
	UseMemcache bool
}

// this is a layer of abstraction to wrap the datastore
// save function
// ex. err := Save(c, User)
func Save(c appengine.Context, m interface{}, opts *Options) error {

	// check to call beforesave method
	if _, ok := m.(BeforeSaveInterface); ok {
		reflect.ValueOf(m).MethodByName("BeforeSave").Call([]reflect.Value{})
	}

	// store object in datastore
	entityName := strings.Split(reflect.TypeOf(m).String(), ".")[1] //assumes model is in separate package
	entityKey := datastore.NewIncompleteKey(c, entityName, nil)
	id := entityKey.Encode()
	// set key as Id value on object
	reflect.ValueOf(m).Elem().FieldByName("Id").SetString(id)
	_, err := datastore.Put(c, entityKey, m)
	if err != nil {
		return err
	}

	// if opts.UseMemcache != false {
	err = setMemcache(m, c)
	if err != nil {
		return err
	}
	// }

	// check to call aftersave callback
	if _, ok := m.(AfterSaveInterface); ok {
		reflect.ValueOf(m).MethodByName("AfterSave").Call([]reflect.Value{})
	}

	return nil
}

func Update(c appengine.Context, m interface{}) error {
	// check to call beforeupdate method
	if _, ok := m.(BeforeUpdateInterface); ok {
		reflect.ValueOf(m).MethodByName("BeforeUpdate").Call([]reflect.Value{})
	}

	id := reflect.ValueOf(m).Elem().FieldByName("Id")

	entityKey, err := datastore.DecodeKey(id.String())
	if err != nil {
		return err
	}

	_, err = datastore.Put(c, entityKey, m)
	if err != nil {
		return err
	}

	// check to call afterupdate method
	if _, ok := m.(AfterUpdateInterface); ok {
		reflect.ValueOf(m).MethodByName("AfterUpdate").Call([]reflect.Value{})
	}

	return nil
}

func Get(c appengine.Context, id string, entity interface{}) error {
	// get from memcache
	if _, err := memcache.Gob.Get(c, id, entity); err == nil {
		c.Infof("testing memcache hit %s", entity)
		return nil
	} else {
		// fetch from datastore
		key, err := datastore.DecodeKey(id)

		if err != nil {
			return err
		}

		err = datastore.Get(c, key, entity)
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}
