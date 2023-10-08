using System.Collections.Generic;

public class Kata
{
    public static string FirstNonRepeatingLetter(string s)
    {
        Dictionary<char, int> letterCount = new();
        foreach (char c in s)
        {
            char letterInLower = char.ToLower(c);
            if (letterCount.ContainsKey(letterInLower)) letterCount[letterInLower]++;
            else letterCount[letterInLower] = 1;
        }

        string res = "";
        foreach (KeyValuePair<char, int> pair in letterCount)
        {
            if (pair.Value == 1)
            {
                res += s.Contains(pair.Key) ? pair.Key : char.ToUpper(pair.Key);
                return res;
            }
        }

        return res;
    }
}
