local cache = {}

local function get_from_cache(req)
    local key = req.url
    if cache[key] then
        return cache[key]
    end
    return nil
end

local function set_cache(req, res)
    local key = req.url
    cache[key] = res.body
end

return {
    get_from_cache = get_from_cache,
    set_cache = set_cache
}
