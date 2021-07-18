package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Tag struct {
	Name string `json:"name,omitempty"`
}

type Category struct {
	Name string `json:"name,omitempty"`
}

func main() {

	http.HandleFunc("/tags", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got requests for fetching tags list")
		tags, err := getTagsJson()
		if err != nil {
			log.Println("Error in tagsJson")
		}
		w.Write(tags)
	})

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got requests for fetching categories list")
		categories, err := getCategoriesJson()
		if err != nil {
			log.Println("Error in getCategoriesJson")
		}

		w.Write(categories)
	})

	// listen to port
	http.ListenAndServe(":5050", nil)

}

func getTagsJson() ([]byte, error) {
	tags := fetchTags()
	tags_json, err := json.Marshal(tags)

	if err != nil {
		return nil, err
	}
	return tags_json, nil
}

func getCategoriesJson() ([]byte, error) {
	tags := fetchCategories()
	tags_json, err := json.Marshal(tags)

	if err != nil {
		return nil, err
	}
	return tags_json, nil
}

func fetchTags() []Tag {
	url := "http://localhost:1313/tags.html"
	log.Printf("Fetching tags from %s \n", url)
	completionData := fetchCompletionData(url)
	var tags []Tag
	for _, d := range completionData {
		tag := Tag{Name: strings.TrimSpace(d)}
		tags = append(tags, tag)
	}

	return tags
}

func fetchCategories() []Category {
	url := "http://localhost:1313/categories.html"
	log.Printf("Fetching categories from %s \n", url)
	completionData := fetchCompletionData(url)
	var categories []Category
	for _, d := range completionData {
		category := Category{Name: strings.TrimSpace(d)}
		categories = append(categories, category)
	}
	return categories
}

func fetchCompletionData(url string) []string {
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	v := strings.Split(string(html), ",")
	return v
}
