package decoprator
//https://blog.csdn.net/m0_38132420/article/details/78326846

import "fmt"

type Company interface {
	Show()
}

type BaseCompany struct {
}

func (c BaseCompany) Show() {
	fmt.Println("公司有老板、销售")
}

type SonCompany struct {
	Company
}

func (c SonCompany) Show() {
	fmt.Println("子公司有：")
	c.Company.Show()
	c.AddWorker()
}

func (c SonCompany) AddWorker() {
	fmt.Println("还有程序员")
}

type ParentCompany struct {
	Company
}

func (c ParentCompany) Show() {
	fmt.Println("母公司有：")
	c.Company.Show()
	c.AddWorker()
}

func (c ParentCompany) AddWorker() {
	fmt.Println("除此之外，各个职能人员应有尽有。。。")

}
