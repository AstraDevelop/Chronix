local function serve_static_files(req, res)
    local file_path = "public" .. req.url
    local file = io.open(file_path, "r")
    if file then
        local content = file:read("*a")
        file:close()
        res:send(content)
    else
        res:status(404):send("File Not Found")
    end
end

return serve_static_files
