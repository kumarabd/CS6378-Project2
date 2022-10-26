#include "mutex.h"

int generate_exponential_random(int mean) {
    double const mean_new = mean;
    double const lambda = 1/mean_new;
    std::random_device rd;
    std::mt19937 gen(rd());
    std::exponential_distribution<> num(lambda);
    return num(gen);
}

Mutex::Mutex() {}
void Mutex::cs_enter(int process_id) {
    // Permission to enter CS for a process
}

void Mutex::cs_leave(int process_id) {
    // Inform the service that the process has finished execution
}