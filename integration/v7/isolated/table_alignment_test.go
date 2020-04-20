package isolated

import (
	"code.cloudfoundry.org/cli/integration/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("table alignment", func() {
	BeforeEach(func() {
		helpers.LoginCF()
	})

	When("output is in English", func() {
		BeforeEach(func() {
			helpers.LoginCF()
			helpers.TargetOrgAndSpace(ReadOnlyOrg, ReadOnlySpace)
		})

		// Developer note: The spacing in this test is significant and explicit. Do
		// not replace with a regex.
		It("aligns the table correctly", func() {
			username, _ := helpers.GetCredentials()
			session := helpers.CF("target")
			Eventually(session).Should(Say("API endpoint:   %s", apiURL))
			Eventually(session).Should(Say(`API version:    [\d.]+`))
			Eventually(session).Should(Say("user:           %s", username))
			Eventually(session).Should(Say("org:            %s", ReadOnlyOrg))
			Eventually(session).Should(Say("space:          %s", ReadOnlySpace))
			Eventually(session).Should(Exit(0))
		})
	})

	When("output is in language with multibyte characters", func() {
		BeforeEach(func() {
			helpers.LoginCF()
			helpers.TargetOrgAndSpace(ReadOnlyOrg, ReadOnlySpace)
		})

		// Developer note: The spacing in this test is significant and explicit. Do
		// not replace with a regex.
		It("aligns the table correctly", func() {
			username, _ := helpers.GetCredentials()
			session := helpers.CFWithEnv(map[string]string{"LANG": "ja-JP.utf8"}, "target")
			Eventually(session).Should(Say("API エンドポイント:   %s", apiURL))
			// TODO: "version" here should be translated for all languages. We have translation resources for "api version"
			// (lowercase), which is what this said in V6, but we don't yet have them for "API version" (uppercase).
			Eventually(session).Should(Say(`API version:      [\d.]+`))
			Eventually(session).Should(Say("ユーザー:            %s", username))
			Eventually(session).Should(Say("組織:                %s", ReadOnlyOrg))
			Eventually(session).Should(Say("スペース:            %s", ReadOnlySpace))
			Eventually(session).Should(Exit(0))
		})
	})
})
