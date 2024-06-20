import {FC} from "react";
import {HornerReading} from "../../api/horner-reading.ts";


const Reading: FC<{ reading: HornerReading}> = ({reading }) => {
    const rows = [
        "Acts",
        "Epistles A",
        "Epistles B",
        "Gospels",
        "Pentateuch",
        "Prophets",
        "Proverbs",
        "Psalms",
        "Wisdom",
        "Writings",
    ]
    return <table><tbody>

        {rows.map( row =>
            <tr key={row}>
                <td >{row}</td>
                <td>{reading[row].bookName} {reading[row].chapter}</td> //@ts-ignore TS7053
            </tr>
        )}
    </tbody>
    </table>
}

export default Reading