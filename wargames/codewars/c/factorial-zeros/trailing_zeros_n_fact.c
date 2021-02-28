#include <stdlib.h>
#include <stdio.h>
#include <math.h>

int trailing_zeros_n_fact(int n) {
  int zeros = 0, i;
  float log5, v;
  log5 = log(5);
  for ( i = 1; i <= n; i++ ) {
    if ( i%10 == 0 ) {
			zeros++;
		} else if ( i%5 == 0 ) {
			v = log(i) / log5;
      /* printf("v: %f\n", v); */
      if (ceilf(v) == v)
				zeros += ( int )v;
			else
				zeros++;
		}
  }
  return zeros;
}

int main(){
  printf("Check [%d]: %d\n", 6,  trailing_zeros_n_fact(6));
  printf("Check [%d]: %d\n", 12, trailing_zeros_n_fact(12));
  printf("Check [%d]: %d\n", 30, trailing_zeros_n_fact(30));
  printf("Check [%d]: %d\n", 1000, trailing_zeros_n_fact(1000));
} 
