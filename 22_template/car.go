package template

import "fmt"

// 汽车

// SedanCar 表示一辆轿车
type SedanCar struct{}

// BuildChassis 轿车构建底盘
func (s *SedanCar) BuildChassis() {
	fmt.Println("Building chassis for Sedan...")
}

// InstallEngine 轿车安装引擎
func (s *SedanCar) InstallEngine() {
	fmt.Println("Installing engine for Sedan...")
}

// Paint 轿车喷漆
func (s *SedanCar) Paint() {
	fmt.Println("Painting Sedan with metallic silver...")
}

// AssembleBody 轿车组装车身
func (s *SedanCar) AssembleBody() {
	fmt.Println("Assembling Sedan body...")
}

// Complete 轿车完成制造
func (s *SedanCar) Complete() {
	fmt.Println("Sedan car is ready!")
}

// SUVCar 表示一辆SUV
type SUVCar struct{}

// BuildChassis SUV构建底盘
func (s *SUVCar) BuildChassis() {
	fmt.Println("Building chassis for SUV...")
}

// InstallEngine SUV安装引擎
func (s *SUVCar) InstallEngine() {
	fmt.Println("Installing engine for SUV...")
}

// Paint SUV喷漆
func (s *SUVCar) Paint() {
	fmt.Println("Painting SUV with matte black...")
}

// AssembleBody SUV组装车身
func (s *SUVCar) AssembleBody() {
	fmt.Println("Assembling SUV body...")
}

// Complete SUV完成制造
func (s *SUVCar) Complete() {
	fmt.Println("SUV car is ready!")
}
