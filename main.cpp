#include <iostream>
#include <cstdlib>
#include <fstream>
#include <pthread.h>
#include <unistd.h>
#include "readConfig.cpp"
#include "network.h"
using namespace std;

typedef struct thread_data {
    int id;
    std::string host;
    int port;
    int dMean;
    int cMean;
} thread_data;

std::vector<Node> nodes;

void * create_nodes(thread_data* args) {
    thread_data *tdata = (thread_data *)args;
    Node addr = Node(tdata->id, tdata->host, tdata->port, tdata->dMean, tdata->cMean);
    nodes.push_back(addr);
    return NULL;
}

int main()
{
    printf("Reading config\n");
	ReadConfig config = ReadConfig();
    config.read_config();

    printf("Gathering nodes\n");
    pthread_t * threads = new pthread_t[config.node];
    thread_data * args = new thread_data[config.node];
    for(int i=0; i<config.node; i++) {
        // Create thread
        args[i].id = i;
        args[i].host = config.hostNames[i];
        args[i].port = config.ports[i];
        args[i].dMean = config.dMean;
        args[i].cMean = config.cMean;

        if(pthread_create(&threads[i], NULL, (void* (*)(void*))  (&create_nodes), (void *) &args[i])) {
            perror("thread create failed");
            exit(EXIT_FAILURE);
        }
        usleep(clock);
    }

    // create a network
    printf("Creating network\n");
    Network network = Network(nodes);
    printf("Network created\n");

    // run topology
    network.run();
    network.save();
    
    // wait for the threads
    void *status;
    for(int i=0; i<config.node; i++) {
        pthread_join(threads[i], &status);
    }

    return 0;
}