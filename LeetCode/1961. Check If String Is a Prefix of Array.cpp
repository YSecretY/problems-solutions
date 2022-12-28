bool isPrefixString(string s, vector<string>& words) {
    string cur;
    for (auto it = words.begin(); it != words.end() && cur != s; ++it) {
        cur += *it;
    }
    return cur == s;
}
