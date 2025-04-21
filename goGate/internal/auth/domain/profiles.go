package domain

type CustomerProfile struct {
	UserID      int64  `json:"-"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
}

type DriverProfile struct {
	UserID        int64  `json:"-"`
	LicenseNumber string `json:"license_number"`
	VehicleInfo   string `json:"vehicle_info"`
}
