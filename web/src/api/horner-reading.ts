type ReadingItem = {
    category: string;
    book: string;
    chapter: number;
}

const firstReading = [
    {
        "category": "Pentateuch",
        "book": "Genesis",
        "chapter": 1
    },
    {
        "category": "Writings",
        "book": "Joshua",
        "chapter": 1
    },
    {
        "category": "Psalms",
        "book": "Psalms",
        "chapter": 1
    },
    {
        "category": "Proverbs",
        "book": "Proverbs",
        "chapter": 1
    },
    {
        "category": "Wisdom",
        "book": "Job",
        "chapter": 1
    },
    {
        "category": "Prophets",
        "book": "Isaiah",
        "chapter": 1
    },
    {
        "category": "Gospels",
        "book": "Matthew",
        "chapter": 1
    },
    {
        "category": "Acts",
        "book": "Acts",
        "chapter": 1
    },
    {
        "category": "Epistles A",
        "book": "Romans",
        "chapter": 1
    },
    {
        "category": "Epistles B",
        "book": "I Thessalonians",
        "chapter": 1
    }
] as Array<ReadingItem>

async function getHornerReading(readingNumber: number = 1) {

    const resp = await fetch(`/api/reader/reading/${readingNumber}`)
    if (resp.status != 200) {
        console.log(`Failed to retrieve reading ${readingNumber}: `, resp.statusText)
        return firstReading
    }

    return (await resp.json()) as Array<ReadingItem>
}

export type { ReadingItem}
export { firstReading }
export default getHornerReading