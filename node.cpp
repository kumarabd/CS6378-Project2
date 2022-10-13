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

Channel::Channel() {};

Channel::Channel(std::string h, int p) {
    // Create socket
    if ((this->server_fd = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
       perror("socket failed");
       exit(EXIT_FAILURE);
    }

    this->address.sin_family = AF_INET;
    //socket_address.sin_addr.s_addr = INADDR_ANY;
    this->address.sin_port = htons(p);

    hostent* hostname = gethostbyname(h.c_str());
    std::string st = std::string(inet_ntoa(**(in_addr**)hostname->h_addr_list));
    //inet_aton(st.c_str(), &socket_address.sin_addr);
    if (inet_aton(st.c_str(), &this->address.sin_addr) < 0) {
            perror("Address not recognized");
            exit(EXIT_FAILURE);
    }

    if (bind(this->server_fd, (struct sockaddr*) &this->address, sizeof(this->address)) < 0) {
        perror("bind failed");
        exit(EXIT_FAILURE);
    }
}

int Channel::fd() {
    return this->server_fd;
}

void Channel::start_socket() {
    // 3 is the max queue limit
    if (listen(this->server_fd, 3) < 0) {
        perror("listening");
        exit(EXIT_FAILURE);
    }
}

void Channel::send_socket(struct sockaddr_in serv_addr, std::string msg) {
    int sock = 0, valread, client_fd;
    if ((sock = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
        perror("\n Socket creation error \n");
        exit(EXIT_FAILURE);
    }
 
    // Convert IPv4 and IPv6 addresses from text to binary
    // form
    if (inet_pton(AF_INET, "127.0.0.1", &serv_addr.sin_addr) <= 0) {
        perror("\nInvalid address/ Address not supported \n");
        exit(EXIT_FAILURE);
    }
 
    if ((client_fd = connect(sock, (struct sockaddr*)&serv_addr, sizeof(serv_addr))) < 0) {
        perror("\nConnection Failed \n");
        exit(EXIT_FAILURE);
    }
    send(sock, msg.c_str(), strlen(msg.c_str()), 0);
    //close(client_fd);
}

Node::Node() {};

Node::Node(int id, std::string h, int p, int dMean, int cMean) {
    printf("Creating Node: %d\n", id);
    this->id = id;
    this->channel = Channel(h, p);

    // Assign random values for d and c
    this->d = generate_exponential_random(dMean);
    this->c = generate_exponential_random(cMean);

    // Create Mutual Exclusion service
    this->mService = MEService();
}

void Node::listen() {
    // start node server
    int newSocket;
    bool status = true;
    this->channel.start_socket();
    // Listen for messages
    while(status) {
        sockaddr_in addr = this->get_address();
        socklen_t addr_size = sizeof(addr);
        if ((newSocket = accept(this->channel.fd(), (struct sockaddr*)&addr, &addr_size)) < 0) {
            perror("unable to accept\n");
            exit(EXIT_FAILURE);
        }

        // read message
        char buffer[1024] = {0};
        int reader = read(newSocket, buffer, 1024);
        std::string delimiter = "///";
        std::string data = convertToString(buffer, sizeof(buffer)/sizeof(char));
        std::string source = data.substr(0, data.find(delimiter));
        std::string time = data.substr(1, data.find(delimiter));

        printf("message received on node %s from node %d for time %s", source.c_str(), this->get_id(), time.c_str());
        // eg: "0///<data>///0"
        // Active node send message to random node
        // if not then
        // Remove the node from network if the node reach maxNumber of messages
        message msg = { atoi(source.c_str()), atoi(time.c_str()) };
        status = this->process_message(msg);
    }
}
        

int Node::get_id() {
    return id;
}

bool Node::process_message(message msg) {
    return false;
}

struct sockaddr_in Node::get_address(){
    return this->channel.address;
}

void Node::send_message(Node * target_node, message msg){
    char* message("time");
    std::string m = std::to_string(this->get_id()) +"///";
    m = m + message;
    this->channel.send_socket(target_node->get_address(), m);
}


void Node::info() {
    printf("id: %d\n",this->id);
    printf("inter-request delay: %d\n",this->d);
    printf("cs-execution time: %d\n",this->c);
    //printf("channel: %d",this->channel);
}