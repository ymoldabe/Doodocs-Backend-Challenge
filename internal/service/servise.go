package service

import (
	"mime/multipart"

	"github.com/ymoldabe/Doodocs-Backend-Challenge/internal/models"
)

type Archive interface {
	ExtractArhiveInfo(file *multipart.FileHeader) (models.ArhiveInfo, error)
}

type Service struct {
	Archive
}

func New() *Service {
	return &Service{
		Archive: NewArchive(),
	}
}
