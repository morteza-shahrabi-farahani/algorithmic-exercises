#include <stdio.h>
int main () {
	int dama;
	scanf("%d" , &dama);
	if (dama > 100) {
		printf("Steam");
	} else if ( 0 <= dama && dama <= 100) {
		printf("Water");
	} else if (dama < 0) {
		printf("Ice");
	}
}

