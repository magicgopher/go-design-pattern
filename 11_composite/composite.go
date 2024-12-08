package composite

import "fmt"

// Component 接口定义了所有参与组合的对象的公共方法
type Component interface {
	GetName() string // 获取名称
	ShowDetails()    // 显示详细信息
}

// Employee 员工
type Employee struct {
	name  string // 姓名
	title string // 职位
}

func NewEmployee(name, title string) *Employee {
	return &Employee{
		name:  name,
		title: title,
	}
}

func (e *Employee) GetName() string {
	return e.name
}

func (e *Employee) ShowDetails() {
	fmt.Printf("Employee: %s, Title: %s\n", e.name, e.title)
}

// Department 即部门，包含多个员工或子部门
type Department struct {
	name       string      // 部门名称
	components []Component // 部门成员
}

func NewDepartment(name string) *Department {
	return &Department{
		name: name,
	}
}

func (d *Department) GetName() string {
	return d.name
}

func (d *Department) ShowDetails() {
	fmt.Printf("Department: %s\n", d.name)
	for _, component := range d.components {
		component.ShowDetails() // 递归显示成员
	}
}

func (d *Department) Add(component Component) {
	d.components = append(d.components, component)
}

func (d *Department) Remove(component Component) {
	for i, c := range d.components {
		if c == component {
			d.components = append(d.components[:i], d.components[i+1:]...)
			return
		}
	}
}

func (d *Department) Search(name string) Component {
	// 首先在当前部门中查找
	for _, component := range d.components {
		if component.GetName() == name {
			return component
		}
	}

	// 如果当前部门没有找到，再递归查找子部门
	for _, component := range d.components {
		if department, ok := component.(*Department); ok {
			// 如果是子部门，则递归调用Search方法
			found := department.Search(name)
			if found != nil {
				return found
			}
		}
	}

	// 如果没有找到，返回nil
	return nil
}

func (d *Department) GetSize() int {
	return len(d.components)
}
