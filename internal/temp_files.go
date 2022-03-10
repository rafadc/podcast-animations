package internal

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func PrepareTempPaths() {
	tmpPath := filepath.Join(".", "tmp")
	err := os.MkdirAll(tmpPath, os.ModePerm)
	if err != nil {
		log.Fatalf("error creating temp directory: %v", err)
	}

	outputPath := filepath.Join(".", "output")
	err = os.MkdirAll(outputPath, os.ModePerm)
	if err != nil {
		log.Fatalf("error creating output directory: %v", err)
	}

}

func CleanupTempFiles() {
	err := os.RemoveAll("tmp")
	if err != nil {
		log.Fatalf("error removing temp directory: %v", err)
	}
}
