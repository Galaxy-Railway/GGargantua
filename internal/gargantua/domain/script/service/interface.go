package service

import "github.com/Galaxy-Railway/GGargantua/internal/gargantua/domain/script/vm"

type ScriptService interface {
	NewScriptVM() (*vm.VM, error)
	RunScript(*vm.VM, string) (string, error)
	ClearScriptVM(*vm.VM) error
}
