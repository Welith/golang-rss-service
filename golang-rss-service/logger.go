package main

import "os"

func LogFileExists(name string) bool {

	if _, err := os.Stat(name); err != nil {

		if os.IsNotExist(err) {

			return false
		}
	}

	return true
}

func CreateLogFile(name string) error {

	fo, err := os.Create(name)

	if err != nil {

		return err
	}

	defer func() {

		fo.Close()
	}()

	return nil
}

