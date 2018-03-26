package hello_test

import (
	"fmt"
	"testing"

	gk "github.com/onsi/ginkgo"
	gm "github.com/onsi/gomega"

	service "github.com/sanksons/reflorest-testapp/src/common"
	"github.com/sanksons/reflorest-testapp/test_tools/fakers/webserver"
	reflorestservice "github.com/sanksons/reflorest/src/core/service"
)

var testHTTPServer *webserver.TestWebserver

func TestHelloAPI(t *testing.T) {
	gm.RegisterFailHandler(gk.Fail)
	gk.RunSpecs(t, "Hello API TEST Suite")
}

var _ = gk.BeforeSuite(func() {
	fmt.Println("Starting webserver")
	service.Register()
	//@todo: need to set init manager in reflorest so that its not needed to be set here explicitely.
	initMgr := new(reflorestservice.InitManager)
	initMgr.Execute()
	testHTTPServer = new(webserver.TestWebserver)
})

var _ = gk.AfterSuite(func() {
	fmt.Println("Crashing webserver")
	testHTTPServer = nil
})
