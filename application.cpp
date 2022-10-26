#include "application.h"

std::string convertToString(char* a, int size) {
    int i;
    std::string s = "";
    for (i = 0; i < size; i++) {
        s = s + a[i];
    }
    return s;
}

Lamport::Lamport() {};

Lamport::Lamport(int id, Channel channel, int nr) {
    this->channel = channel;
    this->nr = nr;
}

void Lamport::process_message(message msg) {
    // For modified: remove the node from reply queue here
    switch(msg.type) {
        case REQUEST:
            // If it is in the cs then put in priority queue
            // If not then send a reply
            printf("%u message\n", msg.type);
            break;
        case REPLY:
            // Remove the node from reply queue
            this->reply_pending.remove(msg.source);
            printf("%u message\n", msg.type);
            break;
        case RELEASE:
            // Remove the node from the priority queue
            printf("%u message\n", msg.type);
            break;
        default:
            printf("message type invalid: %u\n", msg.type);
    }
}

void Lamport::cs_enter(clock_t time) {
    // Construct message
    message msg = {.source = this->id, .type = REQUEST, .time = time};
    std::string m = std::to_string(this->id) +"///";
    m = m + std::to_string(msg.type);
    m = m + std::to_string(msg.time);

    // Broadcast request to other nodes
    for(recepient target: this->recepients) {
        this->channel.send_socket(target.address, m);
        this->reply_pending.push_back(target.id);
    }
}

void Lamport::cs_leave() {
}

void Lamport::listen() {
    // start node server
    int newSocket;
    this->channel.start_socket();
    // Listen for messages
    while(this->nr > 0) {
        printf("listening\n");
        // Receive request from other nodes
        sockaddr_in addr = this->channel.address;
        socklen_t addr_size = sizeof(addr);
        if ((newSocket = accept(this->channel.fd(), (struct sockaddr*)&addr, &addr_size)) < 0) {
            perror("unable to accept\n");
            exit(EXIT_FAILURE);
        }
        char buffer[1024] = {0};
        int reader = read(newSocket, buffer, 1024); // read message
        std::string delimiter = "///";
        std::string data = convertToString(buffer, sizeof(buffer)/sizeof(char));
        std::string source = data.substr(0, data.find(delimiter));
        std::string type = data.substr(1, data.find(delimiter));
        std::string time = data.substr(2, data.find(delimiter));
        printf("message received on node %s from node %d for time %s", source.c_str(), this->id, time.c_str());
        // eg: "0///<data>///0"
        // Active node send message to random node
        // if not then
        // Remove the node from network if the node reach maxNumber of messages
        message msg = { atoi(source.c_str()), static_cast<MessageType>(atoi(type.c_str())), (clock_t)std::stod(time.c_str()) };
        this->process_message(msg);
    }
}