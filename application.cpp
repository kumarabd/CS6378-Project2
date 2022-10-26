#include "application.h"

Lamport::Lamport() {};

Lamport::Lamport(int id) {
    this->id = id;
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