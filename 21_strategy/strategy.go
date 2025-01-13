package strategy

import "fmt"

// PaymentStrategy 是所有支付策略的接口
type PaymentStrategy interface {
	Pay(amount float64) string
}

// CreditCardStrategy 使用信用卡支付的实现
type CreditCardStrategy struct {
	CardNumber string
}

func (c *CreditCardStrategy) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Credit Card (%s)", amount, c.CardNumber)
}

// PayPalStrategy 使用 PayPal 支付的实现
type PayPalStrategy struct {
	Email string
}

func (p *PayPalStrategy) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using PayPal (%s)", amount, p.Email)
}

// CashStrategy 使用现金支付的实现
type CashStrategy struct{}

func (c *CashStrategy) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Cash", amount)
}

// PaymentContext 是支付上下文，持有一个支付策略
type PaymentContext struct {
	strategy PaymentStrategy
}

// SetStrategy 设置支付策略
func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

// Pay 使用当前策略进行支付
func (p *PaymentContext) Pay(amount float64) string {
	if p.strategy == nil {
		return "No payment strategy set"
	}
	return p.strategy.Pay(amount)
}
