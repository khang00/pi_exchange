package io

import (
	"encoding/csv"
	"os"
	"pi_exchange/core"
)

type CSVReader struct {
	objects []core.DataObject
	cursor  int
}

func NewCSVReader(file string) (CSVReader, error) {
	fileReader, err := os.Open(file)
	if err != nil {
		return CSVReader{}, err
	}

	lines, err := csv.NewReader(fileReader).ReadAll()
	if err != nil {
		return CSVReader{}, err
	}

	if len(lines) == 0 {
		return CSVReader{}, err
	}

	objects := make([]core.DataObject, 0)
	header := lines[0]
	for _, line := range lines[1:] {
		object := make(core.DataObject, 0)
		for i, cell := range line {
			object[header[i]] = cell
		}

		objects = append(objects, object)
	}

	return CSVReader{
		objects: objects,
		cursor:  0,
	}, nil
}

func (r CSVReader) GetNextObject() (core.DataObject, error) {
	currObj := r.objects[r.cursor]
	r.cursor += 1

	return currObj, nil
}

func (r CSVReader) IsEmpty() bool {
	return r.cursor == len(r.objects)
}
