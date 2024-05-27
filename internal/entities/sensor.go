package entities

type Sensor struct {
	Ground_temp     float32 `json:"ground_temp"`
	Ground_humidity float32 `json:"ground_humidity"`
	Air_temp        float32 `json:"air_temp"`
	Air_humidity    float32 `json:"air_humidity"`
}
