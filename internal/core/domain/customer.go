package domain

type Customer struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	SignedWaiver  bool   `json:"signed_waiver"`
	AppointmentID string `json:"appointment_id"`

	appt_cache Appointment
}

func (c *Customer) Appointment() Appointment {
	return c.appt_cache
}
