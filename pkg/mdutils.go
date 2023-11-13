package mdutils

import "os"

type MDFile struct {
	file *os.File
	Err  error
}

const MDExtension = ".md"

// New - Defines the New MDUtils object
func New(fileName string) *MDFile {
	file, err := os.Create(fileName + MDExtension)
	return &MDFile{file, err}
}

// Append - Append a new line to the string
func (md MDFile) Append(data, format string) error {
	var formattedData string
	switch format {
	case Bold:
		formattedData = Bold + data + Bold
	case Italics:
		formattedData = Italics + data + Italics
	default:
		formattedData = format + data
	}

	_, err := md.file.Write([]byte(formattedData + "\n"))
	if err != nil {
		return err
	}
	return nil
}

// Close: Closes the Mark-Down file
func (md *MDFile) Close() error {
	return md.file.Close()
}
