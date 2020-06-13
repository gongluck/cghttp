#include <stdio.h>
#include "../cghttp.h"

#pragma comment(lib, "../http.lib")

int main()
{
    char* body = NULL;
    size_t bodylen = 0;
    int ret = Get("http://www.gongluck.icu/", &body, &bodylen);
    printf("%d\n%s\n%zd\n", ret, body, bodylen);
    return 0;
}