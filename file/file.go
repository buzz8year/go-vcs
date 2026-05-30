package file

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func Create(dir, filePath string) {
	file, _ := os.OpenFile(filepath.Join(dir, filePath), os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	file.Close()
}

func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func AppendLine(filePath, line string) {
	indexFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer indexFile.Close()

	_, err = fmt.Fprintln(indexFile, line)
	if err != nil {
		log.Fatal(err)
	}
}

func CopyTo(src, dir string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destination := filepath.Join(dir, filepath.Base(src))

	destinationFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}

func CalculateHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func IsModified(filePath, lastCommitHash string) bool {
	currentHash, err := CalculateHash(filePath)
	if err != nil {
		return false
	}
	return currentHash != lastCommitHash
}
