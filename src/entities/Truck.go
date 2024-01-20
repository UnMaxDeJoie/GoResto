package entities

import "time"

type Truck struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	SlotBuffer uint8     `json:"slot_buffer"`
	Opening    time.Time `json:"opening"`
	Closing    time.Time `json:"closong"`
}
