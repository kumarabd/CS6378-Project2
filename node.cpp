#include "node.h"

std::vector<std::future<void>> pending_futures;

Node::Node() {};

Node::Node(int id, std::vector<std::string> h, std::vector<int> p, int dMean, int cMean, int nr) {
    this->id = id;
    this->nr = nr;
    Channel channel = Channel(h[id-1], p[id-1]);

    // Create Mutual Exclusion service
    std::vector<recepient> rcpts;
    for(int i=0; i<h.size(); i++) {
        if(i+1 != id) {
            struct sockaddr_in addr;
            addr.sin_family = AF_INET;
            addr.sin_addr.s_addr = INADDR_ANY;
            addr.sin_port = htons(p[i]);

            // Convert hostname to address
            //hostent* hostname = gethostbyname(h[i].c_str());
            //std::string st = std::string(inet_ntoa(**(in_addr**)hostname->h_addr_list));
            //inet_aton(st.c_str(), &addr.sin_addr);
            recepient member = { .id=i+1, .address= addr};
            rcpts.push_back(member);
        } 
    }
    this->application = Lamport(id, channel, rcpts);
    this->mutex = Mutex();

    // Assign random values for d and c
    this->d = generate_exponential_random(dMean);
    this->c = generate_exponential_random(cMean);
}


int Node::get_id() {
    return id;
}

void Node::start() {
    // Start the application
    std::future<void> ft = std::async(std::launch::async, &Lamport::listen, this->application);
    pending_futures.push_back(std::move(ft));

    // Start requests
    std::clock_t prev_clock = std::clock();
    this->d = this->d*CLOCKS_PER_SEC + prev_clock;
    while(this->nr > 0) {
        std::clock_t curr_clock = std::clock();
        unsigned long diff = curr_clock - prev_clock;
        // Request CS if difference in current clock value and previous clock is d
        // /10 because the microseconds are skipped so fast
        if(diff/10 == (unsigned long)this->d/10) {
            printf("requesting at: %lu\n", curr_clock);
            this->application.cs_enter(curr_clock);
            this->mutex.execute_cs(); // pending
            this->application.cs_leave(); // pending
            this->nr--;
            prev_clock = curr_clock;
        }
    }
    ft.wait();
}

void Node::info() {
    printf("id: %d\n",this->id);
    printf("inter-request delay: %d\n",this->d);
    printf("cs-execution time: %d\n",this->c);
    //printf("channel: %d",this->channel);
}