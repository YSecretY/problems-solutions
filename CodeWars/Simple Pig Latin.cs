using System;
using System.Collections.Generic;
using System.Linq;

public class Kata
{
    public static string PigIt(string str)
    {
        IEnumerable<string> words = str.Split(' ');
        words = words.Select(word => word.All(char.IsPunctuation) ? word : word[1..] + word[0] + "ay");
        return string.Join(" ", words);
    }
}
