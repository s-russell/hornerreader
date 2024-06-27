package horner

type Book struct {
	Name        string
	VerseCounts []int
}

func (b Book) Chapters() int {
	return len(b.VerseCounts)
}

type Category struct {
	Books []Book
}

type CategoryReading struct {
	BookName string `json:"bookName"`
	Chapter  int    `json:"chapter"`
}

func (c Category) GetReading(readingNum int) CategoryReading {
	readings := make([]CategoryReading, 0, 200)

	for _, book := range c.Books {
		for i := 0; i < book.Chapters(); i++ {
			readings = append(readings, CategoryReading{book.Name, i + 1})
		}
	}

	return readings[(readingNum-1)%len(readings)]
}
