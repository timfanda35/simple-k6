package main

import (
	"embed"
	"github.com/creack/pty"
	"github.com/olahol/melody"
	"io/fs"
	"net/http"
	"os/exec"
)

//go:embed ui
var content embed.FS

func getFileSystem() http.FileSystem {
	webDir, err := fs.Sub(content, "ui")
	if err != nil {
		panic(err)
	}

	return http.FS(webDir)
}

func main() {
	cmd := exec.Command("sh")
	tty, err := pty.Start(cmd)
	if err != nil {
		panic(err)
	}

	m := melody.New()

	go func() {
		for {
			buf := make([]byte, 1024)
			read, err := tty.Read(buf)
			if err != nil {
				return
			}

			m.Broadcast(buf[:read])
		}
	}()

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		tty.Write(msg)
	})

	http.HandleFunc("/webterminal", func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})

	webDir := http.FileServer(getFileSystem())
	http.Handle("/", http.StripPrefix("/", webDir))

	http.ListenAndServe("0.0.0.0:8080", nil)
}
