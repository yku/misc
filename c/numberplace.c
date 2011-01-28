#include <stdio.h>

int board[9][9] = { 6,0,0,0,0,0,1,0,8,
                    0,7,0,0,0,2,3,0,0,
                    0,0,5,8,0,0,9,0,0,
                    0,4,0,9,0,0,0,0,0,
                    0,5,0,0,4,0,0,6,0,
                    0,0,1,0,5,0,0,0,4,
                    0,0,0,6,3,0,0,0,0,
                    0,2,6,0,7,9,0,3,0,
                    0,0,3,0,0,0,0,0,0, };

int col = 9;
int row = 9;

int isPut(x, y, val)
{
    int i, j;
    int offset_x, offset_y;

    offset_x = x - x % 3;
    offset_y = y - y % 3;

    if(x >= col || y >= row) return 0;

    // すでに値が置いてある
    if(board[y][x] != 0) return 0;

    // 縦横で同じ値が存在すれば置けない
    for(i = 0; i < row; i++)
        if(board[i][x] == val || board[y][i] == val) return 0;

    // 3x3のブロック内に同じ値が存在すれば置けない
    for(i = 0; i < 3; i++) {
        for(j = 0; j < 3; j++) {
            if(board[offset_y+i][offset_x+j] == val) return 0; 
        }
    }
    return 1;
}

void dump() {
    int i, j;
    for(i = 0; i < row; i++) {
        for(j = 0; j < col; j++) {
            printf("%d ", board[i][j]);
        }
        printf("\n");
    }
}

int solve(int idx)
{
    int i;
    int x, y;
   
    x = idx % col;
    y = idx / row;

    if(idx >= row * col) {
        dump();
        return 1;
    }

    if(board[y][x] != 0) {
        solve(idx+1);
    }
    else{
        for(i = 1; i <= 9; i++) {
            if(isPut(x, y, i)) {
                board[y][x] = i;
                if(!solve(idx+1))
                    board[y][x] = 0;
                else
                    return 1;
            }
        }
    }
    return 0;
}

int main() 
{
    solve(0);
    return 0;
}
