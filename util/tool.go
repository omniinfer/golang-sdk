package util

import (
	"encoding/base64"
	"os"
)

func ReadImage(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func ReadImageToBase64(name string) (string, error) {
	imageBytes, err := ReadImage(name)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(imageBytes), nil
}
