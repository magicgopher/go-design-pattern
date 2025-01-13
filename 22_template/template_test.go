package template_test

import (
	template "github.com/magicgopher/go-design-pattern/22_template"
	"testing"
)

func TestTemplateMethod(t *testing.T) {
	// 制作一辆轿车
	sedan := &template.SedanCar{}
	sedanMaker := &template.CarTemplate{Car: sedan}
	t.Log("Producing Sedan:")
	sedanMaker.MakeCar()

	// 制作一辆SUV
	suv := &template.SUVCar{}
	suvMaker := &template.CarTemplate{Car: suv}
	t.Log("Producing SUV:")
	suvMaker.MakeCar()
}
