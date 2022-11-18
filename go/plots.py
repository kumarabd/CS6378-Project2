import matplotlib.pyplot as plt

# ********** Metrics with d_mean **********

# Graph 1: (Response Time Vs Mean Inter Request Delay)

# Inter-Req Delay Values (FIXED)
d = list(range(0,12,2))
# print(d)

# Input all the values for the response_time
response_time_lamport = [980.84, 981.98, 985.34, 1000.95, 1000.29, 1000.89]
response_time_ricart = [975.56, 976.88, 976.47, 978.57, 980.79, 984.99]
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
system_throughput_lamport = [0.0000529,0.0000512,0.0000510,0.000050,0.000049,0.000046]
system_throughput_ricart = [0.000050,0.000049,0.0000423,0.0000425,0.0000421,0.0000420]
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
response_time_lamport = [989.47, 991.88, 989.44, 995.98, 996.99, 1000.78]
response_time_ricart = [987.34, 988.78, 989.97, 989.90, 990.69, 990.49]

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
system_throughput_lamport = [0.000054,0.000053,0.000051,0.000049,0.0000489,0.0000479]
system_throughput_ricart = [0.000044,0.000042,0.000043,0.0000432,0.0000431,0.0000430]
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