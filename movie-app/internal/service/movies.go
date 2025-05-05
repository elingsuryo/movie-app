package service

import (
	"context"

	"github.com/elingsuryo/movie-app/internal/entity"
	"github.com/elingsuryo/movie-app/internal/http/dto"
	"github.com/elingsuryo/movie-app/internal/repository"
)

type MovieService interface {
	GetAll(ctx context.Context) ([]entity.Movie, error)
	GetByID(ctx context.Context, id int64) (*entity.Movie, error)
	Insert(ctx context.Context, req dto.CreateMovieRequest) error
	Update(ctx context.Context, req dto.UpdateMovieRequest) error
    Delete(ctx context.Context, movies *entity.Movie) error

}


type movieService struct {
	moviesRepository repository.MoviesRepository
}

func NewMovieService(moviesRepository repository.MoviesRepository) MovieService {
	return &movieService{moviesRepository}
}

 func (s movieService ) GetAll(ctx context.Context) ([]entity.Movie, error){
	 return s.moviesRepository.GetAll(ctx)
 }

func (s movieService) GetByID(ctx context.Context, id int64) (*entity.Movie, error){
	 return s.moviesRepository.GetByID(ctx, id)
}

func (s movieService) Insert(ctx context.Context, req dto.CreateMovieRequest) error{
	movie := &entity.Movie{
		Title: req.Title,
		Year: req.Year,	
		Director: req.Director,
		Description: req.Description,
	}
	return s.moviesRepository.Insert(ctx, movie)
}

func (s movieService) Update(ctx context.Context, req dto.UpdateMovieRequest) error{
	movies, err := s.moviesRepository.GetByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if req.Title != "" {
		movies.Title = req.Title
	}
	if req.Year != 0 {
		movies.Year = req.Year
	}
	if req.Director != "" {
		movies.Director = req.Director
	}
	if req.Description != "" {
		movies.Description = req.Description
	}
	return s.moviesRepository.Update(ctx, movies)
}

func (s movieService) Delete(ctx context.Context, movies *entity.Movie) error{
	return s.moviesRepository.Delete(ctx, movies)
}