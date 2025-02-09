package versions_test

import (
	"embed"
	"net/http"
	"testing"

	itbasisTestUtilsFiles "github.com/itbasis/go-test-utils/v5/files"
	"github.com/itbasis/go-test-utils/v5/ginkgo"
	"github.com/onsi/gomega/ghttp"
	"golang.org/x/tools/godoc/vfs"
)

//go:embed testdata/*
var testHTMLContents embed.FS

func TestVersionSuite(t *testing.T) {
	ginkgo.InitGinkgoSuite(t, "Golang Versions Suite")
}

func initFakeServer(testResponseFile string) *ghttp.Server {
	var server = ghttp.NewServer()

	server.AppendHandlers(
		ghttp.CombineHandlers(
			ghttp.VerifyRequest("GET", "/"),
			ghttp.RespondWith(http.StatusOK, itbasisTestUtilsFiles.ReadFile(vfs.FromFS(testHTMLContents), testResponseFile)),
		),
	)

	return server
}
