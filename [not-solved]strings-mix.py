def mix(s1, s2):
    s1_count = {}
    s2_count = {}
    
    for c in s1:
        if not c.islower():
            continue
        s1_count[c] = s1_count.get(c, 0) + 1
        
    for c in s2:
        if not c.islower():
            continue
        s2_count[c] = s2_count.get(c, 0) + 1
    
    
    s1_count = dict(sorted(s1_count.items(), key=lambda item: (-item[1], item[0])))
    s2_count = dict(sorted(s2_count.items(), key=lambda item: (-item[1], item[0])))
    
    res = ""
    
    print(s1_count)
    print(s2_count)
    
    for key, value in s1_count.items():
        if not key in s2_count:
            if value == 1:
                continue
            res += "/1:" + key * value
        else:
            if value == 1:
                del s2_count[key]
                continue
                
            if value < s2_count[key]:
                res += "/2:" + key * s2_count[key]
            elif value > s2_count[key]:
                res += "/1:" + key * value
            else:
                res += "/=:" + key * value
                
            del s2_count[key]
            
    print(s2_count)
    
    if len(s2_count) != 0:
        for key, value in s2_count.items():
            if value == 1:
                continue
            
            res += "/2:" + key * value

    return res[1:]
