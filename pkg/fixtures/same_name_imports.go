package test

import (
	"net/http"

	number_dir_http "github.com/vektra/mockery/v2/pkg/fixtures/12345678/http"
	my_http "github.com/vektra/mockery/v2/pkg/fixtures/http"
)

// Example is an example
type Example interface {
	A() http.Flusher
	B(fixtureshttp string) my_http.MyStruct
	C(fixtureshttp string) number_dir_http.MyStruct
}
