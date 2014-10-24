package goaeplus

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

type BeforeUpdateInterface interface {
	BeforeUpdate()
}

type AfterUpdateInterface interface {
	AfterUpdate()
}