public class Solution {
    public string ReverseWords(string s)
    {
        string[] words = s.Split(' ');
        StringBuilder builder = new StringBuilder();

        foreach (string word in words)
        {
            char[] arr = word.ToCharArray();
            Array.Reverse(arr);
            string reversedWord = new string(arr);
            builder.Append(reversedWord + ' ');
        }

        builder.Remove(builder.Length - 1, 1);

        return builder.ToString();
    }
}
