#ifndef __HELLOWORLD_H
#define __HELLOWORLD_H
#include <iostream>

#ifdef __cplusplus
extern "C" {
#endif


void hello_world() {
  std::cout << "Hello, CPP from go!";
}


#ifdef __cplusplus
}
#endif

#endif
