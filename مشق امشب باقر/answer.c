#include <stdio.h>
int main () {
	int a , b , c;
	scanf("%d%d%d" , &a , &b , &c);
	if (a == 0 || b == 0 || c == 0) {
		printf("No");
	} else if (a+b+c == 180) {
		printf("Yes");
	} else {
		printf("No");
	}
}

