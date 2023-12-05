#include <ctype.h>
#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct node {
  int x;
  int y;
  int num1;
  int num2;
} n;

int main() {
  FILE *f;
  int sum = 0;
  int c = 0, r = 0, stop = 0;
  int temp;
  char *filename = "3-b.txt";
  f = fopen(filename, "r");

  n save[1000];
  int savelen = 0;

  while (temp != EOF) {
    temp = fgetc(f);
    if (stop == 0 && temp != '\n') {
      c++;
    } else {
      stop = 1;
    }
    if (temp == '\n') {
      r++;
    }
  }

  fclose(f);
  f = fopen(filename, "r");

  int m[c][r];
  for (int i = 0; i < c; i++) {
    for (int j = 0; j < r; j++) {
      temp = fgetc(f);
      if (temp != '\n') {
        m[i][j] = temp;
      } else {
        j--;
      }
    }
  }

  int len;
  int found;
  char buf[3] = {0, 0, 0};

  for (int i = 0; i < c; i++) {
    for (int j = 0; j < r; j++) {
      if (m[i][j] != '.' && isdigit(m[i][j])) {
        found = 0;
        // get the len of the number
        len = 1;
        while (len + j < c && isdigit(m[i][j + len])) {
          len++;
        }
        for (int p = 0; p < len; p++) {
          // printf("%c", m[i][j + p]);
          for (int k = -1; k <= 1; k++) {
            for (int l = -1; l <= 1; l++) {
              if (i + k >= 0 && i + k < c && j + p + l >= 0 && j + p + l < r &&
                  m[i + k][j + p + l] == '*') {
                // printf("%c\n", m[i + k][j + p + l]);
                for (int y = 0; y < len; y++) {
                  buf[y] = m[i][j + y];
                }
                found = 0;
                for (int u = 0; u < savelen; u++) {
                  if (i + k == save[u].x && j + p + l == save[u].y) {
                    save[u].num2 = atoi(buf);
                    found = 1;
                  }
                }
                if (found == 0) {
                  save[savelen].x = i + k;
                  save[savelen].y = j + p + l;
                  save[savelen].num1 = atoi(buf);
                  savelen++;
                }
                buf[0] = 0;
                buf[1] = 0;
                buf[2] = 0;
              }
            }
          }
        }
        j += len - 1;
      }
      // printf("\n");
    }
  }

  for (int u = 0; u < savelen; u++) {
    if (save[u].num1 != 0 && save[u].num2 != 0) {
      printf("NUMS: [%d,%d]\n", save[u].num1, save[u].num2);
    }
    if (save[u].num1 != save[u].num2) {
      sum += save[u].num1 * save[u].num2;
    }
  }

  // for (int i = 0; i < c; i++) {
  //   for (int j = 0; j < r; j++) {
  //     printf("%c", m[i][j]);
  //   }
  //   printf("\n");
  // }
  fclose(f);
  printf("RESULT: %d\n", sum);
  return 0;
}
