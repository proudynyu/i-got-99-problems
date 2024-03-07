local io = require("io")

--- @param grades number[]
--- @return number
local function get_media(grades)
    local sum = 0

    for i = 1, #grades, 1 do
        sum = sum + grades[i]
    end

    local average = sum / 4

    return average
end

local function main()
    print("Insert the four grades to get the average grade\n")
    local grades = {}

    for i = 1, 4, 1 do
        print("Grade " .. i)
        grades[i] = io.read("*n")
    end

    local average = get_media(grades)

    print("The average is " .. average)
end

main()

