package tests

import (
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"os"
	"path/filepath"
	"testing"
)

func TestSuiteRunner(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		panic("не удалось получить рабочую директорию: " + err.Error())
	}

	// Поднимаемся на уровень выше — в корень проекта
	projectRoot := filepath.Dir(cwd)

	// Устанавливаем путь для Allure
	_ = os.Setenv("ALLURE_OUTPUT_PATH", projectRoot)
	_ = os.Setenv("ALLURE_OUTPUT_FOLDER", "allure-results")

	suite.RunSuite(t, new(TestSuiteOne))
	suite.RunSuite(t, new(TestSuiteTwo))
}
