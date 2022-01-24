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
	decorator "github.com/voicurobert/design_patterns_in_go/app/src/section_09_decorator"
	facade "github.com/voicurobert/design_patterns_in_go/app/src/section_10_facade"
	flyweight "github.com/voicurobert/design_patterns_in_go/app/src/section_11_flyweight"
	proxy "github.com/voicurobert/design_patterns_in_go/app/src/section_12_proxy"
	chain_of_responsability "github.com/voicurobert/design_patterns_in_go/app/src/section_13_chain_of_responsibility"
	command "github.com/voicurobert/design_patterns_in_go/app/src/section_14_command"
	interpreter "github.com/voicurobert/design_patterns_in_go/app/src/section_15_interpreter"
	iterator "github.com/voicurobert/design_patterns_in_go/app/src/section_16_iterator"
	mediator "github.com/voicurobert/design_patterns_in_go/app/src/section_17_mediator"
	memento "github.com/voicurobert/design_patterns_in_go/app/src/section_18_memento"
	observer "github.com/voicurobert/design_patterns_in_go/app/src/section_19_observer"
	state "github.com/voicurobert/design_patterns_in_go/app/src/section_20_state"
	strategy "github.com/voicurobert/design_patterns_in_go/app/src/section_21_strategy"
	template_method "github.com/voicurobert/design_patterns_in_go/app/src/section_22_template_method"
	visitor "github.com/voicurobert/design_patterns_in_go/app/src/section_23_visitor"
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

	//runComposite()

	//runDecorator()

	//runFacade()

	//runFlyweight()

	//runProxy()

	//runChainOfResponsibility()

	//runCommand()

	//runInterpreter()

	//runIterator()

	//runMediator()

	//runMemento()

	//runObserver()

	//runState()

	//runStrategy()

	//runTemplateMethod()

	runVisitor()
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

func runDecorator() {
	decorator.MultipleAggregationExample()
	decorator.DecoratorExample()
}

func runFacade() {
	facade.FacadeExample()
}

func runFlyweight() {
	flyweight.TextFormattingExample()
	flyweight.UserNamesExample()
}

func runProxy() {
	proxy.ProtectionProxyExample()
	proxy.VirtualProxyExample()
}

func runChainOfResponsibility() {
	chain_of_responsability.MethodChainExample()
	chain_of_responsability.BrokerChainExample()
}

func runCommand() {
	command.Example()
}

func runInterpreter() {
	interpreter.Example()
}

func runIterator() {
	iterator.IterationExample()
	iterator.TreeTraversalExample()
}

func runMediator() {
	mediator.ChatRoomExample()
}

func runMemento() {
	memento.Example()
	fmt.Println()
	memento.UndeRedoExample()
}

func runObserver() {
	observer.ObserverAndObservableExample()
	fmt.Println()
	observer.PropertyObserversExample()
}

func runState() {
	//state.ClassicExample()
	//state.HandmadeStateMachineExample()
	state.SwitchBasedStateMachineExample()
}

func runStrategy() {
	strategy.Example()
}

func runTemplateMethod() {
	template_method.Example()
}

func runVisitor() {
	visitor.IntrusiveVisitorExample()
}

func printPatternDescription(desc string) {
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println()
	fmt.Println(desc)
	fmt.Println()
	fmt.Println(strings.Repeat("-", 100))
}
