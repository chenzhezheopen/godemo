package main
import(
	"text/template"
)
func main(){
	var templateText = `This is a test for template:
Name: {{.Name}}
Age: {{.Age}}
School: {{.School}}
Married: {{.MarriedOK}}
`
var t, err = t.Parse(templateText)
}