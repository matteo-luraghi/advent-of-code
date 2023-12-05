#include <ctype.h>
#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main() {
  FILE *f;
  int sum = 0;
  int c = 0, r = 0, stop = 0;
  int temp;

  f = fopen("3-a.txt", "r");

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
  f = fopen("3-a.txt", "r");

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
          printf("%c", m[i][j + p]);
          for (int k = -1; k <= 1; k++) {
            for (int l = -1; l <= 1; l++) {
              if (i + k >= 0 && i + k < c && j + p + l >= 0 && j + p + l < r &&
                  m[i + k][j + p + l] != '.' && !isdigit(m[i + k][j + p + l])) {
                printf("%c\n", m[i + k][j + p + l]);
                found = 1;
              }
            }
          }
        }
        if (found == 1) {
          int d = 0;
          for (int h = len - 1; h >= 0; h--) {
            sum += pow(10, h) * ((int)m[i][j + d] - 48);
            printf("SUM: %d\n", sum);
            d++;
          }
        }
        printf("\n");
        j += len - 1;
      }
    }
  }

  for (int i = 0; i < c; i++) {
    for (int j = 0; j < r; j++) {
      printf("%c", m[i][j]);
    }
    printf("\n");
  }

  printf("RESULT: %d\n", sum);
  return 0;
}
