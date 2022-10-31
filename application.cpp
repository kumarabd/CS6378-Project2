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

Lamport::Lamport(int id, Channel channel, std::vector<recepient> rcpts) {
    this->id = id;
    this->channel = channel;
    this->recepients = rcpts;
    this->in_cs = false;
}

void Lamport::send_reply(int target, message msg) {
    std::string m = std::to_string(msg.source) +"///";
    m = m + std::to_string(msg.type);
    m = m + "///";
    m = m + std::to_string(msg.time);
    this->channel.send_socket(this->recepients[target-1].address, m);
}

bool Lamport::process_message(message msg) {
    // For modified: remove the node from reply queue here
    switch(msg.type) {
        case REQUEST: {
            // If it is in the cs then put in priority queue
            // If not then send a reply
            printf("Request message received on node %d from node %d at %lu\n", this->id, msg.source, msg.time);
            // Add to priority queue
            queue_object obj = { .id=msg.source, .time=msg.time};
            this->pq.push(obj);
            if(!this->in_cs) {
                message mobj = { .source=this->id, .type=REPLY, .time=clock() };
                this->send_reply(msg.source, mobj);
            }
            break;
        }
        case REPLY: {
            // Remove the node from reply queue
            this->reply_pending.remove(msg.source);
            printf("Reply message received on node %d from node %d at %lu\n", this->id, msg.source, msg.time);
            break;
        }
        case RELEASE: {               
            this->pq.pop();
            //this->reply_pending.remove(msg.source);
            printf("Release message received on node %d from node %d at %lu\n", this->id, msg.source, msg.time);
            // Remove the node from the priority queue
            break;
        }
        default:
            printf("message type invalid: %u\n", msg.type);
    }
    // Condition to enter CS
    if(!this->reply_pending.size() && this->pq.top().id == this->id) {
        return true;
    }
    return false;
}

void Lamport::cs_enter(clock_t time) {
    // Construct message
    message msg = {.source = this->id, .type = REQUEST, .time = time};
    std::string m = std::to_string(this->id) +"///";
    m = m + std::to_string(msg.type);
    m = m + "///";
    m = m + std::to_string(msg.time);

    // Broadcast request to other nodes
    for(recepient target: this->recepients) {
        printf("Sending request from node %d to node %d at %lu\n", this->id, target.id, time);
        this->channel.send_socket(target.address, m);
        this->reply_pending.push_back(target.id);
    }

    // Add the timestamp to the priority queue
    queue_object obj = { .id=this->id, .time=time };
    this->pq.push(obj);

    while(!this->in_cs) {};
}

void Lamport::cs_leave() {
    // Pop them off the queue
    this->pq.pop();
    
    // Construct message
    message msg = {.source = this->id, .type = RELEASE, .time = clock()};
    std::string m = std::to_string(this->id) +"///";
    m = m + std::to_string(msg.type);
    m = m + "///";
    m = m + std::to_string(msg.time);

    // Send release to all candidates
    for(recepient target: this->recepients) {
        printf("Sending release from node %d to node %d at %lu\n", this->id, target.id, msg.time);
        this->channel.send_socket(target.address, m);
    }
}

void Lamport::listen() {
    // start node server
    int newSocket;
    this->channel.start_socket();
    // Listen for messages
    while(1) {
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
        data.erase(0, data.find(delimiter) + delimiter.length());
        std::string type = data.substr(0, data.find(delimiter));
        data.erase(0, data.find(delimiter) + delimiter.length());
        std::string time = data.substr(0, data.find(delimiter));
        // eg: "0///<data>///0"
        // Active node send message to random node
        // if not then
        // Remove the node from network if the node reach maxNumber of messages
        message msg = { atoi(source.c_str()), static_cast<MessageType>(atoi(type.c_str())), (clock_t)std::stod(time.c_str()) };
        bool result = this->process_message(msg); // Make this future
        if(result) {
            // Enter CS condition
            this->in_cs = true;
        }
    }
}