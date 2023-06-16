package interpret

import (
	"github.com/dop251/goja"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/interpret", New())
}

type (
	// RootModule is the global module instance that will create module
	// instances for each VU.
	RootModule struct{}

	// ModuleInstance represents an instance of the JS module.
	ModuleInstance struct {
		// vu provides methods for accessing internal k6 objects for a VU
		vu modules.VU
	}
)

var (
	_ modules.Module   = &RootModule{}
	_ modules.Instance = &ModuleInstance{}
)

func New() *RootModule {
	return &RootModule{}
}

func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &ModuleInstance{
		vu: vu,
	}
}

type Interpret struct {
	vu modules.VU
}

func (r *Interpret) Run(src string, args ...interface{}) interface{} {
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)

	_, err := i.Eval(src)
	if err != nil {
		common.Throw(r.vu.Runtime(), err)
	}

	v, err := i.Eval("interpret.Run")
	if err != nil {
		common.Throw(r.vu.Runtime(), err)
	}

	run := v.Interface().(func(...interface{}) interface{})

	return run(args...)
}

func (mi *ModuleInstance) Exports() modules.Exports {
	return modules.Exports{
		Named: map[string]interface{}{
			"Interpret": mi.newInterpret,
		},
	}
}

func (mi *ModuleInstance) newInterpret(c goja.ConstructorCall) *goja.Object {
	rt := mi.vu.Runtime()
	obj := &Interpret{
		vu: mi.vu,
	}

	return rt.ToValue(obj).ToObject(rt)
}
