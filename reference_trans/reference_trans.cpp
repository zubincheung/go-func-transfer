//
//  main.cpp
//  test
//
//  Created by Zubin Cheung on 2018/6/13.
//  Copyright © 2018年 Zubin Cheung. All rights reserved.
//

#include <stdio.h>

void inc(int &x)
{
	printf("x 内存地址:%p 值:%d \n", &x, x);
	x++;
	printf("x:%d \n", x);
}

int main(int argc, const char *argv[])
{
	int a = 1;
	printf("a 内存地址:%p 值:%d \n", &a, a);
	inc(a);
	printf("执行后 a:%d \n", a);
	return 0;
}
