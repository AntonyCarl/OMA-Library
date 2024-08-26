package repository

import (
	"io"
	"log"
	"mime/multipart"
	"os"
)

const path = "C:/GoLearn/FileRep/"

func SaveFile(file multipart.File, fileName string) string {
	dst, err := os.Create(path + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	io.Copy(dst, file)
	return path + fileName
}
