package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"xamss.onelab.final/api"
	"xamss.onelab.final/internal/registration/domain"
)

// createUser registration new user
//
//	@Summary      Create user
//	@Description  Create new user
//	@Tags         auth
//	@Accept       json
//	@Produce      json
//	@Param req body api.RegisterRequest true "req body"
//
//	@Success      201
//	@Failure      400  {object}  api.Error
//	@Failure      500  {object}  api.Error
//	@Router       /user/register [post]
func (h *Handler) createUser(ctx *gin.Context) {
	var req api.RegisterRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	u := &domain.User{
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}

	err = h.srvs.CreateAccount(ctx, u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

// loginUser registration new user
//
//	@Summary      User logs in
//	@Description  After successfully identifying oneself, the user is provided with AccessToken for authorization purposes
//	@Tags         auth
//	@Accept       json
//	@Produce      json
//	@Param req body api.LoginRequest true "req body"
//
//	@Success      200  {object}  api.Ok
//	@Failure      400  {object}  api.Error
//	@Failure      500  {object}  api.Error
//	@Router       /user/login [post]
func (h *Handler) loginUser(ctx *gin.Context) {
	var req api.LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	accessToken, err := h.srvs.Login(ctx, req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.Ok{
		Code:    0,
		Message: "success",
		Data:    accessToken,
	})
}

// createAppointment registration new user
//
//	@Summary      User creates a new appointment
//	@Description  Authorized user can create appointments
//	@Tags         appointment
//	@Accept       json
//	@Produce      json
//	@Param req body domain.Appointment true "req body"
//
//	@Success      201  {object}  api.Ok
//	@Failure      400  {object}  api.Error
//	@Failure      400  {object}  api.Error
//	@Failure      500  {object}  api.Error
//	@Router       /user/appointment/create [post]
func (h *Handler) createAppointment(ctx *gin.Context) {

	userID, ok := ctx.MustGet(authUserID).(int64)
	if !ok {
		log.Printf("can't get userID")
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "can't get user id from auth",
		})
		return
	}

	var req domain.Appointment

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	err = h.srvs.CreateAppointment(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    -3,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &api.Ok{
		Code:    0,
		Message: "success",
		Data:    userID,
	})
}
