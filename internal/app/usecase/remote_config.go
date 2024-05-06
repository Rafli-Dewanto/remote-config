package usecase

import (
	"cms-config/internal/app/repository"
	"fmt"
)

type RemoteConfigUsecase interface {
    GetTemplate() (string, error)
}

type usecase struct {
    repo repository.Repository
}

func NewUsecase(repo repository.Repository) RemoteConfigUsecase {
    return &usecase{
        repo: repo,
    }
}

func (uc *usecase) GetTemplate() (string, error) {
    // Call repository method to get template
    template, err := uc.repo.GetTemplate()
    if err != nil {
        return "", err
    }
		fmt.Println(template)
		return "", nil
}

// func processTemplate(template *firebaseremoteconfig.RemoteConfig) string {
//     // Implement logic to process the template response and convert to string
// }
