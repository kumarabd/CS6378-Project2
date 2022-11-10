import matplotlib.pyplot as plt

# ********** Metrics with d_mean **********

# Graph 1: (Response Time Vs Mean Inter Request Delay)

# Inter-Req Delay Values (FIXED)
d = list(range(0,12,2))
# print(d)

# Input all the values for the response_time
response_time_lamport = [12.23, 10.84, 11.34, 10.95, 10.39, 11.11]
response_time_ricart = [11.4, 12.98, 10.47, 10.95, 10.89, 10.81]
f1 = plt.figure(1)
plt.plot(d, response_time_lamport, label="lamport", color='green', linewidth = 3, marker='o', markerfacecolor='blue', markersize=12)
plt.plot(d, response_time_ricart, label="ricart", color='red', linewidth = 2, marker='*', markerfacecolor='black', markersize=12)
plt.title("Response Time Vs Mean Inter Request Delay")
plt.xlabel('Mean Inter-Request Delay (ms)')
plt.ylabel('Response Time (ms)')
plt.legend(["Lamport", "Ricart"])


plt.savefig("RT Vs IR.png")
# plt.ylim([0, 1.5])


# Graph 2: (System Throughput Vs Mean Inter Request Delay)

# Inter-Req Delay Values (FIXED)
d = list(range(0,12,2))
# print(d)

# Input all the values for system_throughput
system_throughput_lamport = [0.034,0.0358,0.0341,0.0357,0.034,0.0346]
system_throughput_ricart = [0.0341,0.0328,0.0356,0.037,0.0331,0.0353]
f2 = plt.figure(2)
plt.plot(d, system_throughput_lamport, label="lamport", color='red', linewidth = 3, marker='d', markerfacecolor='blue', markersize=12)
plt.plot(d, system_throughput_ricart, label="ricart", color='black', linewidth = 2, marker='o', markerfacecolor='green', markersize=12)
plt.title("System Throughput Vs Mean Inter Request Delay")
plt.xlabel('Mean Inter-Request Delay (ms)')
plt.ylabel('System Throughput')
plt.legend(["Lamport", "Ricart"])


plt.savefig("ST Vs IR.png")
# plt.ylim([0, 1.5])


# ********** Metrics with cs_mean **********

# Graph 3: (Response Time Vs Mean CS Time)

# Mean CS Time Values (FIXED)
cs = list(range(0,12,2))
# print(cs)

# Input all the values for the response_time
response_time_lamport = [1.4, 0.98, 0.47, 0.95, 0.89, 0.8]
response_time_ricart = [21.4, 13.98, 20.47, 20.95, 20.89, 30.8]

f3 = plt.figure(3)
plt.plot(d, response_time_lamport, label="lamport", color='green', linewidth = 3, marker='o', markerfacecolor='blue', markersize=12)
plt.plot(d, response_time_ricart, label="ricart", color='red', linewidth = 2, marker='*', markerfacecolor='black', markersize=12)
plt.title("Response Time Vs Mean CS Time")
plt.xlabel('Mean CS Time (ms)')
plt.ylabel('Response Time (ms)')
plt.legend(["Lamport", "Ricart"])


plt.savefig("RT Vs CS.png")
# plt.ylim([0, 1.5])


# Graph 4: (System Throughput Vs Mean CS Time)

# Mean CS Time Values (FIXED)
cs = list(range(0,12,2))
# print(cs)

# Input all the values for system_throughput
system_throughput_lamport = [0.034,0.0358,0.0341,0.0357,0.034,0.0346]
system_throughput_ricart = [0.0341,0.0328,0.0356,0.037,0.0331,0.0353]
f4 = plt.figure(4)

plt.plot(d, system_throughput_lamport, label="lamport", color='red', linewidth = 3, marker='d', markerfacecolor='blue', markersize=12)
plt.plot(d, system_throughput_ricart, label="ricart", color='black', linewidth = 2, marker='o', markerfacecolor='green', markersize=12)
plt.title("System Throughput Vs Mean CS Time")
plt.xlabel('Mean CS Time (ms)')
plt.ylabel('System Throughput')
plt.legend(["Lamport", "Ricart"])

plt.savefig("ST Vs CS.png")
plt.show()
# plt.ylim([0, 1.5])