package handler

import (
	"net/http"

	"github.com/elingsuryo/movie-app/internal/http/dto"
	"github.com/elingsuryo/movie-app/internal/service"
	"github.com/elingsuryo/movie-app/pkg/response"
	"github.com/labstack/echo/v4"
)


type UserHandler struct {
	tokenService service.TokenService
	userService	service.UserService
}

func NewUserHandler(tokenService service.TokenService, userService	service.UserService) UserHandler {
	return UserHandler{tokenService, userService}
}

func (h *UserHandler) Login(ctx echo.Context) error {
	var req dto.UserLoginRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	
	claims, err := h.userService.Login(ctx.Request().Context(), req.Username, req.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	
	token, err := h.tokenService.GenerateAccessToken(ctx.Request().Context(), *claims)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully login", map[string]string{"token": token}))
}

func (h *UserHandler) Register(ctx echo.Context) error {
	var req dto.UserRegisterRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	err := h.userService.Register(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}


	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully register", nil))
}

func (h *UserHandler) GetUsers(ctx echo.Context) error {
	users, err := h.userService.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully showing all users", users))
}

func (h* UserHandler) GetByID(ctx echo.Context) error{
	var req dto.GetUserByIDRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

user, err := h.userService.GetByID(ctx.Request().Context(), req.ID)
if err != nil {
	return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
}
return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully showing User", user))

}

func (h *UserHandler) CreateUser(ctx echo.Context) error {
	var req dto.UserCreateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	err := h.userService.Create(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully creating User", nil))

}

func (h *UserHandler) UpdateUser(ctx echo.Context) error {
	var req dto.UserUpdateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	err := h.userService.Update(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully updating User", nil))
}


func (h *UserHandler) DeleteUser(ctx echo.Context) error {
	var req dto.GetMovieByIDRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	movie,err :=  h.userService.GetByID(ctx.Request().Context(), req.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	err = h.userService.Delete(ctx.Request().Context(), movie)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully deleting User", nil))
}

func (h *UserHandler) ResetPassword(ctx echo.Context) error {
	var req dto.ResetPasswordRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	err := h.userService.ResetPassword(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully reset password", nil))
}

// func (h *UserHandler) ResetPasswordRequest(ctx echo.Context) error {
// 	var req dto.RequestResetPassword
// 	if err := ctx.Bind(&req); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
// 	}

// 	err :=h.userService.RequestResetPassword(ctx.Request().Context(), req.Username)
// 	if err != nil {
// 		return ctx.JSON(http.StatusInternalS erverError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
// 	}

// 	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully request reset password", nil))
// }

func (h *UserHandler) VerifyEmail(ctx echo.Context) error {
	var req dto.VerifyEmailRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	err:= h.userService.VerifyEmail(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully verify email", nil))
}