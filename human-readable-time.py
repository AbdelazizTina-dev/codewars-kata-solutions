def make_readable(seconds):
    ss = seconds % 60
    seconds -= ss
    hh = seconds // 3600
    seconds -= hh * 3600
    mm = seconds // 60

    return "%02d:%02d:%02d" % (hh, mm, ss)
