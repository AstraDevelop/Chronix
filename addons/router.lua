local function handle_request(req, res)
    local path = req.url
    if path == "/" then
        res:send("Welcome to Chronix Server!")
    elseif path == "/about" then
        res:send("This is a simple router plugin")
    else
        res:status(404):send("Page Not Found")
    end
end

return handle_request
