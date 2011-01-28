#include "util.h"

#include <iostream>

using namespace std;

int main()
{
    const string str = "Hello world again  Moo";
    const char *delim = " ";
    vector<string> out;
    yku::Util::SplitString(str, delim, &out);

    vector<string>::iterator it = out.begin();
    while(it != out.end()) {
        cout << *it << endl;
        it++;
    }

}
