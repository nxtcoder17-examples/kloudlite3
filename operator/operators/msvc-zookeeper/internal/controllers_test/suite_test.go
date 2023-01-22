package controllers_test

import (
	"testing"

	zookeeperMsvcv1 "github.com/kloudlite/operator/apis/zookeeper.msvc/v1"
	testlib "github.com/kloudlite/operator/testing"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAPIs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}

var schemes = testlib.AddToSchemes(zookeeperMsvcv1.AddToScheme)
var _ = testlib.PreSuite(schemes)

var _ = testlib.PostSuite()
