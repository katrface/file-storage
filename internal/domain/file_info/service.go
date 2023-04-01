package file_info

import "time"

type FileInfoRepository interface {
	FindAll() ([]FileInfo, error)
	FindByID(id uint) (FileInfo, error)
	Create(fileInfo FileInfo) (FileInfo, error)
	Update(fileInfo FileInfo) (FileInfo, error)
	Delete(fileInfo FileInfo) (FileInfo, error)
}

type FileInfoService struct {
	repo FileInfoRepository
}

func New(repo FileInfoRepository) *FileInfoService {
	return &FileInfoService{repo: repo}
}

func (s *FileInfoService) CreateFileInfo(fileInfo FileInfo) (FileInfo, error) {
	createdFileInfo, _ := s.repo.Create(fileInfo)

	return createdFileInfo, nil
}

func (s *FileInfoService) GetFileInfos() ([]FileInfo, error) {
	fileInfos, _ := s.repo.FindAll()

	return fileInfos, nil
}

func (s *FileInfoService) GetFileInfoByID(id uint) (FileInfo, error) {
	fileInfo, _ := s.repo.FindByID(id)

	return fileInfo, nil
}

func (s *FileInfoService) RemoveFileInfoByID(id uint) (FileInfo, error) {
	fileInfo, _ := s.GetFileInfoByID(id)
	deletedAt := time.Now()
	fileInfo.DeletedAt = &deletedAt
	deletedFileInfo, _ := s.repo.Update(fileInfo)

	return deletedFileInfo, nil
}

func (s *FileInfoService) DeleteFileInfoByID(id uint) (FileInfo, error) {
	fileInfo, _ := s.GetFileInfoByID(id)
	deletedAt := time.Now()
	fileInfo.DeletedAt = &deletedAt
	deletedFileInfo, _ := s.repo.Delete(fileInfo)

	return deletedFileInfo, nil
}
