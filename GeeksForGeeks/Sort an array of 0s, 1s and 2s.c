void sort012(int a[], int n)
{
    int zero_count = 0, one_count = 0, two_count = 0;
    for (int i = 0; i < n; ++i) {
        if (a[i] == 0) ++zero_count;
        else if (a[i] == 1) ++one_count;
        else if (a[i] == 2) ++two_count;
    }
    for (int i = 0; i < n; ++i) {
        if (zero_count != 0) { a[i] = 0; --zero_count; }
        else if (one_count != 0) { a[i] = 1; --one_count; }
        else if (two_count != 0) { a[i] = 2; --two_count; }
    }
}