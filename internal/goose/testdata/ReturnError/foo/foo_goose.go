//+build gooseinject

package main

import (
	"codename/goose"
)

func injectFoo() (Foo, error) {
	panic(goose.Use(Set))
}
