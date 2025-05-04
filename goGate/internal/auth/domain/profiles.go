package domain

type CustomerProfile struct {
	UserID      int64  `json:"user_id"`
	Role        string `json:"role"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
}

type DriverProfile struct {
	UserID        int64  `json:"user_id"`
	Role          string `json:"role"`
	LicenseNumber string `json:"license_number"`
	VehicleInfo   string `json:"vehicle_info"`
}
