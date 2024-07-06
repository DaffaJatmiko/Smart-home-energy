package util

import (
	"encoding/csv"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

func ReadCSV(filename string) (map[string][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return CsvToSlice(string(csvData))
}

func CsvToSlice(data string) (map[string][]string, error) {
	reader := csv.NewReader(strings.NewReader(data))
	headers, err := reader.Read()
	if err != nil {
		return nil, err
	}

	if len(headers) == 0 {
		return nil, errors.New("no headers found")
	}

	result := make(map[string][]string)
	for _, header := range headers {
		result[header] = []string{}
	}

	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		for i, value := range record {
			result[headers[i]] = append(result[headers[i]], value)
		}
	}

	return result, nil
}
