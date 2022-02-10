#include <stdio.h>
int main () {
	int arar1;
	int arar2;
	int tedad;
	int javab;
	scanf("%d%d%d" , &arar1 , &arar2 , &tedad);
	if (tedad % 2 == 0) {
		javab = (tedad/2) * (arar1 + arar2);
		printf("%d" , javab);
	} else {
		javab = ((tedad/2) * (arar1 + arar2)) + arar1 ;
		printf("%d" , javab);
	}
}

