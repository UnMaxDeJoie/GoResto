package entities

type Truck struct {
    ID         int       `json:"id"`
    Name       string    `json:"name"`
    SlotBuffer uint8     `json:"slot_buffer"`
    Opening    string `json:"opening"`
    Closing    string `json:"closing"`
    UserID     int       `json:"user_id"`
}