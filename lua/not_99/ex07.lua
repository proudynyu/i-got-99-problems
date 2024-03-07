local function get_square_area(size)
    return size * size
end

local function get_double_area(size)
    return get_square_area(size) * 2
end

print("Insert a side measure for a square to receive its area and double area.")
print("Square size: ")
local size = io.read("*n")

print("Square with side size: " .. size)
print("Has area: " .. get_square_area(size))
print("And the double area is: " .. get_double_area(size))
