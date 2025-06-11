package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](parameter T) string {
	return parameter.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	name string
}
func (m *MyManager) GetName() string {
	return m.name
}
func (m *MyManager) GetManagerName() string {
	return "Manager: " + m.name
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}
type MyVicePresident struct {
	name string
}
func (vp *MyVicePresident) GetName() string {
	return vp.name
}
func (vp *MyVicePresident) GetVicePresidentName() string {
	return "Vice President: " + vp.name
}

func TestEmployee(t *testing.T) {
	var empName string = GetName[Employee](&MyManager{name: "John Doe"})
	assert.Equal(t, "John Doe", empName)

	var managerName string = GetName[Manager](&MyManager{name: "Jane Smith"})
	assert.Equal(t, "Jane Smith", managerName)

	var vpName string = GetName[VicePresident](&MyVicePresident{name: "Alice Johnson"})
	assert.Equal(t, "Alice Johnson", vpName)
}