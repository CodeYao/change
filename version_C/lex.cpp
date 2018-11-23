#include < stdio.h >
#define BUFLEN 80  //缓冲区大小
int lineLen = 0;   //缓冲区内的数据长度
int readPos = -1;  //读取位置
char line[BUFLEN]; //缓冲区
int lineNum = 1;   //行号
int colNum = 0;    //列号
char lastch;       //上一个字符
char scan(FILE *file)
{
    if (!file) //没有文件
    {
        return -1;
    }
    if (readPos == lineLen - 1) //缓冲区读取完毕/缓冲区为空
    {
        lineLen = fread(line, 1, BUFLEN, file); //重新加载缓冲区数据
        if (lineLen == 0)                       //没有数据了
        {
            lineLen = 1;  //数据长度设置为1
            line[0] = -1; //文件结束标记
        }
        readPos = -1; //恢复读取位置
    }
    readPos++;               //移动读取点
    char ch = line[readPos]; //获取新的字符
    if (lastch == "\n")      //新行
    {
        lineNum++; //行号增加
        colNum=0;  //列号清零
    }
    if (ch == -1) //文件结束，自动关闭
    {
        fclose(file);
        file = NULL
    }
    else if (ch != '\n') //不是新行
    {
        colNum++; //列号递增
    }
    lastch = ch; //记录上个字符
    return ch;   //返回当前字符串
}