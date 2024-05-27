package postgre

import (
	"database/sql"
	"log"

	"github.com/SornchaiTheDev/thep-api/internal/entities"
	"github.com/SornchaiTheDev/thep-api/internal/repositories"
)

type sqlSensorRepo struct {
	db *sql.DB
}

func newSensorRepo(db *sql.DB) repositories.SensorRepo {
	return sqlSensorRepo{db}
}

func (s sqlSensorRepo) Add(sensor *entities.Sensor, treeId uint) error {
	query := "INSERT INTO sensors (ground_temp,group_humidity,air_temp,air_humidity,tree_id) VALUES (?,?,?,?,?)"

	_, err := s.db.Exec(query, sensor.Ground_temp, sensor.Ground_humidity, sensor.Air_temp, sensor.Air_humidity, treeId)

	if err != nil {
		return err
	}

	return nil
}

func (s sqlSensorRepo) GetByTreeId(id uint) ([]entities.Sensor, error) {
	query := "SELECT * FROM sensors WHERE tree_id = ?"
	rows, err := s.db.Query(query, id)

	if err != nil {
		return nil, err
	}

	var sensors = make([]entities.Sensor, 0)

	for rows.Next() {
		var data = entities.Sensor{}
		rows.Scan(&data)
		sensors = append(sensors, data)
	}

	return sensors, nil
}
