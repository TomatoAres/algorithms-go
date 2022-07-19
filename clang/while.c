#include <stdio.h>
 
int main ()
{
   /* 局部变量定义 */
   int a = 10;

   /* while 循环执行 */
   while( a-- )
   {
      printf("a 的值： %d\n", a);
    //   a++;
   }
 
   return 0;
}