(function() {
    const prevButton = document.getElementById("prev")
    const nextButton = document.getElementById("next")
    const currentReadingInput = document.getElementById("currentReading")

    currentReadingInput.addEventListener("change", ({target: {value}}) => {
        prevButton.disabled = value == "1"
    })
})()