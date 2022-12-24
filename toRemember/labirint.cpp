#include <iostream>

using namespace std;

int n, m;
char a[26][26];

void go(int y, int x)
{
    a[y][x] = '.';
    if (a[y + 1][x] == ' ') go(y + 1, x);
    if (a[y][x - 1] == ' ') go(y, x - 1);
    if (a[y][x + 1] == ' ') go(y, x + 1);
    if (a[y - 1][x] == ' ') go(y - 1, x);
}

int main()
{
    cin >> n >> m;
    getchar();
    for (int i = 0; i < n; ++i) {
        for (int j = 0; j < m; ++j)
            a[i][j] = getchar();
        getchar();
    }

    go (1, 1);

    for (int i = 0; i < n; ++i) {
        for (int j = 0; j < m; ++j)
            cout << a[i][j];
        cout << endl;
    }
    return 0;
}
