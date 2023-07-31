package model

var Books = []Book{
	{"1", "What is a Camel", "Robb Banks"},
	{"2", "Money, Unlimited Pleasure", "Howard Coward"},
	{"3", "Tell Me Why", "Agus Supriadi"},
}

type Book struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (book *Book) EmptyingAllField() {
	book.Id = ""
	book.Title = ""
	book.Author = ""
}
