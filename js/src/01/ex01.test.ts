import {describe, expect, it} from "bun:test"
import { getLast } from "."

describe("getLast suite", () => {
    it("should return the undefined if empty array is passed", () => {
        const c: unknown[] = []
        const response = getLast(c)

        expect(response).toBeUndefined()
    })

    it("should return the first element if is the unique element on the array", () => {
        const c: string[] = ["a"]
        const r = getLast<string>(c)

        expect(r).toBe("a")
    })

    it("should return the last element if there is more then one element on the array", () => {
        const c: string[] = ["a", "b", "c"]
        const r = getLast<string>(c)

        expect(r).toBe("c")
    })
})

