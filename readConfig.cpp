#include "readConfig.h"

int str_to_int(string s){
    int num = 0;
    for (int i = 0; i < s.length(); i++){
        char c = s[i];
        if (c == ' ' || int(c) != unsigned(int(c))) {
            continue;
        } else {
            num*=10;
            num+=c-'0';
        } 
    }
    return num;
}

ReadConfig::ReadConfig() {};

void ReadConfig::read_config() {
    ifstream testFile("config.txt");
    if (testFile.is_open()){
        while(getline(testFile, line)){
            //line.erase(remove_if(line.begin(), line.end(), ::isspace),line.end());
            if(line[0] == '#' || line.empty()) {
                    continue;
            }

            stringstream ss(line);
            getline(ss, str_nodestr, ' ');
            getline(ss, str_dMean, ' ');
            getline(ss, str_cMean, ' ');
            getline(ss, str_nr, ' ');

            node = str_to_int(str_nodestr);
            dMean = str_to_int(str_dMean);
            cMean = str_to_int(str_cMean);
            nr = str_to_int(str_nr);
            break;
        }

        for (int i = 0; i < node; i++) {
            getline(testFile, line);
            if(line[0] == '#' || line.empty()) {
                i--;
                continue;
            } 
            else if (line.find('#') != string::npos) {
                line = line.substr(0,line.find('#'));
            } 
            int nodeid, portno;
            string nodeidstr, portnostr, hostname;
            stringstream ss(line);
            ss >> nodeidstr;
            ss >> hostname;
            ss >> portnostr;

            nodeid = str_to_int(nodeidstr);
            portno = str_to_int(portnostr);
            hostNames.push_back(hostname);
            ports.push_back(portno);

        }
     } else{
        cerr << "File failed to open" << endl;
    }
}