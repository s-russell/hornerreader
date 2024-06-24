import {expect, test, describe} from "vitest";

describe('arithemetic', () => {

    test('addition', () => {
        expect(2+2).equals(4, "2 and 2 should be 4")
    })

    test('multiplication', () => {
        expect(2*2).equals(4, "2 by 2 should be 4")
    })
})