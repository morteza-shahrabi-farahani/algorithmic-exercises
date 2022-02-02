#include <stdio.h>
int main () {
	int input;
	int i;
	int j;
	scanf("%d" , &input);
	for (i=1 ; i <= input ; i++) {
		for (j=1 ; j <= input ; j++) {
			printf("%d\t" , i*j);
		}
		j=1;
		printf("\n");
	}
}

