package positive

import (
	"testing"
	"se-test/entity"
	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestNegaitive1(t *testing.T) {

	t.Run("Check Negative1", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		employ := entity.Employees{
			Name: "boss",
			Salary: 10000,
			EmployeeCode: "BB4450",
		}
		ok ,err := govalidator.ValidateStruct(employ)

		g.Expect(ok).To(gomega.BeFalse())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("Salary must be between 15000 and 200000"))
	})

}
