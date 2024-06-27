function addProgressEventListeners(localStorageKey, firstLoad) {
    const boxes = [...document.getElementsByClassName("readingProgress")]
    for (let i = 0; i < boxes.length; i++) {
        const box = boxes[i]
        if (firstLoad) {
            box.checked = localStorage.getItem(`${localStorageKey}:${i}`) === "true"
        } else {
            localStorage.setItem(`${localStorageKey}:${i}`, false)
        }
        box.addEventListener("change", ({target: {checked}}) => {
            localStorage.setItem(`${localStorageKey}:${i}`, box.checked)
        })
    }
}

(function () {
    const localStorageKey = "hornerReaderCurrentReading"
    let firstLoad = true

    const boxes = [...document.getElementsByClassName("readingProgress")]
    for (let i = 0; i < boxes.length; i++) {
        const box = boxes[i]
        box.checked = localStorage.getItem(`${localStorageKey}:${i}`) === "true"
    }

    document.body.addEventListener("newReading", ({detail: {value}}) => {
        localStorage.setItem(localStorageKey, value)
        addProgressEventListeners(localStorageKey, firstLoad)
        firstLoad = false
    })

    const currentReading = parseInt(localStorage.getItem(localStorageKey)) || 1

    htmx.ajax('GET', `/reading/${currentReading}`, '#reading')
})()