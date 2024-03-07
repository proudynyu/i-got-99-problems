local function calculate_circle_area(radius)
    return string.format("%.2f",math.pi * math.pow(radius, 2))
end

print("Insert a radius for a circle: ")
local radius = io.read("*n")
print("The area for the circle with radius " .. radius .. " is " .. calculate_circle_area(radius))
