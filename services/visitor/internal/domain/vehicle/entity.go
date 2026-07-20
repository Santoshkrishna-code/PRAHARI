package vehicle

import "time"

// Detail represents visitor vehicle registration (license plate, type, authorization).
type Detail struct {
	ID           string    `json:"id"`
	VisitID      string    `json:"visit_id"`
	LicensePlate string    `json:"license_plate"`
	VehicleType  string    `json:"vehicle_type"` // CAR, TRUCK, VAN, MOTORCYCLE
	DriverName   string    `json:"driver_name"`
	Approved     bool      `json:"approved"`
	CreatedAt    time.Time `json:"created_at"`
}
