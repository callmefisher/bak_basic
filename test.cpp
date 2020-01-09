#include <iostream>

int main() {
	unsigned int i = 1; // 0000 0000 0000 0001(大端)    1000 0000 0000 0000 (小端)
	if(*((char*)&i) == 0) {
		std::cout << ("this is big endian. \n");
	} else if(*((char*)&i) == 1) {
		std::cout << (" this is little endian. \n");
	}
}
