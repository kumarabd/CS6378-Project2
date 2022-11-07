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
    DEF_REPLY=3
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
        void send_reply(int target, message msg);
    public:
        std::list<int> reply_pending;
        bool in_cs;
        Lamport();
        Lamport(int id, Channel channel, std::vector<recepient> rcpts);
        void cs_enter(clock_t time);
        void cs_leave();
        bool process_message(message msg);
        void listen();
};

class RicartAgarwala {
    private:
        Channel channel;
        int id;
        std::vector<recepient> recepients;
        clock_t csIntendTime;
        bool unfulfilled_request;
        std::list<int> def_reply_pending;
        std::vector<recepient> def_release_send_list;
        void send_reply(int target, message msg);
    public:
        bool in_cs;
        RicartAgarwala();
        RicartAgarwala(int id, Channel channel, std::vector<recepient> rcpts);
        void cs_enter(clock_t time);
        void cs_leave();
        bool process_message(message msg);
        void listen();
};