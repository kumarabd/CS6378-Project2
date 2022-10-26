#include <random>
#include <stdio.h>
#include <map>
#include <iostream>
#include <iomanip>

int generate_exponential_random(int mean);

class Mutex {
    public:
        Mutex();
        void cs_enter(int process_id);
        void cs_leave(int process_id);
};