package main

import "errors"

type artifact struct {
	id string
}

func (a artifact) BuilderId() string {
	return "osbuild.image-builder"
}

func (a artifact) Files() []string {
	return nil
}

func (a artifact) Id() string {
	return a.id
}

func (a artifact) String() string {
	return a.id
}

func (a artifact) State(name string) interface{} {
	return nil
}

func (a artifact) Destroy() error {
	return errors.New("image-builder does not support destroying artifacts")
}
