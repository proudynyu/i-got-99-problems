local function calculate_salary(hour_value, qty_hours)
    local days_work = 5 * 4 -- five day, 4 weeks

    return (days_work * qty_hours) * hour_value
end

print("How much do you receive by hour: ")
local hour_value = io.read("*n")

print("How much hours do you work: ")
local qty_hours = io.read("*n")

local salary = calculate_salary(hour_value, qty_hours)
print("Your salary is: " .. salary)
