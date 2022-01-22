package main

import (
	"fmt"
	"github.com/voicurobert/design_patterns_in_go/app/src/section_01_solid"
	builder2 "github.com/voicurobert/design_patterns_in_go/app/src/section_02_builder"
	factory2 "github.com/voicurobert/design_patterns_in_go/app/src/section_03_factories"
	prototype2 "github.com/voicurobert/design_patterns_in_go/app/src/section_04_prototype"
	singleton "github.com/voicurobert/design_patterns_in_go/app/src/section_05_singleton"
	adapter "github.com/voicurobert/design_patterns_in_go/app/src/section_07_adapter"
	bridge "github.com/voicurobert/design_patterns_in_go/app/src/section_08_bridge"
	composite "github.com/voicurobert/design_patterns_in_go/app/src/section_09_composite"
	"strings"
)

func main() {

	//runSolid()

	//runBuilder()

	//runFactory()

	//runPrototype()

	//runSingleton()

	//runAdapter()

	//runBridge()

	runComposite()

}

func runSolid() {
	solid.SingleResponsibilityPrinciple()
	solid.OpenClosedPrinciple()
	solid.LiskovPrinciple()
	solid.DependencyInversionPrinciple()
}

func runBuilder() {
	printPatternDescription("A builder is a separate component used for building an object. \nTo make builder fluent," +
		"return the receiver - allows chaining. Different facets of an object can be built with different builders working in tandem via a common struct")
	builder2.StringBuilderExample()
	builder2.Example()
	builder2.FacetsExample()
	builder2.ParameterExample()
	builder2.FunctionalExample()
}

func runFactory() {
	printPatternDescription("A factory function (constructor) is a helper function for making struct instances. \n" +
		"A factory is any entity than can take care of object creation. Can be a function or a dedicated struct.")
	factory2.FunctionExample()
	factory2.InterfaceExample()
	factory2.GeneratorExample()
	factory2.PrototypeExample()

}

func runPrototype() {
	printPatternDescription("To implement a prototype, partially construct an object and store it somewhere." +
		"Deep copy the prototype. Customize the resulting instance. A prototype factory provides a convenient API for using prototypes.")
	//prototype.DeepCopyExample()
	prototype2.SerializationExample()
	prototype2.FactoryExample()
}

func runSingleton() {
	singleton.SingletonExample()
	singleton.DependencyInversionSingletonExample()
}

func runAdapter() {
	adapter.AdapterExample()
}

func runBridge() {
	bridge.BridgeExample()
}

func runComposite() {
	composite.GeometricShapesExamples()
}

func printPatternDescription(desc string) {
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println()
	fmt.Println(desc)
	fmt.Println()
	fmt.Println(strings.Repeat("-", 100))
}
