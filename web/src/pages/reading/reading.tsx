import {FC} from "react";
import {ReadingItem} from "../../api/horner-reading.ts";


const Reading: FC<{ reading: Array<ReadingItem>}> = ({reading }) => {
    const rows = [
        "Pentateuch",
        "Writings",
        "Psalms",
        "Proverbs",
        "Wisdom",
        "Prophets",
        "Gospels",
        "Acts",
        "Epistles A",
        "Epistles B",
    ]
    const readingByCategory: { [key: string]: {book: string, chapter: number}} = {}
    reading.forEach(({category, book, chapter}) => readingByCategory[category] = {book, chapter})
    return <table><tbody>

        {rows.map( row =>
            <tr key={row}>
                <td >{row}</td>
                <td>{readingByCategory[row].book} {readingByCategory[row].chapter}</td>
            </tr>
        )}
    </tbody>
    </table>
}

export default Reading