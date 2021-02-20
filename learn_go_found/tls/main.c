#include <stdio.h>
#include <pthread.h>

// int g = 0;  // 1，定义全局变量g并赋初值0 不能用普通变量当全局变量
__thread int g = 0;  // 1，这里增加了__thread关键字，把g定义成私有的全局变量，每个线程都有一个g变量


void* start(void* arg)
{
    printf("start, g[%p] : %d\n", &g, g); // 4，子线程中打印全局变量g的地址和值

    g++; // 5，修改全局变量

    return NULL;
}

// gcc main.c -g -o thread -lpthread

int main(int argc, char* argv[])
{
    pthread_t tid;

    g = 100;  // 2，主线程给全局变量g赋值为100

    pthread_create(&tid, NULL, start, NULL); // 3， 创建子线程执行start()函数
    pthread_join(tid, NULL); // 6，等待子线程运行结束

    printf("main, g[%p] : %d\n", &g, g); // 7，打印全局变量g的地址和值

    return 0;
}

