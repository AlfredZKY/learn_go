#include <stdio.h>
#include <unistd.h>
#include <pthread.h>
#include <asm/prctl.h>
#include <sys/prctl.h>

__thread int g = 0;

void print_fs_base()
{
   unsigned long addr;
   int ret = arch_prctl(ARCH_GET_FS, &addr);  //获取fs段基地址
   if (ret < 0) {
       perror("error");
       return;
   }

   printf("fs base addr: %p\n", (void*)addr); //打印fs段基址

   return;
}

void* start(void* arg)
{
    print_fs_base(); //子线程打印fs段基地址
    printf("start, g[%p] : %d\n", &g, g);

    g++;

    return NULL;
}

int main(int argc, char* argv[])
{
    pthread_t tid;

    g = 100;

    pthread_create(&tid, NULL, start, NULL);
    pthread_join(tid, NULL);

    print_fs_base(); //main线程打印fs段基址
    printf("main, g[%p] : %d\n", &g, g);

    return 0;
}

// 可以看到：
// fs base addr: 0x7f829789a700
// start, g[0x7f829789a6fc] : 0
// fs base addr: 0x7f82980b8740
// main, g[0x7f82980b873c] : 100

// 子线程fs段基地址为0x7f829789a700，g的地址为0x7f829789a6fc，它正好是基地址 - 4

// 主线程fs段基地址为0x7f82980b8740，g的地址为0x7f82980b873c，它也是基地址 - 4