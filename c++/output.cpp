#include "output.h"

extern void generate_output(std::vector< std::vector<int> > snapshots, std::string name) {
    printf("Creating output for %s\n", name.c_str());
    for(int i=0; i<snapshots.size(); i++) {
        std::string data = "";
        std::string file = "test1-"+name;
        file = file+".out";
        std::ofstream outfile(file.c_str());
        for(int j=0; j<snapshots[i].size(); j++) {
            std::vector<int> curr_state = snapshots[i];
            for (auto num: curr_state)
                outfile << num << " ";
            outfile << "\n";
        }
        outfile.close();
    }
}

extern void print_output(std::vector< std::vector<int> > snapshots, std::string name) {
    printf("Creating output for %s\n", name.c_str());
    for(int i=0; i<snapshots.size(); i++) {
        for(int j=0; j<snapshots[i].size(); j++) {
            std::vector<int> curr_state = snapshots[i];
            for (auto num: curr_state)
                std::cout << num << " ";
            std::cout << "\n";
        }
    }
}