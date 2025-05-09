package writer

import (
	"encoding/json"
	"os"
	"path"
)

type Writer[T interface{}] struct {
	DirPath string
}

func NewWriter[T interface{}](dirPath string) (Writer[T], error) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return Writer[T]{}, err
	}

	return Writer[T]{DirPath: dirPath}, nil
}

func (w *Writer[T]) GetFilePath(filename string) string {
	return path.Join(w.DirPath, filename)
}

func (w *Writer[T]) Write(filename string, data []byte) error {
	p := w.GetFilePath(filename)
	return os.WriteFile(p, data, 0664)
}

func (w *Writer[T]) Read(filename string) ([]byte, error) {
	p := w.GetFilePath(filename)
	return os.ReadFile(p)
}

func (w *Writer[T]) JsonWrite(filename string, data T, indent bool) error {
	var bytes []byte
	var err error

	if indent {
		bytes, err = json.MarshalIndent(data, "", "	")
	} else {
		bytes, err = json.Marshal(data)
	}

	if err != nil {
		return err
	}

	return w.Write(filename, bytes)
}

func (w *Writer[T]) JsonRead(filename string) (T, error) {
	var out T
	bytes, err := w.Read(filename)
	if err != nil {
		return out, err
	}

	err = json.Unmarshal(bytes, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}
