#include <stdio.h>
int main () {
	int grade;
	int days;
	int result;
	int counter;
	scanf("%d%d" , &grade , &days);
	if (days == 0) {
		result = 20;
	} else if (days == 7) {
		result = grade;
	} else {
		result = grade - days;
		if (result < 0) {
			result = 0;
		}
	}
	printf("%d" , result);
}

