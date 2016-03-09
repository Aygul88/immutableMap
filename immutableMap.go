package immutableMap

import (
	"errors"
	"sync"
)

// Error code returned when there was something wrong
var (
	ErrNullKey         = errors.New("error: key is null in 'Set' method.")
	ErrExistingKey     = errors.New("error: key already exists in this map.")
	ErrNullValue       = errors.New("error: value is null in 'Set' method.")
	ErrNullValueForKey = errors.New("error: key is not set in this immutable map.")
)

// To ensure the map is thread-safe
var rwLock = sync.RWMutex{}

type ImmutableMap struct {
	data map[interface{}]interface{}
}

// ImmutableMap Constrcutor
func NewImmutableMap() *ImmutableMap {
	return &ImmutableMap{
		data: make(map[interface{}]interface{}),
	}
}

/*
 * Insert new key / value pair into immutable map
 * @param key (interface{})
 * @param value (interface{})
 * @return error (error)
 */
func (self *ImmutableMap) Set(key interface{}, value interface{}) error {
	if key == nil {
		return ErrNullKey
	}
	if value == nil {
		return ErrNullValue
	}
	rwLock.Lock()
	defer rwLock.Unlock()
	if _, ok := self.data[key]; ok {
		return ErrExistingKey
	}
	self.data[key] = value
	return nil
}

/*
 * Read value for certain key from immutable map
 * @param key (interface{})
 * @return interface{}, error
 */
func (self *ImmutableMap) Get(key interface{}) (interface{}, error) {
	if key == nil {
		return nil, ErrNullKey
	}
	rwLock.Lock()
	defer rwLock.Unlock()
	if value, ok := self.data[key]; !ok {
		return nil, ErrNullValueForKey
	} else {
		return value, nil
	}
}

/*
 * Delete key / value pair from immutable map
 * @param key (interface{})
 * @return error (error)
 */
func (self *ImmutableMap) Delete(key interface{}) error {
	if key == nil {
		return ErrNullKey
	}
	rwLock.Lock()
	defer rwLock.Unlock()
	if _, ok := self.data[key]; !ok {
		return ErrNullValueForKey
	} else {
		delete(self.data, key)
		return nil
	}
}

/*
 * Flush data in immutable map away
 * @return error (error)
 */
func (self *ImmutableMap) Flush() error {
	rwLock.Lock()
	defer rwLock.Unlock()
	if len(self.data) == 0 {
		return nil
	}
	for key, _ := range self.data {
		delete(self.data, key)
	}
	return nil
}
