package graph

import "github.com/Edilberto-Vazquez/website-services/models"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	profile  *models.Profile
	projects []*models.Project
	jobs     []*models.Job
}
