#include <iostream>
#include <fstream>
#include <sstream>
#include <string> 
#include <vector>
#include <algorithm>

using namespace std;

class ReadConfig
{
private:
    string str_nodestr, str_dMean, str_cMean, str_nr, line;
public:
    int node, dMean, cMean, nr;
    vector <string> hostNames;
    vector <int> ports;
    
    ReadConfig();
    void read_config();
};

