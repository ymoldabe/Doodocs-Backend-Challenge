package service

import (
	"archive/zip"
	"mime"
	"mime/multipart"
	"path/filepath"

	"github.com/ymoldabe/Doodocs-Backend-Challenge/internal/models"
)

type ArchiveT struct{}

func NewArchive() *ArchiveT {
	return &ArchiveT{}
}

func (a ArchiveT) ExtractArhiveInfo(file *multipart.FileHeader) (models.ArhiveInfo, error) {
	var zipTotalSize float64
	zipName := file.Filename
	zipFileSize := file.Size
	zipFile, err := file.Open()
	if err != nil {
		return models.ArhiveInfo{}, err
	}

	defer zipFile.Close()

	zipReader, err := zip.NewReader(zipFile, file.Size)
	if err != nil {
		return models.ArhiveInfo{}, err
	}
	zipTotalFile := len(zipReader.File)

	var files models.ArhiveInfo

	files.FileName = zipName
	files.ArchiveSize = float64(zipFileSize)
	files.TotalFiles = float64(zipTotalFile)

	for _, f := range zipReader.File {
		if !f.FileInfo().IsDir() {
			fileInfo := models.Files{}
			fileInfo.FilePath = f.Name
			fileInfo.Size = float64(f.UncompressedSize64)
			zipTotalSize += float64(f.UncompressedSize64)
			fileInfo.Mimetype = mime.TypeByExtension(filepath.Ext(f.Name))
			if err != nil {
				return models.ArhiveInfo{}, err
			}
			files.Files = append(files.Files, fileInfo)
		}
	}

	files.TotalSize = zipTotalSize

	return files, nil
}
