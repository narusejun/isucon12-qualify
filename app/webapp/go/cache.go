package isuports

import (
	"database/sql"
	"sync"
)

const SHARD_COUNT = 16

type stringCache[T any] struct {
	mux0         *sync.RWMutex
	mux1         *sync.RWMutex
	mux2         *sync.RWMutex
	mux3         *sync.RWMutex
	mux4         *sync.RWMutex
	mux5         *sync.RWMutex
	mux6         *sync.RWMutex
	mux7         *sync.RWMutex
	mux8         *sync.RWMutex
	mux9         *sync.RWMutex
	mux10        *sync.RWMutex
	mux11        *sync.RWMutex
	mux12        *sync.RWMutex
	mux13        *sync.RWMutex
	mux14        *sync.RWMutex
	mux15        *sync.RWMutex
	data0        map[string]*T
	data1        map[string]*T
	data2        map[string]*T
	data3        map[string]*T
	data4        map[string]*T
	data5        map[string]*T
	data6        map[string]*T
	data7        map[string]*T
	data8        map[string]*T
	data9        map[string]*T
	data10       map[string]*T
	data11       map[string]*T
	data12       map[string]*T
	data13       map[string]*T
	data14       map[string]*T
	data15       map[string]*T
	notfoundFunc func(key string) (*T, error)
}

func newStringCache[T any](size int, notfoundFunc func(key string) (*T, error)) *stringCache[T] {
	return &stringCache[T]{
		mux0:         &sync.RWMutex{},
		mux1:         &sync.RWMutex{},
		mux2:         &sync.RWMutex{},
		mux3:         &sync.RWMutex{},
		mux4:         &sync.RWMutex{},
		mux5:         &sync.RWMutex{},
		mux6:         &sync.RWMutex{},
		mux7:         &sync.RWMutex{},
		mux8:         &sync.RWMutex{},
		mux9:         &sync.RWMutex{},
		mux10:        &sync.RWMutex{},
		mux11:        &sync.RWMutex{},
		mux12:        &sync.RWMutex{},
		mux13:        &sync.RWMutex{},
		mux14:        &sync.RWMutex{},
		mux15:        &sync.RWMutex{},
		data0:        make(map[string]*T, size),
		data1:        make(map[string]*T, size),
		data2:        make(map[string]*T, size),
		data3:        make(map[string]*T, size),
		data4:        make(map[string]*T, size),
		data5:        make(map[string]*T, size),
		data6:        make(map[string]*T, size),
		data7:        make(map[string]*T, size),
		data8:        make(map[string]*T, size),
		data9:        make(map[string]*T, size),
		data10:       make(map[string]*T, size),
		data11:       make(map[string]*T, size),
		data12:       make(map[string]*T, size),
		data13:       make(map[string]*T, size),
		data14:       make(map[string]*T, size),
		data15:       make(map[string]*T, size),
		notfoundFunc: notfoundFunc,
	}
}

func (c *stringCache[T]) reset() {
	c.mux0.Lock()
	c.mux1.Lock()
	c.mux2.Lock()
	c.mux3.Lock()
	c.mux4.Lock()
	c.mux5.Lock()
	c.mux6.Lock()
	c.mux7.Lock()
	c.mux8.Lock()
	c.mux9.Lock()
	c.mux10.Lock()
	c.mux11.Lock()
	c.mux12.Lock()
	c.mux13.Lock()
	c.mux14.Lock()
	c.mux15.Lock()
	c.data0 = make(map[string]*T, len(c.data0))
	c.data1 = make(map[string]*T, len(c.data1))
	c.data2 = make(map[string]*T, len(c.data2))
	c.data3 = make(map[string]*T, len(c.data3))
	c.data4 = make(map[string]*T, len(c.data4))
	c.data5 = make(map[string]*T, len(c.data5))
	c.data6 = make(map[string]*T, len(c.data6))
	c.data7 = make(map[string]*T, len(c.data7))
	c.data8 = make(map[string]*T, len(c.data8))
	c.data9 = make(map[string]*T, len(c.data9))
	c.data10 = make(map[string]*T, len(c.data10))
	c.data11 = make(map[string]*T, len(c.data11))
	c.data12 = make(map[string]*T, len(c.data12))
	c.data13 = make(map[string]*T, len(c.data13))
	c.data14 = make(map[string]*T, len(c.data14))
	c.data15 = make(map[string]*T, len(c.data15))
	c.mux0.Unlock()
	c.mux1.Unlock()
	c.mux2.Unlock()
	c.mux3.Unlock()
	c.mux4.Unlock()
	c.mux5.Unlock()
	c.mux6.Unlock()
	c.mux7.Unlock()
	c.mux8.Unlock()
	c.mux9.Unlock()
	c.mux10.Unlock()
	c.mux11.Unlock()
	c.mux12.Unlock()
	c.mux13.Unlock()
	c.mux14.Unlock()
	c.mux15.Unlock()
}

func (c *stringCache[T]) get(key string) (*T, bool) {
	shard := fnv32(key) % SHARD_COUNT
	switch shard {
	case 0:
		c.mux0.RLock()
		if v, ok := c.data0[key]; ok {
			c.mux0.RUnlock()
			return v, ok
		}
		c.mux0.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux0.Lock()
		c.data0[key] = data
		c.mux0.Unlock()
		return data, true

	case 1:
		c.mux1.RLock()
		if v, ok := c.data1[key]; ok {
			c.mux1.RUnlock()
			return v, ok
		}
		c.mux1.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux1.Lock()
		c.data1[key] = data
		c.mux1.Unlock()
		return data, true

	case 2:
		c.mux2.RLock()
		if v, ok := c.data2[key]; ok {
			c.mux2.RUnlock()
			return v, ok
		}
		c.mux2.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux2.Lock()
		c.data2[key] = data
		c.mux2.Unlock()
		return data, true

	case 3:
		c.mux3.RLock()
		if v, ok := c.data3[key]; ok {
			c.mux3.RUnlock()
			return v, ok
		}
		c.mux3.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux3.Lock()
		c.data3[key] = data
		c.mux3.Unlock()
		return data, true

	case 4:
		c.mux4.RLock()
		if v, ok := c.data4[key]; ok {
			c.mux4.RUnlock()
			return v, ok
		}
		c.mux4.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux4.Lock()
		c.data4[key] = data
		c.mux4.Unlock()
		return data, true

	case 5:
		c.mux5.RLock()
		if v, ok := c.data5[key]; ok {
			c.mux5.RUnlock()
			return v, ok
		}
		c.mux5.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux5.Lock()
		c.data5[key] = data
		c.mux5.Unlock()
		return data, true

	case 6:
		c.mux6.RLock()
		if v, ok := c.data6[key]; ok {
			c.mux6.RUnlock()
			return v, ok
		}
		c.mux6.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux6.Lock()
		c.data6[key] = data
		c.mux6.Unlock()
		return data, true

	case 7:
		c.mux7.RLock()
		if v, ok := c.data7[key]; ok {
			c.mux7.RUnlock()
			return v, ok
		}
		c.mux7.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux7.Lock()
		c.data7[key] = data
		c.mux7.Unlock()
		return data, true

	case 8:
		c.mux8.RLock()
		if v, ok := c.data8[key]; ok {
			c.mux8.RUnlock()
			return v, ok
		}
		c.mux8.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux8.Lock()
		c.data8[key] = data
		c.mux8.Unlock()
		return data, true

	case 9:
		c.mux9.RLock()
		if v, ok := c.data9[key]; ok {
			c.mux9.RUnlock()
			return v, ok
		}
		c.mux9.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux9.Lock()
		c.data9[key] = data
		c.mux9.Unlock()
		return data, true

	case 10:
		c.mux10.RLock()
		if v, ok := c.data10[key]; ok {
			c.mux10.RUnlock()
			return v, ok
		}
		c.mux10.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux10.Lock()
		c.data10[key] = data
		c.mux10.Unlock()
		return data, true

	case 11:
		c.mux11.RLock()
		if v, ok := c.data11[key]; ok {
			c.mux11.RUnlock()
			return v, ok
		}
		c.mux11.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux11.Lock()
		c.data11[key] = data
		c.mux11.Unlock()
		return data, true

	case 12:
		c.mux12.RLock()
		if v, ok := c.data12[key]; ok {
			c.mux12.RUnlock()
			return v, ok
		}
		c.mux12.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux12.Lock()
		c.data12[key] = data
		c.mux12.Unlock()
		return data, true

	case 13:
		c.mux13.RLock()
		if v, ok := c.data13[key]; ok {
			c.mux13.RUnlock()
			return v, ok
		}
		c.mux13.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux13.Lock()
		c.data13[key] = data
		c.mux13.Unlock()
		return data, true

	case 14:
		c.mux14.RLock()
		if v, ok := c.data14[key]; ok {
			c.mux14.RUnlock()
			return v, ok
		}
		c.mux14.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux14.Lock()
		c.data14[key] = data
		c.mux14.Unlock()
		return data, true

	case 15:
		c.mux15.RLock()
		if v, ok := c.data15[key]; ok {
			c.mux15.RUnlock()
			return v, ok
		}
		c.mux15.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux15.Lock()
		c.data15[key] = data
		c.mux15.Unlock()
		return data, true
	}
	return nil, false
}

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 uint32 = 16777619
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

func (c *stringCache[T]) set(key string, data *T) {
	shard := fnv32(key) % SHARD_COUNT
	switch shard {
	case 0:
		c.mux0.Lock()
		c.data0[key] = data
		c.mux0.Unlock()
	case 1:
		c.mux1.Lock()
		c.data1[key] = data
		c.mux1.Unlock()
	case 2:
		c.mux2.Lock()
		c.data2[key] = data
		c.mux2.Unlock()
	case 3:
		c.mux3.Lock()
		c.data3[key] = data
		c.mux3.Unlock()
	case 4:
		c.mux4.Lock()
		c.data4[key] = data
		c.mux4.Unlock()
	case 5:
		c.mux5.Lock()
		c.data5[key] = data
		c.mux5.Unlock()
	case 6:
		c.mux6.Lock()
		c.data6[key] = data
		c.mux6.Unlock()
	case 7:
		c.mux7.Lock()
		c.data7[key] = data
		c.mux7.Unlock()
	case 8:
		c.mux8.Lock()
		c.data8[key] = data
		c.mux8.Unlock()
	case 9:
		c.mux9.Lock()
		c.data9[key] = data
		c.mux9.Unlock()
	case 10:
		c.mux10.Lock()
		c.data10[key] = data
		c.mux10.Unlock()
	case 11:
		c.mux11.Lock()
		c.data11[key] = data
		c.mux11.Unlock()
	case 12:
		c.mux12.Lock()
		c.data12[key] = data
		c.mux12.Unlock()
	case 13:
		c.mux13.Lock()
		c.data13[key] = data
		c.mux13.Unlock()
	case 14:
		c.mux14.Lock()
		c.data14[key] = data
		c.mux14.Unlock()
	case 15:
		c.mux15.Lock()
		c.data15[key] = data
		c.mux15.Unlock()
	}
}

type intCache[T any] struct {
	mux0         *sync.RWMutex
	mux1         *sync.RWMutex
	mux2         *sync.RWMutex
	mux3         *sync.RWMutex
	mux4         *sync.RWMutex
	mux5         *sync.RWMutex
	mux6         *sync.RWMutex
	mux7         *sync.RWMutex
	mux8         *sync.RWMutex
	mux9         *sync.RWMutex
	mux10        *sync.RWMutex
	mux11        *sync.RWMutex
	mux12        *sync.RWMutex
	mux13        *sync.RWMutex
	mux14        *sync.RWMutex
	mux15        *sync.RWMutex
	data0        map[int]*T
	data1        map[int]*T
	data2        map[int]*T
	data3        map[int]*T
	data4        map[int]*T
	data5        map[int]*T
	data6        map[int]*T
	data7        map[int]*T
	data8        map[int]*T
	data9        map[int]*T
	data10       map[int]*T
	data11       map[int]*T
	data12       map[int]*T
	data13       map[int]*T
	data14       map[int]*T
	data15       map[int]*T
	notfoundFunc func(key int) (*T, error)
}

func newIntCache[T any](size int, notfoundFunc func(key int) (*T, error)) *intCache[T] {
	return &intCache[T]{
		mux0:         &sync.RWMutex{},
		mux1:         &sync.RWMutex{},
		mux2:         &sync.RWMutex{},
		mux3:         &sync.RWMutex{},
		mux4:         &sync.RWMutex{},
		mux5:         &sync.RWMutex{},
		mux6:         &sync.RWMutex{},
		mux7:         &sync.RWMutex{},
		mux8:         &sync.RWMutex{},
		mux9:         &sync.RWMutex{},
		mux10:        &sync.RWMutex{},
		mux11:        &sync.RWMutex{},
		mux12:        &sync.RWMutex{},
		mux13:        &sync.RWMutex{},
		mux14:        &sync.RWMutex{},
		mux15:        &sync.RWMutex{},
		data0:        make(map[int]*T, size),
		data1:        make(map[int]*T, size),
		data2:        make(map[int]*T, size),
		data3:        make(map[int]*T, size),
		data4:        make(map[int]*T, size),
		data5:        make(map[int]*T, size),
		data6:        make(map[int]*T, size),
		data7:        make(map[int]*T, size),
		data8:        make(map[int]*T, size),
		data9:        make(map[int]*T, size),
		data10:       make(map[int]*T, size),
		data11:       make(map[int]*T, size),
		data12:       make(map[int]*T, size),
		data13:       make(map[int]*T, size),
		data14:       make(map[int]*T, size),
		data15:       make(map[int]*T, size),
		notfoundFunc: notfoundFunc,
	}
}

func (c *intCache[T]) reset() {
	c.mux0.Lock()
	c.mux1.Lock()
	c.mux2.Lock()
	c.mux3.Lock()
	c.mux4.Lock()
	c.mux5.Lock()
	c.mux6.Lock()
	c.mux7.Lock()
	c.mux8.Lock()
	c.mux9.Lock()
	c.mux10.Lock()
	c.mux11.Lock()
	c.mux12.Lock()
	c.mux13.Lock()
	c.mux14.Lock()
	c.mux15.Lock()
	c.data0 = make(map[int]*T, len(c.data0))
	c.data1 = make(map[int]*T, len(c.data1))
	c.data2 = make(map[int]*T, len(c.data2))
	c.data3 = make(map[int]*T, len(c.data3))
	c.data4 = make(map[int]*T, len(c.data4))
	c.data5 = make(map[int]*T, len(c.data5))
	c.data6 = make(map[int]*T, len(c.data6))
	c.data7 = make(map[int]*T, len(c.data7))
	c.data8 = make(map[int]*T, len(c.data8))
	c.data9 = make(map[int]*T, len(c.data9))
	c.data10 = make(map[int]*T, len(c.data10))
	c.data11 = make(map[int]*T, len(c.data11))
	c.data12 = make(map[int]*T, len(c.data12))
	c.data13 = make(map[int]*T, len(c.data13))
	c.data14 = make(map[int]*T, len(c.data14))
	c.data15 = make(map[int]*T, len(c.data15))
	c.mux0.Unlock()
	c.mux1.Unlock()
	c.mux2.Unlock()
	c.mux3.Unlock()
	c.mux4.Unlock()
	c.mux5.Unlock()
	c.mux6.Unlock()
	c.mux7.Unlock()
	c.mux8.Unlock()
	c.mux9.Unlock()
	c.mux10.Unlock()
	c.mux11.Unlock()
	c.mux12.Unlock()
	c.mux13.Unlock()
	c.mux14.Unlock()
	c.mux15.Unlock()
}

func (c *intCache[T]) get(key int) (*T, bool) {
	shard := key % SHARD_COUNT
	switch shard {
	case 0:
		c.mux0.RLock()
		if v, ok := c.data0[key]; ok {
			c.mux0.RUnlock()
			return v, ok
		}
		c.mux0.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux0.Lock()
		c.data0[key] = data
		c.mux0.Unlock()
		return data, true

	case 1:
		c.mux1.RLock()
		if v, ok := c.data1[key]; ok {
			c.mux1.RUnlock()
			return v, ok
		}
		c.mux1.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux1.Lock()
		c.data1[key] = data
		c.mux1.Unlock()
		return data, true

	case 2:
		c.mux2.RLock()
		if v, ok := c.data2[key]; ok {
			c.mux2.RUnlock()
			return v, ok
		}
		c.mux2.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux2.Lock()
		c.data2[key] = data
		c.mux2.Unlock()
		return data, true

	case 3:
		c.mux3.RLock()
		if v, ok := c.data3[key]; ok {
			c.mux3.RUnlock()
			return v, ok
		}
		c.mux3.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux3.Lock()
		c.data3[key] = data
		c.mux3.Unlock()
		return data, true

	case 4:
		c.mux4.RLock()
		if v, ok := c.data4[key]; ok {
			c.mux4.RUnlock()
			return v, ok
		}
		c.mux4.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux4.Lock()
		c.data4[key] = data
		c.mux4.Unlock()
		return data, true

	case 5:
		c.mux5.RLock()
		if v, ok := c.data5[key]; ok {
			c.mux5.RUnlock()
			return v, ok
		}
		c.mux5.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux5.Lock()
		c.data5[key] = data
		c.mux5.Unlock()
		return data, true

	case 6:
		c.mux6.RLock()
		if v, ok := c.data6[key]; ok {
			c.mux6.RUnlock()
			return v, ok
		}
		c.mux6.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux6.Lock()
		c.data6[key] = data
		c.mux6.Unlock()
		return data, true

	case 7:
		c.mux7.RLock()
		if v, ok := c.data7[key]; ok {
			c.mux7.RUnlock()
			return v, ok
		}
		c.mux7.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux7.Lock()
		c.data7[key] = data
		c.mux7.Unlock()
		return data, true

	case 8:
		c.mux8.RLock()
		if v, ok := c.data8[key]; ok {
			c.mux8.RUnlock()
			return v, ok
		}
		c.mux8.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux8.Lock()
		c.data8[key] = data
		c.mux8.Unlock()
		return data, true

	case 9:
		c.mux9.RLock()
		if v, ok := c.data9[key]; ok {
			c.mux9.RUnlock()
			return v, ok
		}
		c.mux9.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux9.Lock()
		c.data9[key] = data
		c.mux9.Unlock()
		return data, true

	case 10:
		c.mux10.RLock()
		if v, ok := c.data10[key]; ok {
			c.mux10.RUnlock()
			return v, ok
		}
		c.mux10.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux10.Lock()
		c.data10[key] = data
		c.mux10.Unlock()
		return data, true

	case 11:
		c.mux11.RLock()
		if v, ok := c.data11[key]; ok {
			c.mux11.RUnlock()
			return v, ok
		}
		c.mux11.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux11.Lock()
		c.data11[key] = data
		c.mux11.Unlock()
		return data, true

	case 12:
		c.mux12.RLock()
		if v, ok := c.data12[key]; ok {
			c.mux12.RUnlock()
			return v, ok
		}
		c.mux12.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux12.Lock()
		c.data12[key] = data
		c.mux12.Unlock()
		return data, true

	case 13:
		c.mux13.RLock()
		if v, ok := c.data13[key]; ok {
			c.mux13.RUnlock()
			return v, ok
		}
		c.mux13.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux13.Lock()
		c.data13[key] = data
		c.mux13.Unlock()
		return data, true

	case 14:
		c.mux14.RLock()
		if v, ok := c.data14[key]; ok {
			c.mux14.RUnlock()
			return v, ok
		}
		c.mux14.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux14.Lock()
		c.data14[key] = data
		c.mux14.Unlock()
		return data, true

	case 15:
		c.mux15.RLock()
		if v, ok := c.data15[key]; ok {
			c.mux15.RUnlock()
			return v, ok
		}
		c.mux15.RUnlock()

		data, err := c.notfoundFunc(key)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, false
			}
			return nil, false
		}
		c.mux15.Lock()
		c.data15[key] = data
		c.mux15.Unlock()
		return data, true
	}
	return nil, false
}

func (c *intCache[T]) set(key int, data *T) {
	shard := key % SHARD_COUNT
	switch shard {
	case 0:
		c.mux0.Lock()
		c.data0[key] = data
		c.mux0.Unlock()
	case 1:
		c.mux1.Lock()
		c.data1[key] = data
		c.mux1.Unlock()
	case 2:
		c.mux2.Lock()
		c.data2[key] = data
		c.mux2.Unlock()
	case 3:
		c.mux3.Lock()
		c.data3[key] = data
		c.mux3.Unlock()
	case 4:
		c.mux4.Lock()
		c.data4[key] = data
		c.mux4.Unlock()
	case 5:
		c.mux5.Lock()
		c.data5[key] = data
		c.mux5.Unlock()
	case 6:
		c.mux6.Lock()
		c.data6[key] = data
		c.mux6.Unlock()
	case 7:
		c.mux7.Lock()
		c.data7[key] = data
		c.mux7.Unlock()
	case 8:
		c.mux8.Lock()
		c.data8[key] = data
		c.mux8.Unlock()
	case 9:
		c.mux9.Lock()
		c.data9[key] = data
		c.mux9.Unlock()
	case 10:
		c.mux10.Lock()
		c.data10[key] = data
		c.mux10.Unlock()
	case 11:
		c.mux11.Lock()
		c.data11[key] = data
		c.mux11.Unlock()
	case 12:
		c.mux12.Lock()
		c.data12[key] = data
		c.mux12.Unlock()
	case 13:
		c.mux13.Lock()
		c.data13[key] = data
		c.mux13.Unlock()
	case 14:
		c.mux14.Lock()
		c.data14[key] = data
		c.mux14.Unlock()
	case 15:
		c.mux15.Lock()
		c.data15[key] = data
		c.mux15.Unlock()
	}
}
