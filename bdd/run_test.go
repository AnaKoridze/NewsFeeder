package tests

import (
	"github.com/cucumber/godog"
	"github.com/gin-gonic/gin"
	"os"
	"testing"
)

// This function is used for setup before executing the test functions
func TestMain(m *testing.M) {

	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Set directory to root
	os.Chdir("..")

	// Configure godog
	format := "progress"
	for _, arg := range os.Args[1:] {
		if arg == "-test.v=true" { // go test transforms -v option
			format = "pretty"
			break
		}
	}
	status := godog.RunWithOptions("godog", func(s *godog.Suite) {
		FeatureContext(s)
		godog.SuiteContext(s)
	}, godog.Options{
		Format: format,
		Paths:  []string{"bdd/features"},
	})

	// Run tests
	if st := m.Run(); st > status {
		status = st
	}

	// Exit with appropriate code
	os.Exit(status)

}
