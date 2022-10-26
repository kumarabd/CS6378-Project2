#include <iostream>
#include <cstdlib>
#include <fstream>
#include <pthread.h>
#include <unistd.h>
#include "node.h"
#include "readConfig.cpp"
using namespace std;

int main(int argc, char** argv)
{
    printf("Reading config\n");
	ReadConfig config = ReadConfig();
    config.read_config();

    printf("Creating node: %d\n", atoi(argv[1]));
    Node addr = Node(atoi(argv[1]), config.hostNames[atoi(argv[1])], config.ports[atoi(argv[1])], config.dMean, config.cMean, config.nr);
    return 0;
}