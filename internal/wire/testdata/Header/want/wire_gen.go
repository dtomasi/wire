// This is a sample header file.
//
// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func injectFoo() Foo {
	foo := provideFoo()
	return foo
}
