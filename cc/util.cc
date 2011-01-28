
#include "util.h"

#include <iterator>
#include <vector>
#include <string>

using namespace std;

namespace {
template <class ITR>
void SplitStringToIteratorUsing(const string &str,
                                const char *delim,
                                ITR &result) {
    string::size_type begin_index, end_index;
    begin_index = str.find_first_not_of(delim);
    while(begin_index != string::npos) {
        end_index = str.find_first_of(delim, begin_index);
        if (end_index == string::npos) {
            *result++ = str.substr(begin_index);
            return;
        }
        *result++ = str.substr(begin_index, (end_index - begin_index));
        begin_index = str.find_first_not_of(delim, end_index);
    }
}
} // namespace

namespace yku {
void Util::SplitString(const string &str,
                       const char *delim,
                       vector<string> *output) {
    back_insert_iterator<vector<string> > it(*output);
    SplitStringToIteratorUsing(str, delim, it);
}
} // namespace yku

