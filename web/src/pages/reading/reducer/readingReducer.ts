import {
    ReaderEvent,
    ReadingCategory,
    RecievedReading,
    RequestedReading,
    SentUpdate,
    ToggleCategory,
    UpdateAcknowledged
} from "./events.ts";
import {ReadingItem} from "../../../api/horner-reading.ts";

export interface ReaderState {
    isLoading: boolean
    currentReading: number
    nextReading: number | null
    previousReading: number | null
    reading: ReadingItem | null
    progress: Array<{ category: ReadingCategory, complete: boolean }>
    unacknowledgedUpdates: Array<SentUpdate>
}


export default function (state: ReaderState, evt: ReaderEvent): ReaderState {
    let nextState = state
    switch (evt.name) {
        case 'recievedReading':
            nextState = recievedReading(state, evt as RecievedReading)
            break
        case "requestedReading":
            nextState = requestedReading(state, evt as RequestedReading)
            break
        case "sentUpdate":
            nextState = sentUpdate(state, evt as SentUpdate)
            break
        case "updateAcknowledged":
            nextState = updateAcknowledged(state, evt as UpdateAcknowledged)
            break
        case "toggleCategory":
            nextState = toggleCategory(state, evt as ToggleCategory)
            break
    }
    return nextState
}

function requestedReading(state: ReaderState, evt: RequestedReading) {
    return state;
}

function recievedReading(state: ReaderState, evt: RecievedReading) {
    return state;
}

function sentUpdate(state: ReaderState, evt: SentUpdate) {
    return state;
}

function updateAcknowledged(state: ReaderState, evt: UpdateAcknowledged): ReaderState {
    return state;
}

function toggleCategory(state: ReaderState, evt: ToggleCategory): ReaderState {
    return state
}