package postgres

import (
	"file-storage/internal/domain/file_info"
	"fmt"

	"gorm.io/gorm"
)

type FileInfoPostgresRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *FileInfoPostgresRepository {
	return &FileInfoPostgresRepository{db: db}
}

func (r *FileInfoPostgresRepository) Create(fileInfo file_info.FileInfo) (file_info.FileInfo, error) {
	if err := r.db.Create(&fileInfo).Error; err != nil {
		err := fmt.Errorf("failed creating file info: %w", err)

		return fileInfo, file_info.NewRepositoryError(err)
	}

	return fileInfo, nil
}

func (r *FileInfoPostgresRepository) FindAll() ([]file_info.FileInfo, error) {
	fileInfos := []file_info.FileInfo{}

	if err := r.db.Find(&fileInfos).Error; err != nil {
		err := fmt.Errorf("failed finding file infos: %w", err)

		return fileInfos, file_info.NewRepositoryError(err)
	}

	return fileInfos, nil
}

func (r *FileInfoPostgresRepository) FindByID(id uint) (file_info.FileInfo, error) {
	fileInfo := file_info.FileInfo{}

	if err := r.db.Find(&fileInfo, "id = ?", id).Error; err != nil {
		err := fmt.Errorf("failed finding file info by id=%v: %w", id, err)

		return fileInfo, file_info.NewRepositoryError(err)
	}

	if fileInfo.ID == 0 {
		return fileInfo, file_info.NewNotFoundError(id)
	}

	return fileInfo, nil
}

func (r *FileInfoPostgresRepository) Update(fileInfo file_info.FileInfo) (file_info.FileInfo, error) {
	if err := r.db.Save(&fileInfo).Error; err != nil {
		err := fmt.Errorf("failed updating file info (id=%v): %w", fileInfo.ID, err)

		return fileInfo, file_info.NewRepositoryError(err)
	}

	return fileInfo, nil
}

func (r *FileInfoPostgresRepository) Delete(fileInfo file_info.FileInfo) (file_info.FileInfo, error) {
	if err := r.db.Delete(&fileInfo).Error; err != nil {
		err := fmt.Errorf("failed deleting file info (id=%v): %w", fileInfo.ID, err)

		return fileInfo, file_info.NewRepositoryError(err)
	}

	return fileInfo, nil
}
