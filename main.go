package main

import (
	"fmt"
	"github.com/voicurobert/design_patterns_in_go/section_01_solid"
	builder "github.com/voicurobert/design_patterns_in_go/section_02_builder"
	factory "github.com/voicurobert/design_patterns_in_go/section_03_factories"
	prototype "github.com/voicurobert/design_patterns_in_go/section_04_prototype"
	"strings"
)

func main() {

	//runSolid()

	//runBuilder()

	//runFactory()

	runPrototype()
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
	builder.StringBuilderExample()
	builder.Example()
	builder.FacetsExample()
	builder.ParameterExample()
	builder.FunctionalExample()
}

func runFactory() {
	printPatternDescription("A factory function (constructor) is a helper function for making struct instances. \n" +
		"A factory is any entity than can take care of object creation. Can be a function or a dedicated struct.")
	factory.FunctionExample()
	factory.InterfaceExample()
	factory.GeneratorExample()
	factory.PrototypeExample()

}

func runPrototype() {
	printPatternDescription("To implement a prototype, partially construct an object and store it somewhere." +
		"Deep copy the prototype. Customize the resulting instance. A prototype factory provides a convenient API for using prototypes.")
	//prototype.DeepCopyExample()
	prototype.SerializationExample()
	prototype.FactoryExample()
}

func printPatternDescription(desc string) {
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println()
	fmt.Println(desc)
	fmt.Println()
	fmt.Println(strings.Repeat("-", 100))
}
