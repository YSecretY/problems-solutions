std::string alphabet_position(const std::string &text) {
    std::string res;
    for (char c: text) {
        c = std::tolower(c);
        if (std::isalpha(c))
            res += std::to_string(c - 'a' + 1) + " ";
    }

    if (!res.empty()) res.pop_back();
    return res;
}
