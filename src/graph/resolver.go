package graph

import (
	"github.com/Edilberto-Vazquez/website-services/src/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db repository.DBRepository
}

func NewResolver(db repository.DBRepository) *Resolver {
	return &Resolver{
		db: db,
	}
}
