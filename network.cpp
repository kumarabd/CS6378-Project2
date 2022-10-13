#include "network.h"
#include "output.h"

std::vector<std::future<void>> pending_futures;

Network::Network(std::vector<Node> nodes) {
    this->nodes = nodes;
    this->number_of_nodes = this->nodes.size();
}

void Network::run() {
    printf("Running the network\n");

    // Start the nodes
    for(int i =0; i< this->nodes.size();i++) {
        std::future<void> ft = std::async(std::launch::async, &Node::listen, this->nodes[i]);
        pending_futures.push_back(std::move(ft));
    }

    usleep(clock);

    // Send process messages to start with process 0
    message msg = { 0, 0 };
    this->nodes[0].process_message(msg);
}

void Network::save() {
    //for(int i=0; i<this->nodes.size(); i++) {
    //    generate_output(this->nodes[i].states, std::to_string(this->nodes[i].get_id()));
    //}
}