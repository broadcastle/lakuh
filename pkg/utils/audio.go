package utils

import (
	"errors"
	"os"

	filetype "gopkg.in/h2non/filetype.v1"
)

// AudioCheck checks if the src is a valid audio file.
func AudioCheck(src string) (string, error) {

	file, err := os.Open(src)
	if err != nil {
		return "", err
	}

	defer file.Close()

	header := make([]byte, 261)
	if _, err := file.Read(header); err != nil {
		return "", err
	}

	check := make([]byte, 261)
	copy(check, header)

	if !filetype.IsAudio(header) {
		return "", errors.New("not a audio file")
	}

	kind, err := filetype.Match(check)
	if err != nil {
		return "", err
	}

	mime := kind.MIME.Value

	if mime == "audio/midi" || mime == "audio/m4a" || mime == "audio/amr" {
		return "", errors.New("unplayable audio file")
	}

	return mime, nil

}
