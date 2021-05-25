package main
import (
	"github.com/FirdavsMF/http/cmd/app"
	"github.com/FirdavsMF/http/pkg/banners"
	"net"
	"net/http"
	"os"
	"fmt"
)

func main() {
	host := "0.0.0.0"
	port := "3939"

	if err := execute(host, port); err != nil {
		os.Exit(1)
	}
}

func execute(host string, port string) (err error) {
	mux := http.NewServeMux()
	bannersSvc := banners.NewService()
	server := app.NewServer(mux, bannersSvc)
	server.Init()

	srv := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: server,
	}

	fmt.Println("Server is listening...")

	return srv.ListenAndServe()
}