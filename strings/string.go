//（1）func HasPrefix(s, prefix string) bool
//判断字符串s是否有前缀字符串prefix；
//
//（2）func HasSuffix(s, suffix string) bool
//判断字符串s是否有后缀字符串suffix；
//
//（3）func Contains(s, substr string) bool
//判断字符串s是否包含子串substr；
//
//（4）func Count(s, sep string) int
//返回字符串s有几个重复sep子串；
//
//（5）func Index(s, sep string) int
//返回字符串s中子串sep第一次出现的位置；不存在返回-1；
//
//（6）func ToLower(s string) string
//返回字符串s转小写的拷贝；
//
//（7）func ToUpper(s string) string
//返回字符串s转大写的拷贝；
//
//（8）func Repeat(s string, count int) string
//返回count个字符串s串联的字符串；
//
//（9）func Replace(s, old, new string, n int) string
//返回字符串s前n个不重复old子串替换为new子串的新字符串；n<0替换所有old子串；
//
//（10）func Trim(s string, cutset string) string
//返回去掉字符串s前后端所有cutset子串的字符串；
//
//（11）func TrimSpace(s string) string
//返回去掉字符串s前后端空白字符（unicode.IsSpace指定）的字符串；
//
//（12）func TrimLeft(s string, cutset string) string
//返回去掉字符串s前端所有cutset子串的字符串；
//
//（13）func TrimRight(s string, cutset string) string
//返回去掉字符串s后端所有cutset子串的字符串；
//
//（14）func TrimPrefix(s, prefix string) string
//返回去掉字符串s的前缀prefix子串的字符串；
//
//（15）func TrimSuffix(s, suffix string) string
//返回去掉字符串s的后缀suffix子串的字符串；
//
//（16）func Fields(s string) []string
//返回将字符串s按一个或多个空白（unicode.IsSpace）字符分割的多个字符串切片；空白字符串或空字符串返回空切片；
//
//（17）func Split(s, sep string) []string
//返回将字符串s按一个sep子串分割的字符串切片；sep为空字符串时，将s分割为每一个unicode码值的字符串切片；
//
//（18）func Join(a []string, sep string) string
//返回将字符串切片a以子串sep连接的字符串；
//
//（19）func NewReader(s string) *Reader
//创建从字符串s读取数据的Reader指针；

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("查找子串是否在指定的字符串中")
	fmt.Println(" Contains 函数的用法")
	fmt.Println(strings.Contains("seafood", "foo")) //true
	fmt.Println(strings.Contains("seafood", "bar")) //false
	fmt.Println(strings.Contains("seafood", ""))    //true
	fmt.Println(strings.Contains("", ""))           //true 这里要特别注意
	fmt.Println(strings.Contains("我是中国人", "我"))     //true

	fmt.Println("")
	fmt.Println(" ContainsAny 函数的用法")
	fmt.Println(strings.ContainsAny("team", "i"))        // false
	fmt.Println(strings.ContainsAny("failure", "u & i")) // true
	fmt.Println(strings.ContainsAny("foo", ""))          // false
	fmt.Println(strings.ContainsAny("", ""))             // false

	fmt.Println("")
	fmt.Println(" ContainsRune 函数的用法")
	fmt.Println(strings.ContainsRune("我是中国", '我')) // true 注意第二个参数，用的是字符

	fmt.Println("")
	fmt.Println(" Count 函数的用法")
	fmt.Println(strings.Count("cheese", "e")) // 3
	fmt.Println(strings.Count("five", ""))    // before & after each rune result: 5 , 源码中有实现 substr "" utf8.RuneCountInString(s) + 1

	fmt.Println("")
	fmt.Println(" EqualFold 函数的用法")
	fmt.Println(strings.EqualFold("Go", "go")) //大小写忽略

	fmt.Println("")
	fmt.Println(" Fields 函数的用法")
	fmt.Println("Fields are:", strings.Fields("  foo bar  baz   "))
	//["foo" "bar" "baz"] 返回一个列表

	//相当于用函数做为参数，支持匿名函数
	for _, record := range []string{" aaa*1892*122", "aaa\taa\t", "124|939|22"} {
		fmt.Println(strings.FieldsFunc(record, func(ch rune) bool {
			switch {
			case ch > '5': // ch > '5' 作为分开 string 的标志
				return true
			}
			return false
		}))
	}

	fmt.Println("")
	fmt.Println(" HasPrefix 函数的用法")
	fmt.Println(strings.HasPrefix("NLT_abc", "NLT")) //前缀是以NLT开头的

	fmt.Println("")
	fmt.Println(" HasSuffix 函数的用法")
	fmt.Println(strings.HasSuffix("NLT_abc", "abc")) //后缀是以abc开头的

	fmt.Println("")
	fmt.Println(" Index 函数的用法")
	fmt.Println(strings.Index("NLT_abc", "abc")) // 返回第一个匹配字符的位置，这里是4
	fmt.Println(strings.Index("NLT_abc", "aaa")) // 不存在返回 -1
	fmt.Println(strings.Index("我是中国人", "中"))     // 在存在返回 6

	fmt.Println("")
	fmt.Println(" IndexAny 函数的用法")
	fmt.Println(strings.IndexAny("我是中国人", "中")) // 在存在返回 6
	fmt.Println(strings.IndexAny("我是中国人", "和")) // 在存在返回 -1

	fmt.Println("")
	fmt.Println(" Index 函数的用法")
	fmt.Println(strings.IndexRune("NLT_abc", 'b')) // 返回第一个匹配字符的位置，这里是4
	fmt.Println(strings.IndexRune("NLT_abc", 's')) // 在存在返回 -1
	fmt.Println(strings.IndexRune("我是中国人", '中'))   // 在存在返回 6

	fmt.Println("")
	fmt.Println(" Join 函数的用法")
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", ")) // 返回字符串：foo, bar, baz

	fmt.Println("")
	fmt.Println(" LastIndex 函数的用法")
	fmt.Println(strings.LastIndex("go gopher", "go")) // 3

	fmt.Println("")
	fmt.Println(" LastIndexAny 函数的用法")
	fmt.Println(strings.LastIndexAny("go gopher", "go")) // 4
	fmt.Println(strings.LastIndexAny("我是中国人", "中"))      // 6

	fmt.Println("")
	fmt.Println(" Map 函数的用法")
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))

	fmt.Println("")
	fmt.Println(" Repeat 函数的用法")
	fmt.Println("ba" + strings.Repeat("na", 2)) //banana

	fmt.Println("")
	fmt.Println(" Replace 函数的用法")
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))

	fmt.Println("")
	fmt.Println(" Split 函数的用法")
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))

	fmt.Println("")
	fmt.Println(" SplitAfter 函数的用法")
	fmt.Printf("%q\n", strings.SplitAfter("/home/m_ta/src", "/")) //["/" "home/" "m_ta/" "src"]

	fmt.Println("")
	fmt.Println(" SplitAfterN 函数的用法")
	fmt.Printf("%q\n", strings.SplitAfterN("/home/m_ta/src", "/", 2))  //["/" "home/m_ta/src"]
	fmt.Printf("%q\n", strings.SplitAfterN("#home#m_ta#src", "#", -1)) //["/" "home/" "m_ta/" "src"]

	fmt.Println("")
	fmt.Println(" SplitN 函数的用法")
	fmt.Printf("%q\n", strings.SplitN("/home/m_ta/src", "/", 1))

	fmt.Printf("%q\n", strings.SplitN("/home/m_ta/src", "/", 2))  //["/" "home/" "m_ta/" "src"]
	fmt.Printf("%q\n", strings.SplitN("/home/m_ta/src", "/", -1)) //["" "home" "m_ta" "src"]
	fmt.Printf("%q\n", strings.SplitN("home,m_ta,src", ",", 2))   //["/" "home/" "m_ta/" "src"]

	fmt.Printf("%q\n", strings.SplitN("#home#m_ta#src", "#", -1)) //["/" "home/" "m_ta/" "src"]

	fmt.Println("")
	fmt.Println(" Title 函数的用法") //这个函数，还真不知道有什么用
	fmt.Println(strings.Title("her royal highness"))

	fmt.Println("")
	fmt.Println(" ToLower 函数的用法")
	fmt.Println(strings.ToLower("Gopher")) //gopher

	fmt.Println("")
	fmt.Println(" ToLowerSpecial 函数的用法")

	fmt.Println("")
	fmt.Println(" ToTitle 函数的用法")
	fmt.Println(strings.ToTitle("loud noises"))
	fmt.Println(strings.ToTitle("loud 中国"))

	fmt.Println("")
	fmt.Println(" Replace 函数的用法")
	fmt.Println(strings.Replace("ABAACEDF", "A", "a", 2)) // aBaACEDF
	//第四个参数小于0，表示所有的都替换， 可以看下golang的文档
	fmt.Println(strings.Replace("ABAACEDF", "A", "a", -1)) // aBaaCEDF

	fmt.Println("")
	fmt.Println(" ToUpper 函数的用法")
	fmt.Println(strings.ToUpper("Gopher")) //GOPHER

	fmt.Println("")
	fmt.Println(" Trim  函数的用法")
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! ")) // ["Achtung"]

	fmt.Println("")
	fmt.Println(" TrimLeft 函数的用法")
	fmt.Printf("[%q]", strings.TrimLeft(" !!! Achtung !!! ", "! ")) // ["Achtung !!! "]

	fmt.Println("")
	fmt.Println(" TrimSpace 函数的用法")
	fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n")) // a lone gopher

}
