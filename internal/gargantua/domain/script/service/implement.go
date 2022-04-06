package service

import "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/script/vm"

type ScriptServiceImpl struct {
}

func (s ScriptServiceImpl) NewScriptVM() (*vm.VM, error) {
	return vm.NewVM()
}

func (s ScriptServiceImpl) RunScript(v *vm.VM, str string) (string, error) {
	return v.RunScript(str)
}

func (s ScriptServiceImpl) ClearScriptVM(v *vm.VM) error {
	v.Ot = nil
	v.Cache.CMap = nil
	return nil
}

func NewScriptService() ScriptService {
	return &ScriptServiceImpl{}
}
