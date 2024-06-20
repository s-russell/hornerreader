import {useParams} from "wouter";
import {useEffect, useState} from "react";
import {firstReading} from "../../api/horner-reading.ts";
import getHornerReading from "../../api/horner-reading.ts";
import Reading from "./reading.tsx";

function App() {

  const {readingStr} = useParams()
  const [reading, setReading] = useState(firstReading)

  useEffect(() => {

    (async () => {
      const readingNumber = readingStr ? parseInt(readingStr) : 1
      const nextReading = await getHornerReading(readingNumber)
      setReading(nextReading)
    })()

  }, [readingStr])
  return <Reading reading={reading}/>
}

export default App
