package service

import (
	"bytes"
	"mime/multipart"

	"github.com/ymoldabe/Doodocs-Backend-Challenge/configs"
	"github.com/ymoldabe/Doodocs-Backend-Challenge/internal/models"
)

type Archive interface {
	ExtractArhiveInfo(file *multipart.FileHeader) (models.ArhiveInfo, error)
	CreateArchive(files []*multipart.FileHeader) (*bytes.Buffer, int, error)
}

type Mail interface {
	SendLetters(file *multipart.FileHeader, emails string) error
}

type Service struct {
	Archive
	Mail
	Config configs.Config
}

func New(cnf configs.Config) *Service {
	return &Service{
		Archive: NewArchive(),
		Mail:    NewMail(cnf),
		Config:  cnf,
	}
}
