#include <stdio.h>
int main () {
	int c;
	int r;
	scanf ("%d%d" , &r , &c);
	if (c>10) {
		printf ("Left ");
	} else  {
		printf ("Right ");
	} 
	printf("%d " , 11 - r);
	if (c > 10) {
		printf ("%d " , 21-c);
	} else {
		printf("%d ", c);
	}
}

