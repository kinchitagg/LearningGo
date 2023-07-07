package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  string
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, []byte(p.Body), 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: string(body)}, nil
}

func userInput() (string, string) {
	var title string
	bodyReader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to your wiki")
	fmt.Println("Kindly enter your document title")
	fmt.Scanln(&title)
	fmt.Println("Enter your body")
	body, _ := bodyReader.ReadString('\n')
	return title, body
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	p, _ := loadPage(title)
	fmt.Println(p.Title, p.Body)
	fmt.Fprintf(w, "<h1>%s<h1><p>%s<p>", p.Title, p.Body)

}
func main() {

	title, body := userInput()
	pIn := &Page{
		Title: title,
		Body:  body,
	}
	pIn.save()
	//pOut, _ := loadPage(pIn.Title)

	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
