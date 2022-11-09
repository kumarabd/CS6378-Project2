#include "channel.h"
#include <stdio.h>


Channel::Channel() {};

Channel::Channel(std::string h, int p) {
    // Create socket
    if ((this->server_fd = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
       perror("socket failed");
       exit(EXIT_FAILURE);
    }
    this->address.sin_family = AF_INET;
    this->address.sin_addr.s_addr = INADDR_ANY; // Use local loop back
    this->address.sin_port = htons(p);

    //// Convert hostname to address
    //struct hostent* host = gethostbyname(h.c_str());
    //unsigned long int addr=inet_addr(host);
    //host = gethostbyaddr(&addr, sizeof(addr), AF_INET);
    //std::string st = std::string(inet_ntoa(**(in_addr**)host->h_addr_list));
    //if (inet_aton(st.c_str(), &this->address.sin_addr) < 0) {
    //        perror("Address not recognized");
    //        exit(EXIT_FAILURE);
    //}

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
    if (listen(this->server_fd, 10) < 0) {
        perror("listening");
        exit(EXIT_FAILURE);
    }
}

void Channel::send_socket(struct sockaddr_in serv_addr, std::string msg) {
    this->client_socks[serv_addr.sin_port] = 0;
    if ((this->client_socks[serv_addr.sin_port] = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
        perror("\n Socket creation error \n");
        exit(EXIT_FAILURE);
    }
 
    // Convert IPv4 and IPv6 addresses from text to binary
    // form
    if (inet_pton(AF_INET, "0.0.0.0", &serv_addr.sin_addr) <= 0) {
        perror("\nInvalid address/ Address not supported \n");
        exit(EXIT_FAILURE);
    }

    int status = connect(this->client_socks[serv_addr.sin_port], (struct sockaddr*)&serv_addr, sizeof(serv_addr));
    if(status < 0) {
        perror("\nConnection Failed \n");
        exit(EXIT_FAILURE);
    }
    send(this->client_socks[serv_addr.sin_port], msg.c_str(), strlen(msg.c_str()), 0);
    //close(client_fd);
}

int Channel::if_socket(struct sockaddr_in serv_addr) {
    this->client_socks[serv_addr.sin_port] = 0;
    if ((this->client_socks[serv_addr.sin_port] = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
        perror("\n Socket creation error \n");
        exit(EXIT_FAILURE);
    }
 
    // Convert IPv4 and IPv6 addresses from text to binary
    // form
    if (inet_pton(AF_INET, "0.0.0.0", &serv_addr.sin_addr) <= 0) {
        perror("\nInvalid address/ Address not supported \n");
        exit(EXIT_FAILURE);
    }

    int status = connect(this->client_socks[serv_addr.sin_port], (struct sockaddr*)&serv_addr, sizeof(serv_addr));
    return status;
}