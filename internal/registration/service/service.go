package service

import (
	"context"
	"xamss.onelab.final/internal/registration/domain"
)

type Service interface {
	CreateAccount(ctx context.Context, u *domain.User) error
	Login(ctx context.Context, username, password string) (string, error)
	VerifyToken(token string) (int64, error)

	CreateAppointment(ctx context.Context, a *domain.Appointment) error
	GetAppointment(ctx context.Context, id int64) (*domain.Appointment, error)
	GetAppointmentsByAccountID(ctx context.Context, id int64) ([]domain.Appointment, error)
	GetDoctorsBySpeciality(ctx context.Context, speciality string) ([]domain.User, error)
}
