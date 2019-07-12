package main

import (
	"net/http"
	"text/template"
)

type temp struct {
	Name string
	Sex  string
}
type blo struct {
	man []temp
}

func (blo *blo) user(a string, b string) {
	blo.man = append(blo.man, temp{a, b})
}

func main() {

	http.HandleFunc("/a", blank)
	http.ListenAndServe(":8888", nil)

	// fmt.Println(tmpl)
}
func blank(w http.ResponseWriter, req *http.Request) {
	var num blo
	num.user("nihao", "chenzhe")
	num.user("nihao", "liujiang")
	num.user("nihao", "liujin")
	num.user("nihao", "dushuai")
	mubar := "{{range .}}{{.Name}}姓名{{.Sex}}\n{{end}}" //建立一个模板，内容是"hello, {{.}}"
	tmpl, err := template.New("test").Parse(mubar)     //建立一个模板
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, num.man) //将struct与模板合成，合成结果放到os.Stdout里
	if err != nil {
		panic(err)
	}

}
