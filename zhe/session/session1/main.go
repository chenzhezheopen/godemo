package main //Sess SessionMgr实例
import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var Sess *SessionMgr
var Logger *log.Logger

func init() {
	Sess = NewSessionMgr("2222", 0)
}
func main() {
	logfile, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_APPEND, 0)
	Logger = log.New(logfile, "\r\n", log.Ldate|log.Lshortfile)
	server := &http.Server{
		Addr: ":2323",
	}
	http.HandleFunc("/regular", reg)
	http.HandleFunc("/change", cha)
	server.ListenAndServe()
}
func reg(w http.ResponseWriter, r *http.Request) {
	sessid := Sess.StartSession(w, r) //生成
	Sess.SetSessionVal(sessid, "usrname", "usrname")
	fmt.Println(sessid)
	w.Write([]byte("ooooooooooooooooooooo"))
}

func cha(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.RequestURI)
	fmt.Println(r.Cookie("2222"))
	sessid := Sess.CheckCookieValid(w, r)
	agentID, _ := Sess.GetSessionVal(sessid, "usrname") //获取
	fmt.Println(agentID)
}
