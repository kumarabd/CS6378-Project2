#include <iostream>
#include <random>
#include<time.h>

int generate_exponential_random(int mean);

class Mutex {
    private:
        int c;
        int id;
    public:
        Mutex();
        Mutex(int id, int c);
        void execute_cs(std::clock_t payload);
};