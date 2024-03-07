local function convert_celsius_to_fahrenheit(celsius)
    return (celsius * (9 / 5)) + 32
end

print("Enter a celsius number to be converted to fahrenheit: ")
local celsius = io.read("*n")
print("Celsius: " .. celsius)
print("Fahrenheit: " .. convert_celsius_to_fahrenheit(celsius))
