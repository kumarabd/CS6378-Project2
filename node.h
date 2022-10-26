#include <iostream>
#include <string>
#include <iterator>
#include <algorithm>
#include <vector>
#include <pthread.h>
#include <unistd.h>
#include <stdio.h>
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
        Channel channel;
        Lamport application; // Top Module
        Mutex mutex; // Bottom Module 

    public:
        Node();
        Node(int id, std::string h, int p, int d, int c, int nr);
        void cs_enter();
        void cs_leave();
        int get_id();
        void info();
        void listen();
        void start_clock();
        void process_message(message msg);
        void send_message(Node * target_node, message msg);
};
