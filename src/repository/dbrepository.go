package repository

import (
	"context"

	"github.com/Edilberto-Vazquez/website-services/src/models"
)

type DBRepository interface {
	GetProfile(ctx context.Context, lang string) (*models.Profile, error)
	GetProjects(ctx context.Context, lang string) ([]*models.Project, error)
	GetJobs(ctx context.Context, lang string) ([]*models.Job, error)
	GetFullProfile(ctx context.Context, lang string) (*models.FullProfile, error)
}
