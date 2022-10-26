#include <iostream>
#include <stdio.h>
#include <unistd.h>
#include <list>
#include "mutex.h"
#include "channel.h"

enum MessageType {
    REQUEST=0,
    REPLY=1,
    RELEASE=2,
};

typedef struct {
    int source;
    MessageType type;
    clock_t time;
} message;

class Application {
    public:
        virtual void process_message(message msg);
};

class Lamport {
    private:
        int id;
    public:
        std::list<int> reply_pending;
        Lamport();
        Lamport(int id);
        void process_message(message msg);
};

class RicartAgarwala {
    public:
        RicartAgarwala();
        RicartAgarwala(int id);
        void process_message(message msg);
};