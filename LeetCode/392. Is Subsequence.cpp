class Solution {
public:
    bool isSubsequence(string s, string t) {
        int j = 0, n = s.length();
        for (int i = 0; i < t.length() && j < n; ++i)
            if (s[j] == t[i])
                ++j;
        
        return j == n;
    }
};
