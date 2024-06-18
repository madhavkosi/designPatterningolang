package strategypattern

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache() *Cache {
	storage := make(map[string]string)
	lfu := &Lfu{}
	return &Cache{
		storage:      storage,
		evictionAlgo: lfu,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) Get(key string) {
	delete(c.storage, key)
	c.capacity--
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}
