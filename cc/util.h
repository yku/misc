#ifndef _UTIL_H_
#define _UTIL_H_

#include <vector>
#include <string>

using namespace std;

namespace yku {
class Util {
public:
    static void SplitString(const string &str,
                            const char *delim,
                            vector<string> *output);

private:
    Util() {}
    virtual ~Util() {}
};
} // namespace yku

#endif
