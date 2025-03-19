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
func push_suivant(my_sigma [2][]int,un_cycle []int,est_finie *bool) {
	x:=un_cycle[len(un_cycle)-1];
	y:=find_image(my_sigma,x);
	fmt.Println("y",y);
		if is_inV1(un_cycle,y){
			 *est_finie=true;	
			 fmt.Println("un_cycle:",un_cycle);
			 return;
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
	cycle :=[][] int{
		{my_sigma[0][0]},
	} 
	est_finie:=false;
	i:=0;
	for !all_is_in(my_sigma[0],cycle){
	if est_finie{
		est_finie=false;
		i++;
		x:=find_another_int(my_sigma[0],cycle);
		if x==-1{
				break;
		}
		fmt.Println("on a x=",x);
		temp:=append(cycle[i],x);
		cycle=append(cycle,temp);
	}	
	push_suivant(my_sigma,cycle[i],&est_finie);		
	}
	return my_sigma;
}

// }
func main(){
	permutation := [2][]int{
        {1, 2, 3, 4, 5}, // Les éléments originaux
        {3, 5, 1, 2, 4}, // L'image des éléments après la permutation
    }
	permutation1 := [2][]int{
         {2}, // Les éléments originaux
    }
	est_finie:=false
	push_suivant(permutation,permutation1,est_finie);
	fmt.Println(find_image(permutation,1));
	
//	fmt.Println(a);

	 fmt.Println("done");	

}