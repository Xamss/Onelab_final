package domain

import "time"

type Appointment struct {
	ID        int64     `json:"id" db:"id"`
	PatientID int64     `json:"patient_id" db:"patient_id"`
	DoctorID  int64     `json:"doctor_id" db:"doctor_id"`
	Time      time.Time `json:"time" db:"time"`
}
