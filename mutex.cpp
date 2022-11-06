#include "mutex.h"

int generate_exponential_random(int mean) {
    double const mean_new = mean;
    double const lambda = 1/mean_new;
    std::random_device rd;
    std::mt19937 gen(rd());
    std::exponential_distribution<> num(lambda);
    return num(gen);
}

Mutex::Mutex() {};

Mutex::Mutex(int id, int c) {
    this->id = id;
    this->c = c;
};

void Mutex::execute_cs(std::clock_t payload) {
    // Use payload to cover up the time difference between when the program started and the time when this function starts
    this->c = this->c*CLOCKS_PER_SEC;
    std::clock_t prev_clock = std::clock();
    unsigned long diff = 0;
    printf("Node %d is entering CS at %lu\n", this->id, prev_clock);
    while(diff/10 == (unsigned long)this->c/10) {
        std::clock_t curr_clock = std::clock();
        unsigned long diff = curr_clock - prev_clock;
    }
    printf("Node %d is leaving CS at %lu\n", this->id, curr_clock);
};
