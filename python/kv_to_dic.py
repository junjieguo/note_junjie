    msg_tmp_list = msg.split('&')
    msg_tmp_dist = {}
    for data in msg_tmp_list:
        if '=' not in data:
            continue
        key=data.split('=',1)[0]
        value=data.split('=',1)[1]
        msg_tmp_dist[key] = value
