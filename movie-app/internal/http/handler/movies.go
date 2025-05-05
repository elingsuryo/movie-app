package handler

import (
	"net/http"

	"github.com/elingsuryo/movie-app/internal/http/dto"
	"github.com/elingsuryo/movie-app/internal/service"
	"github.com/elingsuryo/movie-app/pkg/response"
	"github.com/labstack/echo/v4"
)

type MovieHandler struct {
	movieService service.MovieService
}

func NewMovieHandler(movieService service.MovieService) MovieHandler {
	return MovieHandler{movieService}
}

func (h *MovieHandler) GetAllMovies(ctx echo.Context) error {
	users, err := h.movieService.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully showing all movies", users))
}

func (h* MovieHandler) GetMovie(ctx echo.Context) error{
	var req dto.GetMovieByIDRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

movie, err := h.movieService.GetByID(ctx.Request().Context(), req.ID)
if err != nil {
	return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
}
return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully showing movie", movie))

}

func (h *MovieHandler) CreateMovie(ctx echo.Context) error {
	var req dto.CreateMovieRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	err := h.movieService.Insert(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully creating movie", nil))

}

func (h *MovieHandler) UpdateMovie(ctx echo.Context) error {
	var req dto.UpdateMovieRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	err := h.movieService.Update(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully updating movie", nil))
}

func (h *MovieHandler) DeleteMovie(ctx echo.Context) error {
	var req dto.GetMovieByIDRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	movie,err :=  h.movieService.GetByID(ctx.Request().Context(), req.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	err = h.movieService.Delete(ctx.Request().Context(), movie)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse("successfully deleting movie", nil))
}