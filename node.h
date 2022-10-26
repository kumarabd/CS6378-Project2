#include <iostream>
#include <string>
#include <iterator>
#include <algorithm>
#include <vector>
#include <pthread.h>
#include <queue>
#include <ctime>
#include "application.h"

class Node {
    private:
        int id;
        int d; // inter-request delay
        int c; // cs-execution time
        int nr; // number of requests
        std::vector<Node *> neighbours;
        Lamport application; // Top Module
        Mutex mutex; // Bottom Module 

    public:
        Node();
        Node(int id, std::string h, int p, int d, int c, int nr);
        void start();
        int get_id();
        void info();
};
