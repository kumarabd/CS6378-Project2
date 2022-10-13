#include<iostream>
#include<string>
#include <iterator>
#include <algorithm>
#include <vector>
#include<pthread.h>
#include <unistd.h>
#include<sys/socket.h>
#include<sys/types.h> 
#include<cstring>
#include <stdio.h>
#include<netinet/in.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <queue>
#include "MEService.h"
#define clock 1000000

typedef struct {
    int sock;
    struct sockaddr sock_addr;
    int addr_len;
} connection_t;

typedef struct {
    int source;
    int time;
} message;

class Channel {
    private:
        int size;
        int server_fd;
    public:
        struct sockaddr_in address;
        Channel();
        Channel(std::string h, int p);
        void start_socket();
        void send_socket(struct sockaddr_in serv_addr, std::string msg);
        int fd();
};

class Node {
    private:
        int id;
        int d; // inter-request delay
        int c; // cs-execution time
        std::priority_queue<int> queue;
        Channel channel;
        MEService mService;

    public:
        Node();
        Node(int id, std::string h, int p, int d, int c);
        int get_id();
        bool process_message(message msg);
        void send_message(Node * node, message msg);
        struct sockaddr_in get_address();
        void info();
        void listen();
};
