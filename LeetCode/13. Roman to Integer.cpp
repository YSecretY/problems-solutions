class Solution {
public:
    int romanToInt(const std::string &s) {
        if (s.empty())
            return 0;

        std::unordered_map<char, unsigned> mp;
        mp['I'] = 1;
        mp['V'] = 5;
        mp['X'] = 10;
        mp['L'] = 50;
        mp['C'] = 100;
        mp['D'] = 500;
        mp['M'] = 1000;

        unsigned res = 0;
        for (int i = s.size() - 1; i >= 0; --i) {
            unsigned cur = mp[s[i]];
            unsigned prev = i > 0 ? mp[s[i - 1]] : 0;
            if (cur > prev) {
                res += cur - prev;
                --i;
            } else
                res += cur;
        }

        return res;
    }
};
