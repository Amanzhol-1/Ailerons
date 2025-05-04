package domain

type CustomerProfile struct {
	UserID      int64  `json:"user_id"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
}

type DriverProfile struct {
	UserID        int64  `json:"user_id"`
	LicenseNumber string `json:"license_number"`
	VehicleInfo   string `json:"vehicle_info"`
}
