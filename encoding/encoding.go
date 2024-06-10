package encoding

import (
	"encoding/json"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	// ниже реализуйте метод

	fileIn, err := os.ReadFile(j.FileInput)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(fileIn, &j.DockerCompose); err != nil {
		return err
	}

	fileOut, err := os.Create(j.FileOutput)
	if err != nil {
		return err
	}
	defer fileOut.Close()

	data, err := yaml.Marshal(j.DockerCompose)
	if err != nil {
		return err
	}
	_, err = fileOut.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод

	fileIn, err := os.ReadFile(y.FileInput)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(fileIn, &y.DockerCompose); err != nil {
		return err
	}

	fileOut, err := os.Create(y.FileOutput)
	if err != nil {
		return err
	}
	defer fileOut.Close()

	data, err := json.Marshal(y.DockerCompose)
	if err != nil {
		return err
	}
	_, err = fileOut.Write(data)
	if err != nil {
		return err
	}

	return nil
}
