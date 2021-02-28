#include <stdlib.h>
#include <stdio.h>

long zeros(long n) {
  long zeros = 0, i;
  for ( i = 1; i <= n; i++ ) {
    if ( i%10 == 0 ) {
      long j = i;
      do{
        zeros++;
        j = j/10;
      }while (j%10==0);
      while(j%5==0) {
        zeros++;
        j = j/5;
      }
		} else if(i%5==0) {
      long j = i;
      do {
        zeros++;
        j = j/5;
      }while(j%5==0);
		}
  }
  return zeros;
}

int main(){
  printf("Check [%d]: %d\n", 6,  zeros(6));
  printf("Check [%d]: %d\n", 12, zeros(12));
  printf("Check [%d]: %d\n", 30, zeros(30));
  printf("Check [%d]: %d\n", 1000, zeros(1000));
  printf("Check [%d]: %d\n", 100000, zeros(100000));
  printf("Check [%d]: %d\n", 1000000000, zeros(1000000000));
} 
