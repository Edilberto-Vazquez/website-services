package graph

import (
	"context"
	"fmt"

	"github.com/Edilberto-Vazquez/website-services/src/graph/generated"
	"github.com/Edilberto-Vazquez/website-services/src/i18n"
	"github.com/Edilberto-Vazquez/website-services/src/models"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

// Profile is the resolver for the profile field.
func (r *queryResolver) Profile(ctx context.Context, input models.Languaje) (*models.Profile, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	value, exist := i18n.CheckLanguage(input.Lang)
	if !exist {
		return nil, gqlerror.Errorf("The app cannot provide that language")
	}
	profile, err := r.db.GetProfile(gc, value)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

// Projects is the resolver for the projects field.
func (r *queryResolver) Projects(ctx context.Context, input models.Languaje) ([]*models.Project, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	value, exist := i18n.CheckLanguage(input.Lang)
	if !exist {
		return nil, gqlerror.Errorf("The app cannot provide that language")
	}
	projects, err := r.db.GetProjects(gc, value)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

// Jobs is the resolver for the jobs field.
func (r *queryResolver) Jobs(ctx context.Context, input models.Languaje) ([]*models.Job, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	value, exist := i18n.CheckLanguage(input.Lang)
	if !exist {
		return nil, gqlerror.Errorf("The app cannot provide that language")
	}
	jobs, err := r.db.GetJobs(gc, value)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
