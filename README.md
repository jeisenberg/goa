Go App engine plus
==================

## Description:

Go app engine plus is a package made to help developers build faster and more efficient web applications on Google App Engine using Go. It has 4 main purposes:

1. Add 'automatic' support for a caching layer.
2. Keep the Id of an entity on the struct itself, as opposed to being returned separate from that entity on queries.
3. Perform repeated CRUD functions in fewer lines of code.
4. Allow support for callbacks on read/write of entites to datastore

### Built-in memcache

GOA will automatically write or read an entity to memcache, which can drastically speed up performance on datastore Get calls.

For example:

```
import (
	"goa"
)

type User struct {
	Name  string
	Id    string
}

func userHandler(c appengine.Context, u User){
	err := goa.Save(c, &u)
}

```
Here the save method will first store the entity in the datastore, and will then encode the struct (using the Gob encoder) into memcache. The memcache key will be the same is both the Id field in the struct, or the string that comes from calling (*key datastore.Key).Encode()

### Entities and Keys

One quirk with the datastore on app engine is that entity elements and keys are returned separately. For example, a typical datastore query in the app engine SDK looks something like this:

```
q := datastore.NewQuery("Photo").Ancestor(tomKey)
var photos []Photo
keys, err := q.GetAll(c, &photos)
```
Which returns both an array of photos and a separate array of keys. This might work for certain elements of data processing, but when rendering JSON or server side templates to the client, it is helpful to have the actual Id on the actual entity to do things like build dynamic links. 

Every struct that is persisted to the datastore must have a field called Id, that is a string. This will be the encoded datastore key, which is used to read/write/update/delete the entity from the datastore. An example is:

```
type User struct {
	Name  string
	Id    string
}
```
The Id does not need to be set on a newly saved entity, but will be created automatically.

### CRUD Methods

Right now, supported methods for CRUD actions are
```
goa.Save(c appengine.Context, i interface{})

goa.Update(c appengine.Context, i interface{})

goa.Get(c appengine.Context, id string, i interface{})

goa.Delete(c appengine.Context, id string)
```
Please reference the source test server at /goaeplus_test to see how these can be implemented.

### Callbacks

Callbacks are a common feature amongst web frameworks like Rails, that can kick off a function before or after an entity has been created, updated, or deleted from the datastore. In order to initiate a callback, simply add the function to a pointer instance of the struct entity that is persisted. For example:

```
type User struct {
	Id string
	FirstName string
	LastName string
	FullName string

}

func (u *User) BeforeSave() {
	u.FullName := strings.Join([]string{u.FirstName, u.LastName}, " ")
}

```

Here, when I call Save() on a user, it will first call the BeforeSave callback, and add the first name and last name to create the full name.

Currently supported callbacks are 
* BeforeSave()
* AfterSave()
* BeforeUpdate()
* AfterUpdate()

Please note that at the moment callbacks are not allowed to take arguments. This is something we are working on implementing.

