package service

import (
	"archive/zip"
	"bytes"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"time"

	"github.com/ymoldabe/Doodocs-Backend-Challenge/internal/models"
)

type ArchiveType struct{}

func NewArchive() *ArchiveType {
	return &ArchiveType{}
}

func (a *ArchiveType) ExtractArhiveInfo(file *multipart.FileHeader) (models.ArhiveInfo, error) {
	var zipTotalSize float64

	zipFile, err := file.Open()
	if err != nil {
		return models.ArhiveInfo{}, err
	}

	defer zipFile.Close()

	zipReader, err := zip.NewReader(zipFile, file.Size)
	if err != nil {
		return models.ArhiveInfo{}, err
	}

	var files models.ArhiveInfo

	files.FileName = file.Filename
	files.ArchiveSize = float64(file.Size)
	files.TotalFiles = float64(len(zipReader.File))

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

func (a *ArchiveType) CreateArchive(files []*multipart.FileHeader) (*bytes.Buffer, int, error) {

	zipFile := new(bytes.Buffer)
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range files {
		switch file.Header.Get("Content-Type") {
		case "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
			"application/xml",
			"image/jpeg",
			"image/png":

			f, err := file.Open()

			if err != nil {
				return nil, http.StatusInternalServerError, err
			}
			defer f.Close()

			filename := file.Filename

			h := &zip.FileHeader{Name: filename, Method: zip.Deflate, Flags: 0x800, Modified: time.Now()}

			zipFileInArchive, err := zipWriter.CreateHeader(h)
			if err != nil {
				return nil, http.StatusInternalServerError, err
			}

			_, err = io.Copy(zipFileInArchive, f)
			if err != nil {
				return nil, http.StatusInternalServerError, err
			}
		default:
			return nil, http.StatusBadRequest, nil
		}
	}

	err := zipWriter.Close()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return zipFile, http.StatusOK, nil
}
