import sys
import traceback
import re


def concurrency_check(clock_values):
    num_nodes = len(clock_values[0])
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

def get_vector_from_log(file):
    f = open(file, 'r')
    vecs = []
    for line in f:
        if "vector time:" in line:
            m = re.search(r'\[.*?\]', line)
            st = m.group(0)[1:]
            m = re.search(r'app=(.*\s)', line)
            node = m.group(0).split(' ')[0].split('=')[1].split('-')[1]
            st = st[:len(st)-1]
            vec = st.split(' ')
            if len(vecs) == 0:
                vecs = [{} for _ in range(len(vec))]
            for idx, v in enumerate(vec):
                vecs[idx][node] = v
    return vecs

if __name__ == "__main__":
    vecs = get_vector_from_log("./build/log.txt")
    #size = len(vecs['0'])
    for v in range(len(vecs)-1):
        val1 = list(vecs[v].values())
        val2 = list(vecs[v+1].values())
        if concurrency_check([val1,val2]):
            print("{0} and {1} are concurrent".format(val1, val2))
            res = False

        

#if __name__ == "__main__":

#    if len(sys.argv) != 2:
#        print("Wrong number of arguments")
#        sys.exit(-1)

#    num_nodes = int(sys.argv[1])
#    # Store all the vector clock values in file clocks.txt
#    res, vec, clockFile = True, [], "clocks.txt"

#    print("Number of Processes {0}".format(num_nodes))
#    vec.append([0 for i in range(num_nodes)])

#    try:
#        cf = open(clockFile, 'r')
#        for line in cf:
#            temp = line
#            temp = temp.replace('[','')
#            temp = temp.replace(']', '')
#            vec.append([int(token) for token in temp.split(',')])

#            if concurrency_check(vec):
#                print("{0} and {1} are concurrent".format(vec[0], vec[1]))
#                res = False

#            del vec[0]

#        print("Result: {0}".format("PASS!" if res else "FAIL!"))
#    except IOError:
#        traceback.print_exc()