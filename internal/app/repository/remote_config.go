package repository

import (
    "google.golang.org/api/firebaseremoteconfig/v1"
)

type Repository interface {
    GetTemplate() (*firebaseremoteconfig.RemoteConfig, error)
}

type firebaseRepository struct {
    client *firebaseremoteconfig.Service
}

func NewFirebaseRepository(rc *firebaseremoteconfig.Service) *firebaseRepository {
    return &firebaseRepository{
        client: rc,
    }
}

func (fr *firebaseRepository) GetTemplate() (*firebaseremoteconfig.RemoteConfig, error) {
    return fr.client.Projects.GetRemoteConfig("cms-content-generator-3305b").Do()
}
