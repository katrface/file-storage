package postgres

import (
	"file-storage/internal/domain/file_info"

	"gorm.io/gorm"
)

type FileInfoPostgresRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *FileInfoPostgresRepository {
	return &FileInfoPostgresRepository{db: db}
}

func (r *FileInfoPostgresRepository) Create(fileInfo file_info.FileInfo) (file_info.FileInfo, error) {
	r.db.Create(&fileInfo)
	return fileInfo, nil
}

func (r *FileInfoPostgresRepository) FindAll() ([]file_info.FileInfo, error) {
	fileInfos := []file_info.FileInfo{}

	r.db.Find(&fileInfos)

	return fileInfos, nil
}

func (r *FileInfoPostgresRepository) FindById(id uint) (file_info.FileInfo, error) {
	fileInfo := file_info.FileInfo{}

	r.db.Find(&fileInfo, "id = ?", id)

	return fileInfo, nil
}

func (r *FileInfoPostgresRepository) Update(fileInfo file_info.FileInfo) (file_info.FileInfo, error) {
	r.db.Save(&fileInfo)
	return fileInfo, nil
}

func (r *FileInfoPostgresRepository) Delete(fileInfo file_info.FileInfo) (file_info.FileInfo, error) {
	r.db.Delete(&fileInfo)
	return fileInfo, nil
}
