package product_test

import (
	"fmt"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMyPackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Product Service test Suite")
}

var _ = BeforeSuite(func() {
	fmt.Println("BeforeSuite: This runs before any specs in the suite")
})

var _ = AfterSuite(func() {
	fmt.Println("AfterSuite: This runs after all specs in the suite")
})

var _ = Describe("MyPackage", func() {
	var (
		someResource string
	)

	BeforeEach(func() {
		fmt.Println("BeforeEach: This runs before each spec")
		someResource = "initialized"
	})

	AfterEach(func() {
		fmt.Println("AfterEach: This runs after each spec")
	})

	Context("When doing something", func() {
		It("should do the first thing", func() {
			fmt.Println("Spec: It should do the first thing")
			Expect(someResource).To(Equal("initialized"))
		})

		It("should do the second thing", func() {
			fmt.Println("Spec: It should do the second thing")
			Expect(someResource).To(Equal("initialized"))
		})
	})

	Context("When doing something else", func() {
		It("should do another thing", func() {
			fmt.Println("Spec: It should do another thing")
			Expect(someResource).To(Equal("initialized"))
		})
	})
})

func TestMain(m *testing.M) {
	fmt.Println("TestMain: This runs before any tests in the package")
	code := m.Run()
	fmt.Println("TestMain: This runs after all tests in the package")
	os.Exit(code)
}
