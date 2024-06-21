package reader

type HornerReadingPlanItem struct {
	Name string
	Category
}

var hornerReadingPlan = []HornerReadingPlanItem{
	{"Pentateuch", Category{[]Book{Genesis, Exodus, Leviticus, Numbers, Deuteronomy}}},
	{"Writings", Category{[]Book{Joshua, Judges, Ruth, I_Samuel, II_Samuel, I_Kings, II_Kings, I_Chronicles, II_Chronicles, Ezra, Nehemiah, Esther}}},
	{"Psalms", Category{[]Book{Psalms}}},
	{"Proverbs", Category{[]Book{Proverbs}}},
	{"Wisdom", Category{[]Book{Job, Ecclesiastes, SongOfSolomon}}},
	{"Prophets", Category{[]Book{Isaiah, Jeremiah, Lamentations, Ezekiel, Daniel, Hosea, Joel, Amos, Obadiah, Jonah, Micah, Nahum, Habakkuk, Zephaniah, Haggai, Zechariah, Malachi}}},
	{"Gospels", Category{[]Book{Matthew, Mark, Luke, John}}},
	{"Acts", Category{[]Book{Acts}}},
	{"Epistles A", Category{[]Book{Romans, I_Corinthians, II_Corinthians, Galatians, Ephesians, Philippians, Colossians, Hebrews}}},
	{"Epistles B", Category{[]Book{I_Thessalonians, II_Thessalonians, I_Timothy, II_Timothy, Titus, Philemon, James, I_Peter, II_Peter, I_John, II_John, III_John, Jude, Revelations}}},
}

type HornerReading struct {
	Name    string `json:"category"`
	Book    string `json:"book"`
	Chapter int    `json:"chapter"`
}

func GetHornerReadingPlan(readingNum int) []HornerReading {
	var readings = make([]HornerReading, len(hornerReadingPlan))
	for i, item := range hornerReadingPlan {
		categoryReading := item.GetReading(readingNum)
		readings[i] = HornerReading{
			item.Name,
			categoryReading.BookName,
			categoryReading.Chapter,
		}
	}
	return readings
}
