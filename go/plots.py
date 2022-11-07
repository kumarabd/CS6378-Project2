import matplotlib.pyplot as plt

# ********** Metrics with d_mean **********

# Graph 1: (Response Time Vs Mean Inter Request Delay)

# Inter-Req Delay Values (FIXED)
d = list(range(0,12,2))
# print(d)

# Input all the values for the response_time
response_time = [1.4, 0.98, 0.47, 0.95, 0.89, 0.8]
f1 = plt.figure(1)
plt.plot(d, response_time, label="lamport", color='green', linewidth = 3, marker='o', markerfacecolor='blue', markersize=12)
plt.title("Response Time Vs Mean Inter Request Delay")
plt.xlabel('Mean Inter-Request Delay (ms)')
plt.ylabel('Response Time (UNIT)')
plt.legend("Lamport")

# plt.savefig("RT Vs IR.png")
# plt.ylim([0, 1.5])


# Graph 2: (System Throughput Vs Mean Inter Request Delay)

# Inter-Req Delay Values (FIXED)
d = list(range(0,12,2))
# print(d)

# Input all the values for system_throughput
system_throughput = [20,3,5,6,7,8]
f2 = plt.figure(2)
plt.plot(d, system_throughput, label="lamport", color='red', linewidth = 3, marker='d', markerfacecolor='blue', markersize=12)
plt.title("System Throughput Vs Mean Inter Request Delay")
plt.xlabel('Mean Inter-Request Delay (ms)')
plt.ylabel('System Throughput (UNIT)')
plt.legend("Lamport")

# plt.savefig("ST Vs IR.png")
# plt.ylim([0, 1.5])


# ********** Metrics with cs_mean **********

# Graph 3: (Response Time Vs Mean CS Time)

# Mean CS Time Values (FIXED)
cs = list(range(0,12,2))
# print(cs)

# Input all the values for the response_time
response_time = [1.4, 0.98, 0.47, 0.95, 0.89, 0.8]
f3 = plt.figure(3)
plt.plot(cs, response_time, label="lamport", color='green', linewidth = 3, marker='o', markerfacecolor='blue', markersize=12)
plt.title("Response Time Vs Mean CS Time")
plt.xlabel('Mean CS Time (ms)')
plt.ylabel('Response Time (UNIT)')
plt.legend("Lamport")

# plt.savefig("RT Vs CS.png")
# plt.ylim([0, 1.5])


# Graph 4: (System Throughput Vs Mean CS Time)

# Mean CS Time Values (FIXED)
cs = list(range(0,12,2))
# print(cs)

# Input all the values for system_throughput
system_throughput = [20,3,5,6,7,8]
f4 = plt.figure(4)

plt.plot(cs, system_throughput, color='red', linewidth = 3, marker='d', markerfacecolor='blue', markersize=12)
plt.title("System Throughput Vs Mean CS Time")
plt.xlabel('Mean CS Time (ms)')
plt.ylabel('System Throughput (UNIT)')
plt.legend("Lamport")

plt.show()
# plt.savefig("ST Vs CS.png")
# plt.ylim([0, 1.5])