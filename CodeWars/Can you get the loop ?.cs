using System.Collections.Generic;

public class Kata
{
    public static int getLoopSize(LoopDetector.Node startNode)
    {
        int count = 0;
        Dictionary<LoopDetector.Node, int> wasHere = new();
        while (!wasHere.ContainsKey(startNode))
        {
            wasHere.Add(startNode, count);
            startNode = startNode.next;
            ++count;
        }

        return count - wasHere[startNode];
    }
}
