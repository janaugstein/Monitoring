package http

import (
	"net/http"
	//"fmt"
	//"path"
	//"path/filepath"
	//"github.com/gin-gonic/gin"
)

func ServeUI() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./ui/dist/ui/browser/index.html")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	 /* r:= gin.Default()
	 r.NoRoute(func(ctx *gin.Context) {
		dir, file := path.Split(ctx.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			ctx.File("./ui/dist/ui/browser/index.html")
		} else {
			ctx.File("./ui/dist/ui/browser" + path.Join(dir, file))
		}
	 })
	 fmt.Println("Listening on port 8080")
	 err := r.Run(":8080")
	 if err != nil {
		panic(err)
	 } */
}