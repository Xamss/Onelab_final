package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"xamss.onelab.final/api"
	"xamss.onelab.final/internal/registration/domain"
	"xamss.onelab.final/internal/registration/service/mock"
)

func TestRegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockService(ctrl)

	handler := NewHandler(mockService)

	tests := []struct {
		req         *api.RegisterRequest
		status      int
		mockReturn  error
		requestJSON interface{}
	}{
		{
			&api.RegisterRequest{
				Username:  "aa",
				FirstName: "bb",
				LastName:  "cc",
				Email:     "ee",
				Password:  "dd",
			},
			http.StatusCreated,
			nil,
			&api.RegisterRequest{
				Username:  "aa",
				FirstName: "bb",
				LastName:  "cc",
				Email:     "ee",
				Password:  "dd",
			},
		},
		{
			&api.RegisterRequest{
				Username:  "aa",
				FirstName: "bb",
				LastName:  "cc",
				Email:     "211022@astanait.edu.kz",
				Password:  "dd",
			},
			http.StatusInternalServerError,
			errors.New("Email is duplicated! SQL error"),
			&api.RegisterRequest{
				Username:  "aa",
				FirstName: "bb",
				LastName:  "cc",
				Email:     "211022@astanait.edu.kz",
				Password:  "dd",
			},
		},
		//{
		//	&api.RegisterRequest{
		//		FirstName: "bb",
		//		LastName:  "cc",
		//		Email:     "ee",
		//		Password:  "dd",
		//	},
		//	http.StatusBadRequest,
		//	nil,
		//	struct {
		//		Username string
		//	}{
		//		Username: "aa",
		//	},
		//},
	}

	for _, test := range tests {
		recorder := httptest.NewRecorder()
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(test.requestJSON)
		require.NoError(t, err)

		u := &domain.User{
			Username:  test.req.Username,
			FirstName: test.req.FirstName,
			LastName:  test.req.LastName,
			Email:     test.req.Email,
			Password:  test.req.Password,
		}

		mockService.EXPECT().CreateAccount(gomock.Any(), u).Return(test.mockReturn).Times(1)

		url := fmt.Sprintf("/signup")
		request, err := http.NewRequest(http.MethodPost, url, &buf)
		require.NoError(t, err)

		handler.InitRouter().ServeHTTP(recorder, request)

		require.Equal(t, test.status, recorder.Code)
	}
}
