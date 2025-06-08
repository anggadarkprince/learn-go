//go:build wireinject
// +build wireinject

package simple

import (
	"io"
	"os"
	"github.com/google/wire"
)

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository,
		NewSimpleService,
	)
	return nil, nil
}

func InitializedDatabaseRespository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMongoDB,
		NewDatabaseMySQL,
		NewDatabaseRepository,
	)
	return nil
}

// Provider sets
var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FooBarService {
	wire.Build(
		fooSet,
		barSet,
		NewFooBarService,
	)
	return nil
}


// Interface provider
var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)), // new: to make pointer
)
func InitializedHelloService() *HelloService {
	wire.Build(
		helloSet,
 		NewHelloService,
	)
	return nil
}

// Struct provider
func InitializeFooBar() *FooBar {
	wire.Build(
		NewFoo,
		NewBar,
		//wire.Struct(new(FooBar), "*"),
		wire.Struct(new(FooBar), "Foo", "Bar"),
	)
	return nil
}

// Binding value
var fooValue = &Foo{}
var barValue = &Bar{}

func InitializeFooBarUsingValue() *FooBar {
	wire.Build(
		wire.Value(fooValue),
		wire.Value(barValue),
		wire.Struct(new(FooBar), "Foo", "Bar"),
	)
	return nil
}

func InitializeReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

// Struct field provider
func InitializedConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Config"),
	)
	return nil
}

// Cleanup provider
func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(
		NewFile,
		NewConnection,
	)
	return nil, nil
}