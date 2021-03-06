package main

//go:generate ${GOPATH}/src/github.com/maxvalor/instancefactory/update.sh ${GOPATH}/src/github.com/maxvalor/instancefactory/sample/sample1/setting

import (
	"github.com/maxvalor/instancefactory"
	"github.com/maxvalor/instancefactory/sample/sample1/instancefactoryimpl"
	"github.com/maxvalor/instancefactory/sample/sample1/printer"
)

func main()  {
	instancefactory.SetFactory(instancefactoryimpl.NewInstanceFactoryImpl())
	pt := instancefactory.GetFactory().GetStructImpl(printer.Class).(printer.Printer)
	pt.Print("Hello instance factory")
}
