package repository

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ideatocode/go/config/pkg/interfaces"
)

// File .
type File struct {
	Extension string
}

// Load .
func (r File) Load(cfg interfaces.Config) error {

	file, err := os.Open(r.Path(cfg))
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	err = cfg.Marshaler().Unmarshal(byteValue, cfg)

	if err != nil {
		return fmt.Errorf("File Unmarshal Error: %s", err)
	}

	return err
}

// Save .
func (r File) Save(cfg interfaces.Config) error {
	cfg.Logger().Log("File: Flushing changes to disk")
	cfg.Lock()
	defer cfg.Unlock()
	b, err := cfg.Marshaler().Marshal(cfg)
	if err != nil {
		return fmt.Errorf("Json Marshal Error: %s", err)
	}
	err = ioutil.WriteFile(r.Path(cfg), b, 0644)

	if err != nil {
		return fmt.Errorf("Failed to write e: %s, p: %s", err, r.Path(cfg))
	}

	return nil
}

// Path .
func (r File) Path(cfg interfaces.Config) string {
	return cfg.Path() + r.Extension
}
