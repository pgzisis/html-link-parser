package htmlparser

import (
	"log"
	"os"
	"reflect"
	"testing"
)

// TestEx1 calls Parse with the file ex1.html, checking for a valid return value
func TestEx1(t *testing.T) {
	want := []Link{{Href: "/other-page", Text: "A link to another page"}}

	file, err := os.Open("../ex1.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	result := Parse(file)

	if !reflect.DeepEqual(want, result) {
		t.Fatalf("Result: %v is not equal to wanted: %v", result, want)
	}
}

// TestEx2 calls Parse with the file ex2.html, checking for a valid return value
func TestEx2(t *testing.T) {
	want := []Link{
		{Href: "https://www.twitter.com/joncalhoun", Text: "Check me out on twitter"},
		{Href: "https://github.com/gophercises", Text: "Gophercises is on"},
	}

	file, err := os.Open("../ex2.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	result := Parse(file)

	if !reflect.DeepEqual(want, result) {
		t.Fatalf("Result: %v is not equal to wanted: %v", result, want)
	}
}

// TestEx3 calls Parse with the file ex3.html, checking for a valid return value
func TestEx3(t *testing.T) {
	want := []Link{
		{Href: "#", Text: "Login"},
		{Href: "/lost", Text: "Lost? Need help?"},
		{Href: "https://twitter.com/marcusolsson", Text: "@marcusolsson"},
	}

	file, err := os.Open("../ex3.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	result := Parse(file)

	if !reflect.DeepEqual(want, result) {
		t.Fatalf("Result: %v is not equal to wanted: %v", result, want)
	}
}

// TestEx4 calls Parse with the file ex4.html, checking for a valid return value
func TestEx4(t *testing.T) {
	want := []Link{{Href: "/dog-cat", Text: "dog cat"}}

	file, err := os.Open("../ex4.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	result := Parse(file)

	if !reflect.DeepEqual(want, result) {
		t.Fatalf("Result: %v is not equal to wanted: %v", result, want)
	}
}
