package positive

import (
	"se-test/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestNegaitive2(t *testing.T) {

	t.Run("Check Negative2", func(t *testing.T) {
		g := gomega.NewGomegaWithT(t)
		employ := entity.Employees{
			Name:         "boss",
			Salary:       17000,
			EmployeeCode: "-",
		}
		ok, err := govalidator.ValidateStruct(employ)

		g.Expect(ok).To(gomega.BeFalse())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by - and 4 digits (0-9)"))
	})

}
