#include <iostream>
#include <stdio.h>
#include <unistd.h>
#include <list>
#include <queue>
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
    int id;
    clock_t time;
} queue_object;

class Compare
{
public:
    bool operator() (queue_object a, queue_object b) {
        if(a.time > b.time) {
            return true;
        } else {
            return false;
        }
    }
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
        int id;
        std::vector<recepient> recepients;
        std::priority_queue<queue_object, std::vector<queue_object>, Compare> pq;
        //std::priority_queue<queue_object> pq;
    public:
        std::list<int> reply_pending;
        Lamport();
        Lamport(int id, Channel channel, std::vector<recepient> rcpts);
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