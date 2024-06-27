(function() {

    const localStorageKey = "hornerReaderCurrentReading"

    const currentReading = parseInt(localStorage.getItem(localStorageKey)) || 1

    htmx.ajax('GET', `/reading/${currentReading}`, '#reading')
})()