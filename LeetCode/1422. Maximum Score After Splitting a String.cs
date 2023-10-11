public class Solution
{
    public int MaxScore(string s)
    {
        int maxScore = 0;
        for (int i = 1; i < s.Length; ++i)
        {
            int currentScore = 0;
            for (int j = 0; j < s.Length; ++j)
            {
                if (j < i && s[j] == '0') ++currentScore;
                if (j >= i && s[j] == '1') ++currentScore;
            }

            if (currentScore > maxScore) maxScore = currentScore;
        }

        return maxScore;
    }
}
