import _ "net/http/pprof"

func main() {
	// if don't have http server yet, then start like
	go func() { log.Println(http.ListenAndServe("localhost:6060", nil)) }()
}