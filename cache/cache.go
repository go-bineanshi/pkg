package cache

type ICache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}) error
	Delete(key string)
}

type Cache struct {
	store ICache
}

func NewCache(store ICache) *Cache {
	return &Cache{
		store: store,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	return c.store.Get(key)
}
func (c *Cache) Set(key string, value interface{}) error {
	return c.store.Set(key, value)
}

func (c *Cache) Delete(key string) {
	c.store.Delete(key)
}
