package cache

import (
	"errors"
	"strconv"
)

type Cache struct {
	CMap map[string]*CacheValue
}

func NewCache() Cache {
	return Cache{
		CMap: make(map[string]*CacheValue),
	}
}

func (c *Cache) Set(keyOfCache, typeOfCache string, value string) error {
	t, err := StringToCacheType(typeOfCache)
	if err != nil {
		return err
	}
	cv := CacheValue{
		TypeOfCache: t,
	}
	switch t {
	case IntCacheType:
		var tmp int
		tmp, err = strconv.Atoi(value)
		if err != nil {
			return err
		}
		cv.Value = tmp
	case Int64CacheType:
		var tmp int64
		tmp, err = strconv.ParseInt(value, 0, 64)
		if err != nil {
			return err
		}
		cv.Value = tmp
	case FloatCacheType:
		var tmp float64
		tmp, err = strconv.ParseFloat(value, 32)
		if err != nil {
			return err
		}
		cv.Value = float32(tmp)
	case Float64CacheType:
		var tmp float64
		tmp, err = strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		cv.Value = tmp
	case BoolCacheType:
		var tmp bool
		tmp, err = strconv.ParseBool(value)
		if err != nil {
			return err
		}
		cv.Value = tmp
	}
	c.CMap[keyOfCache] = &cv
	return nil
}

func (c *Cache) Get(keyOfCache string) interface{} {
	if result, ok := c.CMap[keyOfCache]; ok {
		return result.Value
	}
	return nil
}

type CacheValue struct {
	TypeOfCache CacheType
	Value       interface{}
}

type CacheType int

const (
	BoolCacheType CacheType = iota
	StringCacheType
	IntCacheType
	Int64CacheType
	FloatCacheType
	Float64CacheType

	UnknownCacheType CacheType = -1
)

var (
	UnknownCacheTypeError = errors.New("type of this cache is not supported, yet. Please choose one type of (bool, string, int, int64, float, float64)")
)

func StringToCacheType(str string) (CacheType, error) {
	switch str {
	case "bool":
		return BoolCacheType, nil
	case "string":
		return StringCacheType, nil
	case "int":
		return IntCacheType, nil
	case "int64":
		return Int64CacheType, nil
	case "float":
		return FloatCacheType, nil
	case "float64":
		return Float64CacheType, nil
	}
	return UnknownCacheType, UnknownCacheTypeError
}
