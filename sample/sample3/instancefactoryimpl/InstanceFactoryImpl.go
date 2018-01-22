package instancefactoryimpl

// This file is generated by update.sh.

import (
	impl "github.com/maxvalor/instancefactory/sample/sample3/printera/impl"
	impl1 "github.com/maxvalor/instancefactory/sample/sample3/printerb/impl"
)

type InstanceFactoryImpl struct {
	implMap map[string]func() interface{}
}

func NewInstanceFactoryImpl() (instance *InstanceFactoryImpl) {
	instance = &InstanceFactoryImpl{}
	instance.implMap = make(map[string]func() interface{})
	instance.implMap["github.com/maxvalor/instancefactory/sample/sample3/printera.PrinterA"] = getimplPrinterAImpl
	instance.implMap["github.com/maxvalor/instancefactory/sample/sample3/printerb.PrinterB"] = getimpl1PrinterBImpl
	return
}

func (fac *InstanceFactoryImpl)GetStructImpl(interfaceName string) (instance interface{}) {
	f := fac.implMap[interfaceName]
	if f != nil {
		instance = f()
	}
	return
}

func getimplPrinterAImpl() interface{} {
	return &impl.PrinterAImpl{}
}

func getimpl1PrinterBImpl() interface{} {
	return &impl1.PrinterBImpl{}
}

