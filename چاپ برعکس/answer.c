#include <stdio.h>
int main () {
	int input[1000];
	int counter , counter2 , counter3;
	for (counter = 0 ; counter < 1000 ; counter++ ) {
		scanf("%d" , &input[counter]);
		if (input[counter] == 0) {
			break;
		}
	}
	counter3 = counter - 1;
	for(counter2 = 0 ; counter2 <= counter3 ; counter3-- ) {
		printf("%d\n" , input[counter3]);
	}
}

