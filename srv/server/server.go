package server

func InitServer() {
	r := NewRouter()
	r.Run(":80")
}
