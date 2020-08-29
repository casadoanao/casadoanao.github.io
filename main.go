package main

import (
	"fmt"
	"log"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("$PORT não definido.")
	}
	fmt.Println("I'm alive, bitch!")
	go func() {
		for {
			fmt.Printf("\r")
			fmt.Printf("eae, não morre, não!")
			time.Sleep(10 * time.Minute)
		}
	}()

	http.Handle("/", http.FileServer(http.Dir(".")))
	//http.HandleFunc("/index/", indexFunc)
	http.HandleFunc("/pages/", pagesFunc)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

type File struct {
	Name string
	Body []byte
}

/*
func defaultFunc(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/index", http.StatusFound)
}
*/

/*
func indexFunc(w http.ResponseWriter, r *http.Request) {
	pagesDir, err := ioutil.ReadDir("pages")
	if err != nil {
		log.Fatal(err)
		return
	}

	originalPage, err := load("index")
	if err != nil {
		log.Fatal(err)
		return
	}

	changedPage := string(originalPage.Body)

	for _, page := range pagesDir {
		pageToLoad := strings.Replace(page.Name(), ".html", "", -1)
		if pageToLoad != "notfound" {
			changedPage += fmt.Sprintf("\n<p><a href=\"/pages/%s\" target=\"_parent\"> %s </a></p>\n", pageToLoad,
			pageToLoad)
		}
	}

	fmt.Fprintf(w, "<body>%s</body>", []byte(changedPage))
}
*/

func pagesFunc(w http.ResponseWriter, r *http.Request) {
	pageTitle := r.URL.Path[len("/pages/"):]
	t, _ := template.ParseFiles("pages.html")

	if pageTitle != "" {
		page, err := load("pages/"+pageTitle)
		if err != nil {
			http.Redirect(w, r, "/pages/notfound", http.StatusFound)
			return
		}

		fmt.Fprintf(w, "%s", page.Body)
	} else {
		pagesDir, err := ioutil.ReadDir("pages")
		if err != nil {
			log.Fatal(err)
			return
		}

		var changingPage string

		for _, page := range pagesDir {
			pageToLoad := strings.Replace(page.Name(), ".html", "", -1)
			if pageToLoad != "notfound" {
				changingPage += fmt.Sprintf("<li><a href=\"/pages/%s\"> %s </a></li>", pageToLoad,
				strings.Title(pageToLoad))
			}
		}

		// toAddToTemplate := File{
		// 	Body: []byte(changingPage),
		// }
		t.Execute(w, template.HTML(changingPage))
	}

	//fmt.Fprintf(w, "<html>%s</html>", finalPage)
}


		
func load(title string) (File, error) {
	filename := title + ".html"

	fileBody, err := ioutil.ReadFile(filename)
	if err !=  nil {
		return File{}, err
	}

	completeFile := File{
		filename,
		fileBody,
	}

	return completeFile, err
}

func apnd(slice *[]interface{}, value interface{}) {
	*slice = append(*slice, value)
}
