package domain

type Customer struct {
	ID           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	SignedWaiver bool   `json:"signed_waiver"`
	// AppointmentID string `json:"appointment_id"`

	// appt_cache Appointment
}

func (c *Customer) FullName() string {
	return c.FirstName + " " + c.LastName
}
