int singleNumber(vector<int>& nums) {
    map <int, int> mp;
    for (auto it : nums) {
        mp[it++];
    }
    for (auto& it : mp) {
        it.second = count(nums.begin(), nums.end(), it.first);
        if (it.second == 1)
            return it.first;
    }
return 0;
}
