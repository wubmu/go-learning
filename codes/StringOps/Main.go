package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
   	需要掌握的常用的string的相关函数
   		1、len()用来统计一个字符串的长度的（按字节统计一个中文字符在utf-8中是三个字节）
   		2、[]rune(T type) 将一个字符串转化为切片，可以用于遍历含有中文的字符
   		3、i,err := strconv.Atoi()将一个字符串转化为一个整数，返回值有i和error两个
   		4、str := strconv.ItoA() 将一个整数转化为一个字符串
   		5、bytes := []byte(str string) 字符串转化为byte切片
   		6、str := string([]byte{97,98,99}) byte切片转换为字符串
    		7、str := strconv.FormatInt(i int,base b) 将十进制转化为其他进制（2到32进制）
   		8、var judge bool = strings.Contains(BigStr string,littleStr string) 判断一个字符串是否包含另一个字符串
   		9、 var counts int = strings.Count(s, str string) 返回一个字符串有几个指定的字符串
   		10、var judge bool = string.EqualFold("abc","Abc") 不区分大小写的字符串比较，区分的使用 == 就可以
   		11、var index int = strings.Index(s,str string) 寻找字符串第一次出现的位置，没有的话返回-1
   		12、var lastIndex int = strings.LastIndex(s,str) 寻找一个字符在字符串中最后一次返回的位置
   		13、var str string = strings.Replace(s,oldstr,newstr,n)将一个字符串中的字符串替换成另一个字符串，n表示替换几个，
               如果是-1表示替换所有
   		14、stringArr string[] = strings.Split(s,str) 将一个字符串按照指的的分隔符分割为一个字符串数组
   		15、strings.ToLower(s)、stings.ToUpper(s) 将字符串进行大小写的变化
   		16、strings.TrimSpace(s) 去掉字符串两边的空格
   		17、strings.Trim(s，str) 去掉字符串两边的指定字符串（str包括指定字符串和空格）
   		18、strings.TimLeft(s,str) 、strings.TimRight(s,str)去除左右两边指定的字符串，顺便去掉多出的空格
   		19、strings.HasPrefix(s,str)、strings.HasSuffix(s,str)判断一个字符串是否是以指定字符串开头或者结尾的

*/

func main() {
	//定义变量
	var s string
	var str string
	var count int
	var err error
	var bytes []byte
	var judge bool
	//1、len()用来统计一个字符串的长度的（按字节统计一个中文字符在utf-8中是三个字节）
	s = "hello海 "
	fmt.Println(s, " len = ", len(s))
	//2、[]rune(T type) 将一个字符串转化为切片，可以用于遍历含有中文的字符
	s = "晋太元中，武陵人捕鱼为业"
	var runes []rune = []rune(s)
	fmt.Printf("runes = %q\n", runes)
	//3、i,err := strconv.Atoi()将一个字符串转化为一个整数，返回值有i和error两个
	s = "12345789a"
	count, err = strconv.Atoi(s)
	fmt.Printf("count = %d\n", count)
	fmt.Println("err =", err)
	//4、str := strconv.Itoa() 将一个整数转化为一个字符串
	count = 123456
	str = strconv.Itoa(count)
	fmt.Println("str = ", str)
	//5、bytes := []byte(str string) 字符串转化为byte切片
	s = "hello "
	bytes = []byte(s)
	fmt.Printf("bytes = %c\n", bytes)
	//6、str := string([]byte{97,98,99}) byte切片转换为字符串
	str = string([]byte{97, 98, 99})
	fmt.Println("str = ", str)
	//7、str := strconv.FormatInt(i int,base b) 将十进制转化为其他进制（2到32进制）
	formatInt := strconv.FormatInt(100, 2)
	fmt.Println("FormatInt = ", formatInt)
	//8、var judge bool = strings.Contains(BigStr string,littleStr string) 判断一个字符串是否包含另一个字符串
	s = "鹅鹅鹅，曲项向天歌"
	str = "鹅"
	judge = strings.Contains(s, str)
	fmt.Println("字符串\"", s, "\"中包含\"", str, "\"否？", judge)
	//9、 var counts int = strings.Count(s, str string) 返回一个字符串有几个指定的字符串
	s = "鹅鹅鹅，曲项向天歌"
	str = "鹅"
	count = strings.Count(s, str)
	fmt.Printf("字符串%s中包含个%d字符串%s\n", s, count, str)
	//10、var judge bool = string.EqualFold("abc","Abc") 不区分大小写的字符串比较，区分的使用 == 就可以
	s = "keysHe"
	str = "keyshe"
	judge = strings.EqualFold(s, str)
	fmt.Printf("字符串%s和字符串%s相等否%v\n", s, str, judge)
	//11、var index int = strings.Index(s,str string) 寻找字符串第一次出现的位置，没有的话返回-1
	s = "蚕丛及鱼凫，开国何茫然 何茫然"
	str = "何茫然"
	count = strings.Index(s, str)
	fmt.Printf("字符串%s在字符串%s第一次出现的位置是%d\n", s, str, count)
	//12、var lastIndex int = strings.LastIndex(s,str) 寻找一个字符在字符串中最后一次返回的位置
	s = "蚕丛及鱼凫，开国何茫然 何茫然"
	str = "何茫然"
	count = strings.LastIndex(s, str)
	fmt.Printf("字符串%s在字符串%s最后一次出现位置是%d\n", s, str, count)
	//13、var str string = strings.Replace(s,oldstr,newstr,n)将一个字符串中的字符串替换成另一个字符串，n表示替换几个，
	//如果是-1表示替换所有
	s = "鹅鹅鹅，曲项向天歌，白毛浮绿水，红掌拨清波"
	oldstr := "鹅"
	newstr := "骆宾王"
	s = strings.Replace(s, oldstr, newstr, 4)
	fmt.Println("s = ", s)
	//14、stringArr []string = strings.Split(s,str) 将一个字符串按照指的的分隔符分割为一个字符串数组
	s = "西当太白有鸟道，可以横绝峨眉巅"
	str = "，"
	var split = strings.Split(s, str)
	fmt.Println("split = ", split)
	//15、strings.ToLower(s)、stings.ToUpper(s) 将字符串进行大小写的变化
	s = "hello"
	s = strings.ToUpper(s)
	fmt.Println("s =", s)
	//16、strings.TrimSpace(s) 去掉字符串两边的空格
	s = " hello "
	s = strings.TrimSpace(s)
	fmt.Println("s =", s)
	//17、strings.Trim(s，str) 去掉字符串两边的指定字符串（str包括指定字符串和空格）
	s = " hello ?"
	s = strings.Trim(s, "?")
	fmt.Println("s =", s)
	//18、strings.TimLeft(s,str) 、strings.TimRight(s,str)去除左右两边指定的字符串，顺便去掉多出的空格
	s = " hello ?"
	s = strings.TrimRight(s, "?")
	fmt.Println("s =", s)
	//19、strings.HasPrefix(s,str)、strings.HasSuffix(s,str)判断一个字符串是否是以指定字符串开头或者结尾的
	s = " hello ?"
	judge = strings.HasSuffix(s, "?")
	fmt.Println("judge =", judge)
}
