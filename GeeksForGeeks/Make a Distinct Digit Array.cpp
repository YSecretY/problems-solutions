class Solution{
	public:
   	vector<int>  common_digits(vector<int>nums) {
   	    set<int> digits;
   	    for (auto it : nums) {
   	        int save = it;
   	        while (save) {
   	            digits.insert(save % 10);
   	            save /= 10;
   	        }
   	    }
   	    vector<int> res;
   	    for (auto it : digits) {
   	        res.push_back(it);
   	    }
   	    return res;
   	} 
};
