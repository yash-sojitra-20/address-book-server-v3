package utils

import (
	"address-book-server-v3/internal/common/fault"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/samber/mo"
)

type FileDetails struct {
	FilePath string
	FileName string
}

func GenerateCustomAddressesCSV(userID uint64, fields []string, data []map[string]interface{}) mo.Result[*FileDetails] {
	timestamp := time.Now().Format("20060102_150405")

	baseDir, err := os.Getwd()
	if err != nil {
		return mo.Err[*FileDetails](fault.InternalServerError(err))
	}

	fileName := fmt.Sprintf(
		"address_custom_%d_%s.csv",
		userID,
		timestamp,
	)

	dir := filepath.Join(baseDir, "exports")
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return mo.Err[*FileDetails](fault.InternalServerError(err))
	}

	filePath := filepath.Join(dir, fileName)

	// Create file
	file, err := os.Create(filePath)
	if err != nil {
		return mo.Err[*FileDetails](fault.InternalServerError(err))
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	if err := writer.Write(fields); err != nil {
		return mo.Err[*FileDetails](fault.InternalServerError(err))
	}

	// Write rows
	for _, row := range data {
		record := make([]string, len(fields))

		for i, field := range fields {
			if val, ok := row[field]; ok && val != nil {
				record[i] = fmt.Sprint(val)
			} else {
				record[i] = ""
			}
		}

		if err := writer.Write(record); err != nil {
			return mo.Err[*FileDetails](fault.InternalServerError(err))
		}
	}

	return mo.Ok(&FileDetails{
		FilePath: filePath,
		FileName: fileName,
	})
}
