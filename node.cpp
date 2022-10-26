#include "node.h"

Node::Node() {};

Node::Node(int id, std::string h, int p, int dMean, int cMean, int nr) {
    this->id = id;

    Channel channel = Channel(h, p);

    // Create Mutual Exclusion service
    this->application = Lamport(id, channel, nr);
    this->mutex = Mutex();
    

    // Assign random values for d and c
    this->d = generate_exponential_random(dMean);
    this->c = generate_exponential_random(cMean);
}
        

int Node::get_id() {
    return id;
}

void Node::start() {
    std::clock_t prev_clock = std::clock();
    while(this->nr > 0) {
        std::clock_t curr_clock = std::clock();
        float diff = std::round((float)(curr_clock - prev_clock)/CLOCKS_PER_SEC);
        // Request CS if difference in current clock value and previous clock is d
        if((int)diff == this->d) {
            printf("clocking time curr: %lu\n", curr_clock);
            this->application.cs_enter(curr_clock);
            this->mutex.execute_cs(); // pending
            this->application.cs_leave(); // pending
            this->nr--;
            prev_clock = curr_clock;
        }
    }
}

void Node::info() {
    printf("id: %d\n",this->id);
    printf("inter-request delay: %d\n",this->d);
    printf("cs-execution time: %d\n",this->c);
    //printf("channel: %d",this->channel);
}