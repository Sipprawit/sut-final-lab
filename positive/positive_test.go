package positive

import (
	"se-test/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestPositive(t *testing.T) {

	t.Run("Check Positive", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		employ := entity.Employees{
			Name:         "boss",
			Salary:       17000,
			EmployeeCode: "BB-4450",
		}
		ok, err := govalidator.ValidateStruct(employ)

		g.Expect(ok).To(gomega.BeTrue())
		g.Expect(err).To(gomega.BeNil())
	})
}
