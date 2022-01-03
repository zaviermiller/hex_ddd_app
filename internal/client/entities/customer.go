package entities

type Customer struct {
	ID           uint   `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	SignedWaiver bool   `json:"signed_waiver"`
	// Appointments []Appointment
	// AppointmentID string `json:"appointment_id"`

	// appt_cache Appointment
}

func (c *Customer) FullName() string {
	return c.FirstName + " " + c.LastName
}

func (c *Customer) ScheduleAppointment() {}
