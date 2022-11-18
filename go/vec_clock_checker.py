import sys
import traceback


def concurrency_check(clock_values):

    result, less, great = False, False, False
    first, second = clock_values[0], clock_values[1]
    # print(first, second)
    
    if len(first) != len(second):
        return True

    for k in range(num_nodes):
        if first[k] < second[k]:
            less = True
        if first[k] > second[k]:
            great = True
        if less and great:
            result = True
            break

    return result

if __name__ == "__main__":

    if len(sys.argv) != 2:
        print("Wrong number of arguments")
        sys.exit(-1)

    num_nodes = int(sys.argv[1])
    # Store all the vector clock values in file clocks.txt
    res, vec, prev, clockFile = True, [], [], "clocks.txt"

    print("Number of Processes {0}".format(num_nodes))
    vec.append([0 for i in range(num_nodes)])

    try:
        cf = open(clockFile, 'r')
        for line in cf:
            temp = line
            temp = temp.replace('[','')
            temp = temp.replace(']', '')
            vec.append([int(token) for token in temp.split(',')])

            if concurrency_check(vec):
                print("{0} and {1} are concurrent".format(vec[0], vec[1]))
                res = False

            del vec[0]

        print("Result: {0}".format("PASS!" if res else "FAIL!"))
    except IOError:
        traceback.print_exc()