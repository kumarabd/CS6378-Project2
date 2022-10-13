#include<iostream>
#include <vector>
#include <stdio.h>
#include <future>
#include "node.h"

class Network {
    private:
        int number_of_nodes;
        std::vector<Node> nodes;
    public:
        Network(std::vector<Node> nodes);
        void run();
        void save();
};