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
func push_suivant(my_sigma [2][]int,un_cycle []int,est_finie *bool) {
	x:=un_cycle[len(un_cycle)-1];
	y:=find_image(my_sigma,x);
		if is_inV1(un_cycle,y){
			 *est_finie=true;	
		}else{
			un_cycle=append(un_cycle,y);	
		}
		
}

func all_is_in(my_sigma1 []int,cycle [][]int) bool {
	for i:=0;i<len(my_sigma1);i++{
		if !is_inV2((cycle),my_sigma1[i]) {
			return false;
		}
	}
	return true;
}
func chercher_cycle(my_sigma [2][]int) [2][]int{
	var cycle [][] int; 
	est_finie:=false;
	for !all_is_in(my_sigma[0],cycle){
	if !est_finie{
		est_finie=true;
	}		
	}
	return my_sigma;
}
func main(){
	 fmt.Println("done");	
}