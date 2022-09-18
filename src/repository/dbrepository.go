package repository

import (
	"context"

	"github.com/Edilberto-Vazquez/website-services/src/models"
)

type DBRepository interface {
	GetProfile(ctx context.Context, lang string) (*models.Profile, error)
	GetProjects(ctx context.Context, lang string) ([]*models.Project, error)
	GetJobs(ctx context.Context, lang string) ([]*models.Job, error)
}

var implementedDB DBRepository

func SetImplementedDB(implementation DBRepository) {
	implementedDB = implementation
}

func GetProfile(ctx context.Context, lang string) (*models.Profile, error) {
	return implementedDB.GetProfile(ctx, lang)
}

func GetProjects(ctx context.Context, lang string) ([]*models.Project, error) {
	return implementedDB.GetProjects(ctx, lang)
}

func GetJobs(ctx context.Context, lang string) ([]*models.Job, error) {
	return implementedDB.GetJobs(ctx, lang)
}
