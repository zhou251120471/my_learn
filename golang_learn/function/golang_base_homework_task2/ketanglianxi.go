package golang_base_homework_task2

import "fmt"

type PayMethod interface {
	Account
	Pay(account int) bool
}
type Account interface {
	GetBalance() int
}
type CreditCard struct {
	Balance int
	Limit   int
}

func (c *CreditCard) Pay(account int) bool {
	if c.Balance+account >= c.Limit {
		c.Balance += account
		fmt.Println("信用卡支付成功：%d\n", account)
		return true
	}
	fmt.Println("信用卡支付失败：超出额度")
	return false
}
func (c *CreditCard) GetBalance() int {
	return c.Balance
}

type DebitCard struct {
	Balance int
}

func (d *DebitCard) Pay(amount int) bool {
	if d.Balance >= amount {
		d.Balance -= amount
		fmt.Printf("借记卡支付成功: %d\n", amount)
		return true
	}
	fmt.Println("借记卡支付失败: 余额不足")
	return false
}

func (d *DebitCard) GetBalance() int {
	return d.Balance
}

// 使用 PaymentMethod 接口的函数
func PurchaseItem(p PayMethod, price int) {
	if p.Pay(price) {
		fmt.Printf("购买成功, 剩余余额: %d\n", p.GetBalance())
	} else {
		fmt.Println("购买失败")
	}
}
