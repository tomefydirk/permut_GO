package main

import(
	"fmt"
)
func is_inV1(cycle []int,a int) bool{
	for j:=0;j<len(cycle);j++{
		if cycle[j]==a{
			return true;
		}
	}
	return false;
}
func is_inV2(cycle [][]int,a int ) bool{
	for i:=0;i<len(cycle);i++ {
		for j:=0;j<len(cycle[i]);j++{
			if cycle[i][j]==a{
				return true;
			}
		}
	}
	return false;
}

func find_another_int(my_sigma1 []int ,cycle [][]int) int {
	for i:=0; i<len (my_sigma1); i++{
		
			if !is_inV2(cycle,my_sigma1[i]){
				return my_sigma1[i];
			}
		
	}
	return -1;
}
func find_image(my_sigma [2][]int, value int) int {
	 return my_sigma[1][value];
}
func push_suivant(my_sigma [2][]int,un_cycle []int ) int{
	a:=un_cycle[len(un_cycle)-1];
		if a==0{
			
		}
		return -1;
}
func chercher_cycle(my_tab [][]int) [][]int{
	
	return my_tab;
}
func main(){
	 fmt.Println("done");	
}