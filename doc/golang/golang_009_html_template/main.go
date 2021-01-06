package main

import (
	"github.com/Masterminds/sprig"
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {
	r := gin.Default()
	r.GET("/hello", hello)
	r.GET("/ifelse", ifelse)
	r.GET("/range", rangef)
	r.GET("/compare", compare)
	r.GET("/callmethod", callmethod)
	r.GET("/callfunction", callfunction)
	r.GET("/usesprig", usesprig)
	r.GET("/nest", nestedTemplate)
	_ = r.Run(":9090")
}

func hello(c *gin.Context) {
	t := template.Must(template.ParseFiles("tmpl/hello.gohtml"))
	_ = t.Execute(c.Writer, struct {
		Title string
	}{"2021"})
}

func ifelse(c *gin.Context) {
	t := template.Must(template.ParseFiles("tmpl/ifelse.gohtml"))
	_ = t.Execute(c.Writer, struct {
		Name string
	}{"John"})
}

type Item struct {
	Name  string
	Price float32
}

func rangef(c *gin.Context) {
	t := template.Must(template.ParseFiles("tmpl/range.gohtml"))
	_ = t.Execute(c.Writer, struct {
		Items []Item
	}{
		[]Item{
			{"IPhone", 100.242},
			{"Xiaomi", 95.141},
			{"Samsung", 99.889},
			{"Huawei", 98.122},
		},
	})
}

func compare(c *gin.Context) {
	t := template.Must(template.ParseFiles("tmpl/compare.gohtml"))
	_ = t.Execute(c.Writer, struct {
		Name string
		Age  int
	}{"John", 16})
}

type User struct {
	Id   int64
	Name string
}

// 只能在 gohtml 中調用接收者爲結構體的方法，不能調用接收者爲指針類型的方法
func (u User) IsPaySuccess(payId string) bool {
	return u.Id != 0 && payId != ""
}

func callmethod(c *gin.Context) {
	t := template.Must(template.ParseFiles("tmpl/callmet.gohtml"))
	err := t.Execute(c.Writer, User{
		Id:   1234,
		Name: "John",
	})
	if err != nil {
		println(err)
	}
}

func callfunction(c *gin.Context) {
	// 注意：Funcs 必須在 ParseFiles 之前，否則在 ParseFiles 時將會拋出函數未定義錯誤。
	t, err := template.New("cf").Funcs(template.FuncMap{
		"isOK": func(payId string) bool {
			return payId != ""
		},
	}).ParseFiles("tmpl/callfunc.gohtml")
	if err != nil {
		println(err)
	}
	err = t.Execute(c.Writer, struct{}{})
	if err != nil {
		println(err)
	}
}

func usesprig(c *gin.Context) {
	t, err := template.New("us").Funcs(sprig.FuncMap()).ParseFiles("tmpl/usesprig.gohtml")
	if err != nil {
		println(err)
	}
	err = t.Execute(c.Writer, struct{}{})
	if err != nil {
		println(err)
	}
}

func nestedTemplate(c *gin.Context) {
	t, err := template.New("base").Funcs(sprig.FuncMap()).ParseFiles("tmpl/base.gohtml", "tmpl/layout.gohtml", "tmpl/footer.gohtml")
	if err != nil {
		println(err)
	}
	err = t.Execute(c.Writer, struct {
		Layout struct {
			Content string
		}
		Footer struct {
			CopyRight string
		}
	}{
		Layout: struct{ Content string }{Content: "Hey U"},
		Footer: struct{ CopyRight string }{CopyRight: "Me All Right Reserved"},
	})
	if err != nil {
		println(err)
	}
}
