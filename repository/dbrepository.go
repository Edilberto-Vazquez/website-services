package repository

import (
	"context"

	"github.com/Edilberto-Vazquez/website-services/models"
)

type DBRepository interface {
	GetProfile(ctx context.Context, lang string) (profile models.Profile, err error)
	GetProjects(ctx context.Context, lang string) (projects models.Projects, err error)
	GetJobs(ctx context.Context, lang string) (jobs models.Jobs, err error)
}

var implementedDB DBRepository

func SetImplementedDB(implementation DBRepository) {
	implementedDB = implementation
}

func GetProfile(ctx context.Context, lang string) (profile models.Profile, err error) {
	return implementedDB.GetProfile(ctx, lang)
}

func GetProjects(ctx context.Context, lang string) (projects models.Projects, err error) {
	return implementedDB.GetProjects(ctx, lang)
}

func GetJobs(ctx context.Context, lang string) (jobs models.Jobs, err error) {
	return implementedDB.GetJobs(ctx, lang)
}
