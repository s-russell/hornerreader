(function() {

    const localStorageKey = "hornerReaderCurrentReading"

    const prevButton = document.getElementById("prev")
    const nextButton = document.getElementById("next")
    const currentReadingInput = document.getElementById("currentReading")

    currentReadingInput.addEventListener("change", ({target: {value}}) => {
        prevButton.disabled = value == "1"
        localStorage.setItem(localStorageKey, value)
    })

    nextButton.addEventListener("click", () => {
        const previousValue = parseInt(currentReadingInput.value) || 1
        currentReadingInput.value = previousValue + 1
        currentReadingInput.dispatchEvent(new Event('change'))
    })

    prevButton.addEventListener("click", () => {
        const previousValue = parseInt(currentReadingInput.value) || 1
        if(previousValue > 1) {
            currentReadingInput.value = Math.max(1, previousValue - 1)
            currentReadingInput.dispatchEvent(new Event('change'))
        }
    })

    const currentReading = parseInt(localStorage.getItem(localStorageKey)) || 1
    prevButton.disabled = currentReading == 1
    currentReadingInput.value = currentReading
})()