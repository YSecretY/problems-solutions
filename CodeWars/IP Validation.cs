namespace Solution
{
    class Kata
    {
        public static bool IsValidIp(string ipAddres)
        {
            string[] arr = ipAddres.Split('.');
            if (arr.Length != 4)
                return false;

            foreach (string value in arr)
            {
                try
                {
                    if (value.Length > 1 && value[0] == '0')
                        return false;

                    if (value.Contains(' '))
                        return false;

                    Int16 number = Convert.ToInt16(value);
                    if (number is > 255 or < 0)
                        return false;
                }
                catch (FormatException)
                {
                    return false;
                }
            }

            return true;
        }
    }
}
