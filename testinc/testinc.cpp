/*
 * @Author: gongluck 
 * @Date: 2020-06-18 11:08:13 
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-18 11:14:50
 */

#include <stdio.h>
#include <time.h>

#include <thread>

#include "../cghttp.h"

int TESTTIMES = 1000;

void test_get()
{
    char* body = NULL;
    size_t bodylen = 0;
    int ret = Get("https://www.baidu.com", &body, &bodylen);
    //printf("%d\n%s\n%zd\n", ret, body, bodylen);
    Release(&body);
}

int main()
{
    std::thread* ths = new std::thread[TESTTIMES];
    clock_t t1 = clock();
    for (int i = 0; i < TESTTIMES; ++i)
    {
        std::thread th(test_get);
        ths[i].swap(th);
    }
    for (int i = 0; i < TESTTIMES; ++i)
    {
        if (ths[i].joinable())
        {
            ths[i].join();
        }
    }
    clock_t t2 = clock();
    printf("%fms\n", difftime(t2, t1));
    getchar();
    return 0;
}
