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

// FullProfile is the resolver for the fullProfile field.
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

func (r *queryResolver) FullProfile(ctx context.Context, lang string) (*models.FullProfile, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	value, exist := i18n.CheckLanguage(lang)
	if !exist {
		return nil, gqlerror.Errorf("The app cannot provide that language")
	}
	profile, err := r.db.GetFullProfile(gc, value)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
