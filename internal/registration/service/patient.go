package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"xamss.onelab.final/internal/registration/domain"
	"xamss.onelab.final/pkg/util"
)

func (m *Manager) CreateAccount(ctx context.Context, u *domain.User) error {
	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	err = m.Repository.CreateAccount(ctx, u)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) Login(ctx context.Context, username, password string) (string, error) {
	user, err := m.Repository.GetAccount(ctx, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("user not found")
		}

		return "", fmt.Errorf("get user err: %w", err)
	}

	err = util.CheckPassword(password, user.Password)
	if err != nil {
		return "", fmt.Errorf("incorrect password: %w", err)
	}

	accessToken, err := m.Token.CreateToken(user.ID, m.Config.Token.TimeToLive)
	if err != nil {
		return "", fmt.Errorf("create token err: %w", err)
	}

	return accessToken, nil
}

func (m *Manager) VerifyToken(token string) (int64, error) {
	payload, err := m.Token.VerifyToken(token)
	if err != nil {
		return 0, fmt.Errorf("validate token err: %w", err)
	}

	return payload.UserID, nil
}

func (m *Manager) CreateAppointment(ctx context.Context, a *domain.Appointment) error {
	err := m.Repository.CreateAppointment(ctx, a)
	if err != nil {
		return err
	}
	return nil
}
func (m *Manager) GetAppointment(ctx context.Context, id int64) (*domain.Appointment, error) {
	appointment, err := m.Repository.GetAppointment(ctx, id)
	if err != nil {
		return nil, err
	}
	return appointment, nil
}

func (m *Manager) GetAppointmentsByAccountID(ctx context.Context, id int64) ([]domain.Appointment, error) {
	appointments, err := m.Repository.GetAppointmentsByAccountID(ctx, id)
	if err != nil {
		return nil, err
	}
	return appointments, nil
}
func (m *Manager) GetDoctorsBySpeciality(ctx context.Context, speciality string) ([]domain.User, error) {
	doctors, err := m.Repository.GetDoctorsBySpeciality(ctx, speciality)
	if err != nil {
		return nil, err
	}
	return doctors, nil
}
