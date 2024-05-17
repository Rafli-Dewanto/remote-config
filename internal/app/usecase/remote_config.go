package usecase

import (
	"cms-config/internal/pkg/util"

	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"google.golang.org/api/firebaseremoteconfig/v1"
)

type RemoteConfigUseCase struct {
	firebaseClient *firebaseremoteconfig.Service
}

func NewUsecase(firebaseClient *firebaseremoteconfig.Service) *RemoteConfigUseCase {
	return &RemoteConfigUseCase{
		firebaseClient: firebaseClient,
	}
}

func (uc *RemoteConfigUseCase) GetOauthToken() (*oauth2.Token, error) {
	token, err := util.ServiceAccount("service-account.json")
	if err != nil {
		log.Fatalf("Error acquiring token: %v", err)
	}
	return token, nil
}
