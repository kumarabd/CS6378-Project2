#include "node.h"

std::string convertToString(char* a, int size)
{
    int i;
    std::string s = "";
    for (i = 0; i < size; i++) {
        s = s + a[i];
    }
    return s;
}

Node::Node() {};

Node::Node(int id, std::string h, int p, int dMean, int cMean, int nr) {
    this->id = id;
    this->channel = Channel(h, p);
    this->nr = nr;

    // Assign random values for d and c
    this->d = generate_exponential_random(dMean);
    this->c = generate_exponential_random(cMean);

    //// Create Mutual Exclusion service
    //this->application = Lamport(id, channelc);
    //this->mutex = Mutex();
}

void Node::start_clock() {
    //std::clock_t prev_clock = std::clock();
    //while(this->nr > 0) {
    //    std::clock_t curr_clock = std::clock();
    //    float diff = std::round((float)(curr_clock - prev_clock)/CLOCKS_PER_SEC);
    //    // Request CS if difference in current clock value and previous clock is d
    //    if((int)diff == this->d) {
    //        printf("clocking time curr: %lu\n", curr_clock);
    //        // Request CS
    //        message msg = this->application.request_cs(this->get_id(), &this->mutex);
    //        // Broadcast request to other nodes
    //        for(Node * target_node: this->neighbours) {
    //            this->send_message(target_node, msg);
    //            this->application.reply_pending.push_back(target_node->get_id());
    //        }
    //        this->nr--;
    //        prev_clock = curr_clock;
    //    }
    //}
}

void Node::listen() {
    //// start node server
    //int newSocket;
    //this->channel.start_socket();
    //// Listen for messages
    //while(this->nr > 0) {
    //    printf("listening\n");
    //    // Receive request from other nodes
    //    sockaddr_in addr = this->channel.address;
    //    socklen_t addr_size = sizeof(addr);
    //    if ((newSocket = accept(this->channel.fd(), (struct sockaddr*)&addr, &addr_size)) < 0) {
    //        perror("unable to accept\n");
    //        exit(EXIT_FAILURE);
    //    }
    //    char buffer[1024] = {0};
    //    int reader = read(newSocket, buffer, 1024); // read message
    //    std::string delimiter = "///";
    //    std::string data = convertToString(buffer, sizeof(buffer)/sizeof(char));
    //    std::string source = data.substr(0, data.find(delimiter));
    //    std::string type = data.substr(1, data.find(delimiter));
    //    std::string time = data.substr(2, data.find(delimiter));
    //    printf("message received on node %s from node %d for time %s", source.c_str(), this->get_id(), time.c_str());
    //    // eg: "0///<data>///0"
    //    // Active node send message to random node
    //    // if not then
    //    // Remove the node from network if the node reach maxNumber of messages
    //    message msg = { atoi(source.c_str()), static_cast<MessageType>(atoi(type.c_str())),atoi(time.c_str()) };
    //    this->process_message(msg);
    //}
}
        

int Node::get_id() {
    return id;
}

void Node::process_message(message msg) {
    this->application.process_message(msg);
}

void Node::send_message(Node * target_node, message msg){
    //if(msg.type == RELEASE) {
    //    this->nr--; // Count the number of cs requests sent
    //}

    //std::string m = std::to_string(this->get_id()) +"///";
    //m = m + std::to_string(msg.type);
    //m = m + std::to_string(msg.time);
    //this->channel.send_socket(target_node->channel.address, m);
}


void Node::info() {
    printf("id: %d\n",this->id);
    printf("inter-request delay: %d\n",this->d);
    printf("cs-execution time: %d\n",this->c);
    //printf("channel: %d",this->channel);
}