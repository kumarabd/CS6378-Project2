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
    int time;
} message;

class Application {
    public:
        virtual void request_cs(int id, Mutex * service);
        virtual void execute_cs(int duration);
        virtual void process_message(message msg);
};

class Lamport {
    private:
    public:
        std::list<int> reply_pending;
        Lamport();
        void request_cs(int id, Mutex * service);
        void execute_cs(int duration);
        void process_message(message msg);
};

class RicartAgarwala {
    public:
        RicartAgarwala();
        void request_cs(int id, Mutex * service);
        void execute_cs(int duration);
        void process_message(message msg);
};