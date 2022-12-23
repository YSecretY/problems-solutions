int majorityElement(vector<int>& nums) {
    map<int, int> mp;
    for (auto it : nums) {
        mp[it++];
    }
    for (auto it = mp.begin(); it != mp.end(); ++it) {
        it -> second = count(nums.begin(), nums.end(), it -> first);
        if (it -> second > nums.size() / 2)
            return it -> first;
    }
    return 0;
}
