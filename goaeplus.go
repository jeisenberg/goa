package goaeplus

import (
	"appengine"
	"appengine/datastore"
	_ "log"
	"reflect"
	"strings"
)

// these are interfaces used to check
// if an object implements a specific
// callback method
// if it does, the method will be called
type BeforeSaveInterface interface {
	BeforeSave()
}

type AfterSaveInterface interface {
	AfterSave()
}

// go's appengine datastore service returns
// keys and structs separately
// an entity is a map where the key is the
// datastore key of the entity, and the
// value is the entity data
// this keeps the object together in one place

// this is a layer of abstraction to wrap the datastore
// save function
// ex. err := Save(c, User)
func Save(c appengine.Context, m interface{}) error {

	// check to call beforesave method
	if _, ok := m.(BeforeSaveInterface); ok {
		reflect.ValueOf(m).MethodByName("BeforeSave").Call([]reflect.Value{})
	}

	// store object in datastore
	entityName := strings.Split(reflect.TypeOf(m).String(), ".")[1] //assumes model is in separate package
	entityKey := datastore.NewIncompleteKey(c, entityName, nil)
	_, err := datastore.Put(c, entityKey, m)
	if err != nil {
		return err
	}

	// check to call aftersave callback
	if _, ok := m.(AfterSaveInterface); ok {
		reflect.ValueOf(m).MethodByName("AfterSave").Call([]reflect.Value{})
	}

	return nil
}

func Update(c appengine.Context, m interface{}) error {
	return nil
}
