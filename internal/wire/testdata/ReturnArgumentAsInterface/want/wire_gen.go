// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"fmt"
)

// Injectors from wire.go:

func injectStringer(s MyString) fmt.Stringer {
	return s
}
