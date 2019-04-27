package utils

/***
 * 变量交换
 */
func Swap(list []int, i, j int)  {
	tmp := list[i]
	list[i] = list[j]
	list[j] = tmp
}
/***
 * go特有变量交换
 */
func SwapGo(list []int, i, j int)  {
	list[i],list[j]=list[j],list[i]
}
/***
 * go变量高阶交换(不推荐，一般不好理解)
 */
func SwapGoAdvanced(list []int, i, j int)  {
	list[i]=list[i]+list[j]
	list[j]=list[i]-list[j]
	list[i]=list[i]-list[j]
}
