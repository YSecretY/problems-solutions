using System.Text;

public class Solution
{
    public string SortVowels(string s)
    {
        List<char> vowels = s.Where(IsVowel).ToList();
        vowels.Sort();

        StringBuilder builder = new();
        int i = 0;
        foreach (char c in s)
        {
            if (IsVowel(c))
            {
                builder.Append(vowels[i]);
                ++i;
            }
            else
            {
                builder.Append(c);
            }
        }

        return builder.ToString();
    }


    private static bool IsVowel(char c)
    {
        char lowerChar = char.ToLower(c);

        return lowerChar is 'a' or 'e' or 'i' or 'o' or 'u';
    }
}
