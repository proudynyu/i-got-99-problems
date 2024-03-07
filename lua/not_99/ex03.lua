local io = require("io")

local function sum(x, y)
    return x + y
end

print("We will sum two numbers for you!")
print("Please, digit the first one: ")
local x = io.read("*n")
print("Now, digit the second one: ")
local y = io.read("*n")

print("Them sum of the two number is " .. sum(x, y))
