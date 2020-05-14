package controller

import (
	"github.com/example/my-operator/pkg/controller/visitorsapp"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, visitorsapp.Add)
}
