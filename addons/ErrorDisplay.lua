local function error_page(req, res)
    local status_code = req.status_code or 500
    local error_message = "An unexpected error occurred"
    
    if status_code == 404 then
        error_message = "Page not found"
    end

    res:status(status_code):send("<html><body><h1>" .. error_message .. "</h1></body></html>")
end

return error_page
