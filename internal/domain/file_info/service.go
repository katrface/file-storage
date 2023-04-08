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
	createdFileInfo, err := s.repo.Create(fileInfo)
	if err != nil {
		return fileInfo, err
	}

	return createdFileInfo, nil
}

func (s *FileInfoService) GetFileInfos() ([]FileInfo, error) {
	fileInfos, err := s.repo.FindAll()
	if err != nil {
		return fileInfos, err
	}

	return fileInfos, nil
}

func (s *FileInfoService) GetFileInfoByID(id uint) (FileInfo, error) {
	fileInfo, err := s.repo.FindByID(id)
	if err != nil {
		return fileInfo, err
	}

	return fileInfo, nil
}

func (s *FileInfoService) RemoveFileInfoByID(id uint) (FileInfo, error) {
	fileInfo, err := s.GetFileInfoByID(id)
	if err != nil {
		return fileInfo, err
	}

	deletedAt := time.Now()
	fileInfo.DeletedAt = &deletedAt

	deletedFileInfo, err := s.repo.Update(fileInfo)
	if err != nil {
		return fileInfo, err
	}

	return deletedFileInfo, nil
}

func (s *FileInfoService) DeleteFileInfoByID(id uint) (FileInfo, error) {
	fileInfo, err := s.GetFileInfoByID(id)
	if err != nil {
		return fileInfo, err
	}

	deletedAt := time.Now()
	fileInfo.DeletedAt = &deletedAt

	deletedFileInfo, err := s.repo.Delete(fileInfo)
	if err != nil {
		return fileInfo, err
	}

	return deletedFileInfo, nil
}
