package file_info

import "time"

type FileInfoRepository interface {
	FindAll() ([]FileInfo, error)
	FindById(id uint) (FileInfo, error)
	Create(file_info FileInfo) (FileInfo, error)
	Update(file_info FileInfo) (FileInfo, error)
	Delete(file_info FileInfo) (FileInfo, error)
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

func (s *FileInfoService) GetFileInfoById(id uint) (FileInfo, error) {
	fileInfo, _ := s.repo.FindById(id)

	return fileInfo, nil
}

func (s *FileInfoService) RemoveFileInfoById(id uint) (FileInfo, error) {
	fileInfo, _ := s.GetFileInfoById(id)
	deletedAt := time.Now()
	fileInfo.DeletedAt = &deletedAt
	deletedFileInfo, _ := s.repo.Update(fileInfo)
	return deletedFileInfo, nil
}

func (s *FileInfoService) DeleteFileInfoById(id uint) (FileInfo, error) {
	fileInfo, _ := s.GetFileInfoById(id)
	deletedAt := time.Now()
	fileInfo.DeletedAt = &deletedAt
	deletedFileInfo, _ := s.repo.Delete(fileInfo)
	return deletedFileInfo, nil
}
