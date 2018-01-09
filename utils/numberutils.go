package utils
//func GetSum(x,y int) int {
//	return x+y
//}
func GetSum(x,y int ) (int ,string){
	if x<0 || y<0 {
		return -1,"不能小于0"
	}else {
		return x+y,""
	}
}
func GetSplit(sum int) (x,y int){
	x =sum*3/9
	y=sum-9
	return
}
func GetAdress(adr int) *int{
	return  &adr
}
