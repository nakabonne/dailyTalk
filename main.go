package main

import (
	"flag"
	"html/template"
	"log"
	"my_packages/trace"
	"net/http"
	"os"
	"sync"
)

// temp1は1つのテンプレートを表します
type templateHandler struct {
	once     sync.Once
	filename string
	temp1    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// []byte型のテンプレート情報を取得
	tpl, err := Asset("templates/chat.html")
	if err != nil {
		log.Fatal("ListenAndSearver:", err)
		os.Exit(1)
	}
	t.once.Do(func() {
		t.temp1 = template.Must(template.New("templates/" + t.filename).Parse(string(tpl)))
	})
	t.temp1.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse()
	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	// ルート
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	// チャットルーム開始
	go r.run()
	// webサーバーを開始
	log.Println("webサーバーを開始します。ポート：", *addr)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndSearver:", err)
	}
}
