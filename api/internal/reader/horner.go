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

func GetHornerReadingPlan(readingNum int) map[string]CategoryReading {
	var readingsByCategory = make(map[string]CategoryReading, len(hornerReadingPlan))
	for _, item := range hornerReadingPlan {
		readingsByCategory[item.Name] = item.GetReading(readingNum)
	}
	return readingsByCategory
}
