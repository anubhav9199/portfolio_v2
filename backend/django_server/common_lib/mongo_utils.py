def clean_data(func):
    def wrapping_function(*args, **kwargs):
        data = func(*args, **kwargs)
        result = []
        for r in data:
            if "_id" in r:
                _ = r.pop("_id")
            result.append(r)
        return result
    return wrapping_function
