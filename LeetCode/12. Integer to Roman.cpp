class Solution {
public:
    string intToRoman(int num) {
        string roman = "";
        vector<pair<int, string> > storeIntRoman = {{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"}, {100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"}, {10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"}};
        for (const auto &p : storeIntRoman) {
            while (num >= p.first) {
                roman += p.second;
                num -= p.first;
            }
        }
        
        return roman;
    }
};
