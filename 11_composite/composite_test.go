package composite_test

import (
	composite "github.com/magicgopher/go-design-pattern/11_composite"
	"testing"
)

func TestCompositePattern(t *testing.T) {
	// 创建员工
	employee1 := composite.NewEmployee("John Doe", "Software Engineer")
	employee2 := composite.NewEmployee("Jane Smith", "Project Manager")
	employee3 := composite.NewEmployee("Michael Brown", "QA Engineer")

	// 创建部门并添加员工
	engineeringDept := composite.NewDepartment("Engineering")
	engineeringDept.Add(employee1)
	engineeringDept.Add(employee3)

	managementDept := composite.NewDepartment("Management")
	managementDept.Add(employee2)

	// 创建公司并将部门加入
	company := composite.NewDepartment("Company")
	company.Add(engineeringDept)
	company.Add(managementDept)

	// 验证公司结构是否正确
	t.Run("Test ShowDetails", func(t *testing.T) {
		// 捕获输出
		company.ShowDetails()
	})

	// 测试查找员工
	t.Run("Test Search Employee", func(t *testing.T) {
		found := company.Search("John Doe")
		if found == nil || found.GetName() != "John Doe" {
			t.Errorf("Expected 'John Doe', but got %v", found)
		}
	})

	// 测试删除员工
	t.Run("Test Remove Employee", func(t *testing.T) {
		engineeringDept.Remove(employee3)
		// 删除后，员工3不应再存在
		company.ShowDetails()
	})
}
