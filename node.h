#include <iostream>
#include <string>
#include <iterator>
#include <algorithm>
#include <vector>
#include <pthread.h>
#include <queue>
#include <ctime>
#include <future>
#include "application.h"

class Node {
    private:
        int id;
        int d; // inter-request delay
        int c; // cs-execution time
        int nr; // number of requests
        std::clock_t payload; // The dely between the process and the request starts
        std::vector<recepient> neighbours;
        Lamport application; // Top Module
        Mutex mutex; // Bottom Module 

    public:
        Channel channel;
        Node();
        Node(int id, std::vector<std::string> h, std::vector<int> p, int d, int c, int nr);
        void start();
        int get_id();
        void info();
};
