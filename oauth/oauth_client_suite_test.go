package oauth_test

import (
	"github.com/SergeyDonskoy/wl"
	"github.com/SergeyDonskoy/wl/logger"
	"github.com/SergeyDonskoy/wl/oauth"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	"testing"
)

const (
	dummyAccessToken = "dummyAccessToken"
	dummyClientID    = "dummyClientID"
)

var (
	client wl.Client

	server *ghttp.Server
	apiURL string

	testLogger logger.Logger
)

func TestWL(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WL Suite")
}

var _ = BeforeEach(func() {
	server = ghttp.NewServer()
	apiURL = server.URL()

	testLogger = logger.NewTestLogger(GinkgoWriter)
	client = oauth.NewClient(
		dummyAccessToken,
		dummyClientID,
		apiURL,
		testLogger,
	)
})

var _ = AfterEach(func() {
	server.Close()
})
