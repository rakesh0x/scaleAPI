package main
import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"sync"
)

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strconv"
// 	"sync"
// )

// mutex for security to shared state
var counter int
var mutex = &sync.Mutex{}

// func echoString(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "hello")
// }

// func incrementCounter(w http.ResponseWriter, r *http.Request) {
// 	mutex.Lock()
// 	counter++
// 	fmt.Fprint(w, strconv.Itoa(counter))
// 	mutex.Unlock()
// }

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the landing Page!")

	fmt.Println("Endpoint hit: Homepage")
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " Welcome to Dashboard!")
	mutex.Lock()
	fmt.Println("Endpoint hit: Dashboard")
	mutex.Unlock()
}

// func handleRequest() {
// 	http.HandleFunc("/homepage", homepage)
// 	http.HandleFunc("/dashboard", dashboard)
// 	http.HandleFunc("/increment", incrementCounter)
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, r.URL.Path[1:])
// 	})
// 	http.HandleFunc("/article", returnAllArticle)
// 	log.Fatal(http.ListenAndServe(":3000", nil))
// }

// func main() {
// 	handleRequest()
// }

func handleRequests() {
	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/", homepage)
	myrouter.HandleFunc("/article", returnAllArticle)

	log.Fatal(http.ListenAndServe(":3000", myrouter))
}

func main() {
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"}, 
		Article{Title: "Hello 1", Desc: "Article Description 1", Content: "Article Content 1"},
	}
	handleRequests()
}