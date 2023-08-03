package pgrepo

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"strings"
	"xamss.onelab.final/internal/registration/domain"
)

func (p *Postgres) CreateAccount(ctx context.Context, u *domain.User) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			                username, 
			                firstname, 
			                lastname, 
			                email,
			                password 
			                )
		VALUES ($1, $2, $3, $4, $5)
	`, accountTable)

	_, err := p.Pool.Exec(ctx, query, u.Username, u.FirstName, u.LastName, u.Email, u.Password)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetAccount(ctx context.Context, username string) (*domain.User, error) {
	user := new(domain.User)

	query := fmt.Sprintf("SELECT id, username, firstname, lastname, email, password FROM %s WHERE username = $1", accountTable)

	err := pgxscan.Get(ctx, p.Pool, user, query, strings.TrimSpace(username))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *Postgres) CreateAppointment(ctx context.Context, a *domain.Appointment) error {
	query := fmt.Sprintf(`
		INSERT INTO %s ( 
			                patient_id, 
			                doctor_id, 
			                time
			                )
		VALUES ($1, $2, $3)
	`, appointmentTable)

	_, err := p.Pool.Exec(ctx, query, a.PatientID, a.DoctorID, a.Time)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetAppointment(ctx context.Context, id int64) (*domain.Appointment, error) {
	app := new(domain.Appointment)

	query := fmt.Sprintf("SELECT id, patient_id, doctor_id, time FROM %s WHERE id = $1", appointmentTable)

	err := pgxscan.Get(ctx, p.Pool, app, query, id)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (p *Postgres) GetAppointmentsByAccountID(ctx context.Context, id int64) ([]domain.Appointment, error) {
	query := fmt.Sprintf(`
		SELECT id, patient_id, doctor_id, time
		FROM %s
		WHERE patient_id = $1
	`, appointmentTable)

	rows, err := p.Pool.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var appointments []domain.Appointment
	for rows.Next() {
		var appointment domain.Appointment
		err := rows.Scan(
			&appointment.ID,
			&appointment.PatientID,
			&appointment.DoctorID,
			&appointment.Time,
		)
		if err != nil {
			return nil, err
		}

		appointments = append(appointments, appointment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return appointments, nil
}

func (p *Postgres) GetDoctorsBySpeciality(ctx context.Context, speciality string) ([]domain.User, error) {
	query := `
		SELECT id, username, firstname, lastname, email
		FROM accounts inner join account_roles
		ON accounts.id = account_roles.account_id inner join roles
		ON roles.id = account_roles.role_id
		WHERE speciality = $1
	`

	rows, err := p.Pool.Query(ctx, query, strings.TrimSpace(speciality))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var doctors []domain.User
	for rows.Next() {
		var doctor domain.User
		err := rows.Scan(
			&doctor.ID,
			&doctor.Username,
			&doctor.FirstName,
			&doctor.LastName,
			&doctor.Email,
		)
		if err != nil {
			return nil, err
		}

		doctors = append(doctors, doctor)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return doctors, nil
}
