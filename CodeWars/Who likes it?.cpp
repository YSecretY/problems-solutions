#include <string>
#include <vector>

std::string likes(const std::vector<std::string> &names) {
    std::string::size_type n = names.size();
    std::string res = "no one likes this";
    if (n == 1)
        res = names.front() + " likes this";
    else if (n == 2)
        res = names.front() + " and " + names.back() + " like this";
    else if (n == 3)
        res = names.front() + ", " + names[1] + " and " + names.back() + " like this";
    else if (n >= 4)
        res = names.front() + ", " + names[1] + " and " + std::to_string(n - 2) + " others like this";

    return res;
}
