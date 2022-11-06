#Node %d is entering CS at %lu
import re

record = {}
consistent = True

# read output log
output_file = open('./build/log', 'r')
for line in output_file.readlines():
    time = ""
    level = ""
    msg = ""
    app = ""
    for subline in line.split(' '):
        word = subline.split('=')[0]
        if word == "time":
            time = subline.split('=')[1].strip('\n')
        elif word == "level":
            level = subline.split('=')[1].strip('\n')
        elif word == "app":
            app = subline.split('=')[1].strip('\n')
        elif word == "msg":
            msg = subline.split('=')[1].strip('\n')
        else:
            for m in subline.split('='):
                msg = msg+" "+m
    if 'executing cs' in msg:
        time = msg.strip('executing cs at ')
        if time in record.keys():
            consistent = False
        record[time] = app
    elif 'leaving cs' in msg:
        time = msg.strip('leaving cs at ')
        record.pop(time, None)
    #print(app)
    #print(msg)

if consistent:
    print("The algorithm is consistent")
else:
    print("The algorithm is not consistent")

# parse output to check if only one request was in CS at an instant