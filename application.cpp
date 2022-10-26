#include "application.h"

Lamport::Lamport() {}

void Lamport::request_cs(int id, Mutex * service) {
    //// Broadcast request to all nodes
    //printf("Requesting cs\n");
    //service->cs_enter(id);
    //for(int target_node: this.neighbours) {
    //    this.send_message(target_node, msg);
    //    this->reply_pending.push_back(target_id);
    //}
    ////time = clock()
    ////msg = { curr_node.get_id(), REQUEST, time}
    //pair<int, int> p({time, curr_node.get_id})
    ////this.queue.push(p)
}

void Lamport::execute_cs(int duration) {
    usleep(duration);
}

void Lamport::process_message(message msg) {
    // For modified: remove the node from reply queue here
    switch(msg.type) {
        case REQUEST:
            // If it is in the cs then put in priority queue
            // If not then send a reply
            printf("%u message\n", msg.type);
            break;
        case REPLY:
            // Remove the node from reply queue
            this->reply_pending.remove(msg.source);
            printf("%u message\n", msg.type);
            break;
        case RELEASE:
            // Remove the node from the priority queue
            printf("%u message\n", msg.type);
            break;
        default:
            printf("message type invalid: %u\n", msg.type);
    }
}