// db_test.go
package postgres

import (
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize("postgres", "SailIn1985!", "fatsoma")

}
