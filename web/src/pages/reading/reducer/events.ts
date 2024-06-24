export type RequestedReading = {
    name: "requestedReading", readingNumber: number
}

export type RecievedReading = {
    name: "recievedReading", readingNumber: number
}

export type ReadingCategory =
    | "Pentateuch"
    | "Writings"
    | "Psalms"
    | "Proverbs"
    | "Wisdom"
    | "Prophets"
    | "Gospels"
    | "Acts"
    | "Epistles A"
    | "Epistles B"

export type ToggleCategory = {
    name: 'toggleCategory', category: ReadingCategory
}

export type SentUpdate = {
    name: 'sentUpdate',
    updateId: string
}

export type UpdateAcknowledged = {
    name: 'updateAcknowledged',
    updateId: string
}

export type ReaderEvent =
    | RequestedReading
    | RecievedReading
    | ToggleCategory
    | SentUpdate
    | UpdateAcknowledged