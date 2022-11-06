#include <string>
#include <sys/socket.h>
#include <sys/types.h> 
#include <netinet/in.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <cstring>
#include <map>

typedef struct {
    int sock;
    struct sockaddr sock_addr;
    int addr_len;
} connection_t;

class Channel {
    private:
        int size;
        int server_fd;
        std::map<in_port_t,int> client_socks;
    public:
        struct sockaddr_in address;
        Channel();
        Channel(std::string h, int p);
        void start_socket();
        void send_socket(struct sockaddr_in serv_addr, std::string msg);
        int fd();
        int if_socket(struct sockaddr_in serv_addr);
};