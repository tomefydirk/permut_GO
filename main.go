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
	 return my_sigma[1][value-1];
}

// to_debug{
func push_suivant(my_sigma [2][]int,un_cycle *[]int) {
	x:=(*un_cycle)[len(*un_cycle)-1];
	y:=find_image(my_sigma,x);


	for  !is_inV1(*un_cycle,y){
		*un_cycle=append(*un_cycle,y);
		x=(*un_cycle)[len(*un_cycle)-1];
	    y=find_image(my_sigma,x);
		
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
func chercher_cycle(my_sigma [2][]int) [][]int{
	cycle :=[][] int{
		
	} 
	for i:=0;!all_is_in(my_sigma[0],cycle);i++{
	a:=find_another_int(my_sigma[0],cycle);
	cycle=append(cycle,[]int{a});
	push_suivant(my_sigma,&cycle[i]);	
	}
	return cycle;
}

// }
func main(){
	permutation := [2][]int{
        {1, 2, 3, 4, 5,6,7,8,9}, // Les éléments originaux
        {9, 1, 4, 7, 3,8,6,5,2}, // L'image des éléments après la permutation
    }
	//push_suivant(permutation,&permutation1[0]);
	//fmt.Println(permutation1);
//	fmt.Println(find_image(permutation,1));
	fmt.Println(chercher_cycle(permutation));

	fmt.Println(find_another_int(permutation[0],[][]int{{1,2},}));
//	fmt.Println(a);

	 fmt.Println("done");	

}