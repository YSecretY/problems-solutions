int lengthOfLastWord(string s) {
        if (s.empty())
            return 0;
        int length = 0;
        int k = s.length() - 1;
        for (k; k >= 0 && s[k] == ' '; --k) {}
        for (int i = k; i >= 0 && s[i] != ' '; --i) {
            ++length;
        }
        return length;
}
