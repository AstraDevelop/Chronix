local function log_request(req)
    local time = os.date("%Y-%m-%d %H:%M:%S")
    local method = req.method
    local url = req.url
    print(string.format("[%s] %s %s", time, method, url))
end

return log_request
