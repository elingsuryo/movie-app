package repository

import (
	"context"

	"github.com/elingsuryo/movie-app/internal/entity"
	"gorm.io/gorm"
)

type MoviesRepository interface {
    GetByID(ctx context.Context, id int64) (*entity.Movie, error)
    GetAll(ctx context.Context) ([]entity.Movie, error)
    Insert(ctx context.Context, movies *entity.Movie) error
    Update(ctx context.Context, movies *entity.Movie) error
    Delete(ctx context.Context, movies *entity.Movie) error
}

type moviesRepository struct {
	db *gorm.DB
}

func NewMoviesRepository(db *gorm.DB) MoviesRepository {
	return &moviesRepository{db}
}

func (u moviesRepository) GetByID(ctx context.Context, id int64) (*entity.Movie, error) {
    result := new(entity.Movie)

    if err := u.db.WithContext(ctx).Where("id = ?", id).First(&result).Error; err != nil {
        return nil, err
    }
    return result, nil
}

func (u moviesRepository) GetAll(ctx context.Context) ([]entity.Movie, error) {
    result := make([]entity.Movie, 0)

    if err := u.db.WithContext(ctx).Find(&result).Error; err != nil {
        return nil, err 
    }
    return result, nil
}

func (u moviesRepository) Insert(ctx context.Context, movies *entity.Movie) error {
    return u.db.WithContext(ctx).Create(&movies).Error
}

func (u moviesRepository) Update(ctx context.Context, movies *entity.Movie) error {
    return u.db.WithContext(ctx).Updates(&movies).Error
}

func (u moviesRepository) Delete(ctx context.Context, movies *entity.Movie) error {
    return u.db.WithContext(ctx).Delete(&movies).Error
}
