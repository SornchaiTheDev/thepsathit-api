package usecases

import (
	"log"

	"github.com/SornchaiTheDev/thep-api/internal/entities"
	"github.com/SornchaiTheDev/thep-api/internal/repositories"
)

type senserUseCaseImpl struct {
	repo repositories.SensorRepo
}

type SensorUsecase interface {
	AddNewData(s *entities.Sensor, treeId uint) error
}

func newSensorUsecase(repo repositories.SensorRepo) SensorUsecase {
	return &senserUseCaseImpl{repo}
}

func (s senserUseCaseImpl) AddNewData(sensor *entities.Sensor, treeId uint) error {
	err := s.repo.Add(sensor, treeId)
	if err != nil {
		log.Fatal("Can't add new data to the tree")
	}

	return nil
}
