type ReadingItem = {
    bookName: string;
    chapter: number;
}

type HornerReading = {
    "Acts": ReadingItem,
    "Epistles A": ReadingItem,
    "Epistles B": ReadingItem,
    "Gospels": ReadingItem,
    "Pentateuch": ReadingItem,
    "Prophets": ReadingItem,
    "Proverbs": ReadingItem,
    "Psalms": ReadingItem,
    "Wisdom": ReadingItem,
    "Writings": ReadingItem
}

const firstReading = {
    "Acts": {
        "bookName": "Acts",
        "chapter": 1
    },
    "Epistles A": {
        "bookName": "Romans",
        "chapter": 1
    },
    "Epistles B": {
        "bookName": "I Thessalonians",
        "chapter": 1
    },
    "Gospels": {
        "bookName": "Matthew",
        "chapter": 1
    },
    "Pentateuch": {
        "bookName": "Genesis",
        "chapter": 1
    },
    "Prophets": {
        "bookName": "Isaiah",
        "chapter": 1
    },
    "Proverbs": {
        "bookName": "Proverbs",
        "chapter": 1
    },
    "Psalms": {
        "bookName": "Psalms",
        "chapter": 1
    },
    "Wisdom": {
        "bookName": "Job",
        "chapter": 1
    },
    "Writings": {
        "bookName": "Joshua",
        "chapter": 1
    }
} as HornerReading

async function getHornerReading(readingNumber: number = 1) {

    const resp = await fetch(`/api/reader/reading/${readingNumber}`)
    if (resp.status != 200) {
        console.log(`Failed to retrieve reading ${readingNumber}: `, resp.statusText)
        return firstReading
    }

    return (await resp.json()) as HornerReading
}

export type { HornerReading, ReadingItem}
export { firstReading }
export default getHornerReading