#include <iostream>

using namespace std;

int n, m;
char a[26][26];
int rmin = 999;

void go(int y, int x, int k) {
    if (a[y][x] == '$') {
        if (k < rmin) rmin = k;
        return;
    }
    a[y][x] = '.';
    if (a[y - 1][x] == ' ' || a[y - 1][x] == '$') {go(y - 1, x, k + 1);}
    if (a[y][x - 1] == ' ' || a[y][x - 1] == '$') {go(y, x - 1, k + 1);}
    if (a[y][x + 1] == ' ' || a[y][x + 1] == '$') {go(y, x + 1, k + 1);}
    if (a[y + 1][x] == ' ' || a[y + 1][x] == '$') {go(y + 1, x, k + 1);}
    a[y][x] = ' ';
}

int main()
{
    cin >> n >> m;
    int k = 0;
    getchar();
    for (int i = 0; i < n; ++i) {
        for (int j = 0; j < m; ++j)
            a[i][j] = getchar();
        getchar();
    }
    go(1, 1, 0);
    cout << rmin;
}
