int countSmileys(std::vector<std::string> arr) {
    int res = 0;
    for (const std::string &s: arr)
        if ((s.back() == ')' || s.back() == 'D') &&
            (s[s.size() - 2] == ':' || s[s.size() - 2] == ';' || s[s.size() - 2] == '-' || s[s.size() - 2] == '~'))
            ++res;

    return res;
}
