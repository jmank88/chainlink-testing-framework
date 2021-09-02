package contracts

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/smartcontractkit/integrations-framework/environment"
	"github.com/smartcontractkit/integrations-framework/suite/testcommon"
)

var _ = Describe("OCR Feed @ocr", func() {

	DescribeTable("Deploys and watches an OCR feed @ocr", func(
		envInit environment.K8sChainlinkGroupsInit,
		numNodes int,
	) {
		i := &testcommon.OCRSetupInputs{}
		testcommon.DeployOCRForEnv(i, envInit, numNodes)
		testcommon.SetupOCRTest(i)
		testcommon.CheckRound(i)
		By("Tearing down the environment", i.SuiteSetup.TearDown())
	},
		Entry("all the same version", environment.NewChainlinkNodesGroups, 5),
		Entry("different versions", environment.NewMixedVersionChainlinkGroupInit(2), 5),
	)
})
