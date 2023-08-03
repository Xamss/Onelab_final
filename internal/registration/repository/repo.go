package repository

import (
	"context"
	"xamss.onelab.final/internal/registration/domain"
)

type Repository interface {
	CreateAccount(ctx context.Context, u *domain.User) error
	GetAccount(ctx context.Context, username string) (*domain.User, error)

	CreateAppointment(ctx context.Context, a *domain.Appointment) error
	GetAppointment(ctx context.Context, id int64) (*domain.Appointment, error)
	GetAppointmentsByAccountID(ctx context.Context, id int64) ([]domain.Appointment, error)
	GetDoctorsBySpeciality(ctx context.Context, speciality string) ([]domain.User, error)
}
