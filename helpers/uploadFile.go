package helpers

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

func UploadFile(filePath string, fileName string, extension string, file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// // Destination
	// split := strings.Split(file.Filename, ".")
	// filenameArr := split[:len(split)-1]
	// filename := strings.Join(filenameArr, "_")
	// extension := split[len(split)-1]
	fileName = fmt.Sprintf("%s_%s.%s", strings.ReplaceAll(fileName, " ", "_"), time.Now().Format("20060102150405"), extension)
	fileURL := fmt.Sprintf("public/%s/%s", filePath, fileName)
	dst, err := os.Create(fileURL)
	if err != nil {
		fmt.Println("disini")
		fmt.Println(err)
		return "", err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}
	return strings.ReplaceAll(fmt.Sprintf("%s/%s", os.Getenv("SERVER_URL"), fileURL), "public", "api/v1/uploads"), nil
}
