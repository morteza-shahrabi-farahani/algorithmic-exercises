#include <stdio.h>
int main () {
	long long n;
	int a;
	int c;
	int e;
	int d = 0;
	int b = 0;
	int f = 0;
	scanf("%lld" , &n);
	for( ; n >= 1 ; b = b+a) {
		a = n%10;
		n = n/10;
}
if (b < 10) {
	printf("%d" , b);
} else {
	for ( ; b >= 1 ; d = d+c) {
		c = b %10 ;
		b = b/10 ;
	}
    if ( d< 10) {
    	printf ("%d" , d);
	} else {
		for ( ; d >= 1 ; f = f+e) {
			e = d % 10;
			d = d/10 ;
		}
		printf ("%d" , f);
	}
}
}

