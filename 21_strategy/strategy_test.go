package strategy_test

import (
	strategy "github.com/magicgopher/go-design-pattern/21_strategy"
	"testing"
)

// TestPaymentStrategies 策略模式单元测试
func TestPaymentStrategies(t *testing.T) {
	context := &strategy.PaymentContext{}

	// 使用信用卡支付
	creditCard := &strategy.CreditCardStrategy{CardNumber: "1234-5678-9876-5432"}
	context.SetStrategy(creditCard)
	result := context.Pay(100.0)
	expected := "Paid 100.00 using Credit Card (1234-5678-9876-5432)"
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}

	// 使用 PayPal 支付
	payPal := &strategy.PayPalStrategy{Email: "user@example.com"}
	context.SetStrategy(payPal)
	result = context.Pay(200.0)
	expected = "Paid 200.00 using PayPal (user@example.com)"
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}

	// 使用现金支付
	cash := &strategy.CashStrategy{}
	context.SetStrategy(cash)
	result = context.Pay(50.0)
	expected = "Paid 50.00 using Cash"
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}

	// 没有设置支付策略
	context = &strategy.PaymentContext{}
	result = context.Pay(30.0)
	expected = "No payment strategy set"
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
