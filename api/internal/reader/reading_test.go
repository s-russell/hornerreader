package reader

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Chapters(t *testing.T) {
	got := Genesis.Chapters()
	if got != 50 {
		t.Error("wanted: 50, got:", got)
	}
}

func Test_GetReading(t *testing.T) {
	cat := Category{
		[]Book{Jonah, Nahum},
	}

	data := []struct {
		readingNum int
		reading    CategoryReading
	}{
		{1, CategoryReading{Jonah.Name, 1}},
		{2, CategoryReading{Jonah.Name, 2}},
		{3, CategoryReading{Jonah.Name, 3}},
		{4, CategoryReading{Jonah.Name, 4}},
		{5, CategoryReading{Nahum.Name, 1}},
		{6, CategoryReading{Nahum.Name, 2}},
		{7, CategoryReading{Nahum.Name, 3}},
		{8, CategoryReading{Jonah.Name, 1}},
		{9, CategoryReading{Jonah.Name, 2}},
		{10, CategoryReading{Jonah.Name, 3}},
	}

	for _, d := range data {
		t.Run(fmt.Sprintf("Reading: %d", d.readingNum), func(t *testing.T) {
			got := cat.GetReading(d.readingNum)

			if diff := cmp.Diff(d.reading, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
