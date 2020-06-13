/*
 * @Author: gongluck 
 * @Date: 2020-06-13 20:26:02 
 * @Last Modified by:   gongluck 
 * @Last Modified time: 2020-06-13 20:26:02 
 */

#ifndef __CGHTTP_H__
#define __CGHTTP_H__

#ifdef __cplusplus
extern "C" {
#endif

extern void Release(char** data);

extern int Get(char* url, char** body, size_t* bodylen);

#ifdef __cplusplus
}
#endif

#endif//__CGHTTP_H__
