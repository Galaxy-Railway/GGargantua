package vm

import (
	"github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/script/cache"
	"github.com/Galaxy-Railway/GGargantua/pkg/common"
	"github.com/robertkrimen/otto"
)

type VM struct {
	Ot    *otto.Otto
	Cache cache.Cache
	JobId string
}

func NewVM() *VM {
	vm := &VM{
		Ot:    otto.New(),
		Cache: cache.NewCache(),
		JobId: "",
	}
	err := vm.SetCacheSetter()
	if err != nil {
		common.Logger().Errorf("failed to set cacheSet() into vm, err: %v", err)
	}
	err = vm.SetCacheGetter()
	if err != nil {
		common.Logger().Errorf("failed to set cacheGet() into vm, err: %v", err)
	}
	return vm
}

func (v *VM) CleanVM() {
	v.Ot = otto.New()
}

func (v *VM) SetResponse(response interface{}) error {
	return v.Ot.Set("response", response)
}

func (v *VM) RunScript(script string) (otto.Value, error) {
	return v.Ot.Run(script)
}

// SetCacheSetter will set a setCache() function in js vm.
// setCache() function looks like below to user:
// function setCache(key, type, value) {
// 	 return err
// }
// key have to be a string
// type have to be one of (int, int64, float, float64, string, bool)
// type of value should be equal to "type"
// err should be empty
func (v *VM) SetCacheSetter() error {
	return v.Ot.Set("setCache", func(call otto.FunctionCall) otto.Value {
		keyOfCache, err := call.Argument(0).ToString()
		if err != nil {
			result, _ := v.Ot.ToValue(err)
			return result
		}
		typeOfCache, err := call.Argument(1).ToString()
		if err != nil {
			result, _ := v.Ot.ToValue(err)
			return result
		}
		valueOfCache, err := call.Argument(2).ToString()
		if err != nil {
			result, _ := v.Ot.ToValue(err)
			return result
		}
		err = v.SetCache(keyOfCache, typeOfCache, valueOfCache)
		if err != nil {
			result, _ := v.Ot.ToValue(err)
			return result
		}
		return otto.NullValue()
	})
}

func (v *VM) SetCache(keyOfCache, typeOfCache, valueOfCache string) error {
	return v.Cache.Set(keyOfCache, typeOfCache, valueOfCache)
}

// SetCacheGetter will set a getCache() function in js vm.
// getCache() function looks like below to user:
// function getCache(key) {
// 	 return value
// }
// key have to be a string
// if key is not in cache, value will be null
func (v *VM) SetCacheGetter() error {
	return v.Ot.Set("getCache", func(call otto.FunctionCall) otto.Value {
		keyOfCache, err := call.Argument(0).ToString()
		if err != nil {
			result, _ := v.Ot.ToValue(err)
			return result
		}
		value := v.GetCache(keyOfCache)
		if value == nil {
			return otto.NullValue()
		} else {
			result, _ := v.Ot.ToValue(value)
			return result
		}
	})
}

func (v *VM) GetCache(keyOfCache string) interface{} {
	return v.Cache.Get(keyOfCache)
}
