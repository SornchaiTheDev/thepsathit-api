package repositories

import "github.com/SornchaiTheDev/thep-api/internal/entities"

type SensorRepo interface {
	Add(s *entities.Sensor, treeId uint) error
	GetByTreeId(id uint) ([]entities.Sensor, error)
}
