#Node %d is entering CS at %lu
import re

record = {}
consistent = True
intervals=[]

print('✓ Reading log file')

# read output log
output_file = open('./build/log', 'r')

created = 0
running = 0
connected = 0
started = 0

counter = 30
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
                
    # if app != "" or time!= "" or msg!="":
    #     if counter:
    #         counter -=1
    #         print('Time: '+time+' App: '+app+' msg: '+msg)
    
    if 'created' in msg:
        created += 1
    if 'running' in msg:
        running += 1
    if 'started' in msg:    
        started += 1
    if 'connected' in msg:
        connected += 1
        
    if 'executing cs' in msg:
        # Node-1 --> [[3,7], [10,12]]
        clock_val = int(msg.split("clock ",1)[1])
        if app not in record.keys():
            record[app] = [[clock_val]]
        else:
            record[app].append([clock_val])    
    elif 'leaving cs' in msg:
        clock_val = int(msg.split("clock ",1)[1])
        record[app][-1].append(clock_val)

for val in record.values():
    for interval in val:
        intervals.append(interval)
    
output_config_file = open('./config.txt','r')
nodes = 0
num_cs_req = 0
for line in output_config_file.readlines():
    line = line.split('#')[0]
    if line == '' or line[0] < '0' or line[0] > '9':
        continue
    else:
        nodes = int(line.split()[0])
        num_cs_req = int(line.split()[3])
        break

if nodes == len(record):
    print("✓ Number of nodes matches config")
else:
    print("✕ Number of nodes does not match config")
    
if created == nodes:
    print("✓ All nodes created")
else:
    print("✕ Number of created nodes does not match config")
    
if running == nodes:
    print("✓ All nodes are running")
else:
    print("✕ Number of nodes running does not match config")    
    
if connected == nodes:
    print("✓ All nodes are connected")
else:
    print("✕ Number of connected nodes does not match config")
num_cs_flag = True
for k in record.keys():
    if len(record[k]) != num_cs_req:
        num_cs_flag = False
        
if num_cs_flag:
    print("✓ Number of cs requests per node matches config")
else:
    print("✕ Number of cs requests per node does not match config")
    
intervals.sort()
consistent = True
for i in range(1, len(intervals)):
    if intervals[i][0] < intervals[i-1][1]:
        consistent = False
        
if not consistent:
    print("✕ The algorithm is not consistent")
else:
    print("✓ The algorithm is consistent")