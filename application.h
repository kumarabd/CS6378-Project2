#include <iostream>
#include <stdio.h>
#include <unistd.h>
#include <list>
#include "mutex.h"
#include "channel.h"
typedef struct {
    int id;
    struct sockaddr_in address;
} recepient;

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
};

class Lamport {
    private:
        Channel channel;
        int nr;
        int id;
        std::vector<recepient> recepients;
    public:
        std::list<int> reply_pending;
        Lamport();
        Lamport(int id, Channel channel, int nr);
        void cs_enter(clock_t time);
        void cs_leave();
        void process_message(message msg);
        void listen();
};

//class RicartAgarwala {
//    public:
//        RicartAgarwala();
//        RicartAgarwala(int id);
//};