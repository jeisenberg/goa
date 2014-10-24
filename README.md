Go App engine plus
==================

## Description:

Go app engine plus is a package made to help developers build web applications on Google App Engine using go. It has 4 main purposes:

1. Add 'automatic' support for a caching layer.
2. Keep the Id of an entity on the entity itself, as opposed to separate from that entity, which is how the datastore defaults 
3. Allow developers to perform basic Read/Write/Delete functions to the datastore/memcache services in fewer lines of code
4. Allow support for callbacks on read/write of entites to datastore

### Built-in memcache

GOA will automatically write or read an entity to memcache, which will drastically speed up performance on datastore Get and Put calls.

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

### Entity's and Keys

### CRUD Methods

### Callbacks


