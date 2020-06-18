/*
 * @Author: gongluck 
 * @Date: 2020-06-18 11:08:13 
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-18 17:30:42
 */

#include <stdio.h>
#include <time.h>

#include <thread>

#include "../cghttp.h"

int TESTTIMES = 100;

void test_get()
{
    char* body = NULL;
    size_t bodylen = 0;
    int ret = Get("http://47.115.57.81/web", &body, &bodylen);
    //printf("%d\n%*s\n%zd\n", ret, body, bodylen);
    Release(&body);
}

void test_post()
{
    char* keys[] = {"name", "password"};
    char* values[] = {"gongluck", "testtest"};
    char* body = NULL;
    size_t bodylen = 0;
    int ret = Post("http://47.115.57.81/api/regist", keys, values, 2, &body, &bodylen);
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

    t1 = clock();
    for (int i = 0; i < TESTTIMES; ++i)
    {
        std::thread th(test_post);
        ths[i].swap(th);
    }
    for (int i = 0; i < TESTTIMES; ++i)
    {
        if (ths[i].joinable())
        {
            ths[i].join();
        }
    }
    t2 = clock();
    printf("%fms\n", difftime(t2, t1));

    getchar();
    return 0;
}
