package sql

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/DATA-DOG/godog"
)

type dbFeature struct {
	TestDb    *sql.DB
	TestError error
	ConnInfo  *ConnectionInfo
}

func (db *dbFeature) iConnect() error {
	conn, err :=
		openDatabase(db.ConnInfo)
	err = conn.Ping()

	db.TestDb = conn
	db.TestError = err

	return nil

}

func (db *dbFeature) iProvideServerOfWithPort(server string, port int) error {
	db.ConnInfo = &ConnectionInfo{}
	db.ConnInfo.server = server
	db.ConnInfo.port = port
	db.ConnInfo.debug = true

	return nil
}

func (db *dbFeature) iShouldNotBeConnected() error {

	if db.TestError == nil {
		return fmt.Errorf("expected to not be connected but I am")
	}

	return nil
}

func (db *dbFeature) errorShouldSay(errorMessageExpected string) error {
	if !strings.Contains(db.TestError.Error(), errorMessageExpected) {
		return fmt.Errorf("expected %s, got %s", errorMessageExpected, db.TestError.Error())
	}
	return nil
}
func (db *dbFeature) iShouldBeConnected() error {
	if db.TestDb != nil {
		return fmt.Errorf("expected to be connected but I am not: %s", db.TestError.Error())
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	db := &dbFeature{}
	s.Step(`^I connect$`, db.iConnect)
	s.Step(`^I should be connected$`, db.iShouldBeConnected)
	s.Step(`^I provide server of "([^"]*)" with port \'(\d+)\'$`, db.iProvideServerOfWithPort)
	s.Step(`^I should not be connected$`, db.iShouldNotBeConnected)
	s.Step(`^I should see error "([^"]*)"`, db.errorShouldSay)

	s.BeforeScenario(func(interface{}) {

	})
}

func TestMain(m *testing.M) {
	format := "progress"
	for _, arg := range os.Args[1:] {
		if arg == "-test.v=true" { // go test transforms -v option
			format = "pretty"
			break
		}
	}
	status := godog.RunWithOptions("godog", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format: format,
		Paths:  []string{"features"},
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
