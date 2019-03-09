package main

import "fmt"

type Books struct {
	title string
	id int
}

func showTile(b Books) {
	fmt.Println("title of this booK: ", b.title);
}

type Phone interface {
	call()
}

type Huawei struct {
}

func (hw Huawei) call() {
	fmt.Println("this is a Huawei phone")
}

func main() {
   fmt.Println("Hello, World!")

   // function
   fmt.Printf("\nfunction: \n")   
   var result int
   result = getMax(4, 10)
   fmt.Println("max number: ", result)

   x, y := swap("tong", "peng")
   fmt.Println("after swapping: ", x, y)

   // array
   fmt.Printf("\narray: \n")
   var nums1 [5]int
   var nums2 = [5]int{0, 1, 2, 3, 4}
   fmt.Println(nums1)
   fmt.Println(nums2)

   for i,val := range nums2 {
   		fmt.Printf("num at %d is %d\n", i, val)
   }

   // 2d array
   var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
   var i, j int

   /* 输出数组元素 */
   for  i = 0; i < 5; i++ {
      for j = 0; j < 2; j++ {
         fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j] )
      }
   }

   // pointer
   fmt.Printf("\npointer: \n")
   var ipTest int = 22
   var ip *int
   ip = &ipTest
   fmt.Println("&ipTest: ", &ipTest)
   fmt.Println("ip: ", ip)
   fmt.Println("*ip: ", *ip)

   // struct
   fmt.Printf("\nstruct: \n")

   book1 := Books{"math", 1}
   var book2 Books = Books{"nba", 2}
   var book3 Books
   book3 = Books{"java", 3}

   fmt.Println(book1)
   fmt.Println(book2)
   fmt.Println(book3)   

   var bookIp *Books = &book2
   fmt.Println("ip book title: ", bookIp.title)

   // slice
   fmt.Printf("\nslice: \n")   
   slice1 := []int{1, 2, 3}
   slice2 := make([]int, 5, 10)
   fmt.Println("slice1: ", slice1)
   fmt.Println("slice2: ", slice2)

   slice2 = append(slice2, 1, 2, 3)
   fmt.Println("slice2: ", slice2)
   fmt.Println("slice2 part: ", slice2[:6])

   // map
   fmt.Printf("\nmap: \n")
   cmap := make(map[string]string)
   cmap["China"] = "Beijing"
   cmap["US"] = "Washinton"
   cmap["UK"] = "London"

   for key := range cmap {
   		fmt.Println("country: ", key, ",capital: ", cmap[key])		
   }

   delete(cmap, "UK")

   if _,ok := cmap["UK"]; !ok {
   		fmt.Println("UK is not in the map")
   }

   if value,ok := cmap["China"]; ok {
   		fmt.Println("China is in the map, ", value)
   }

   // interface
   fmt.Printf("\ninterface: \n")   
   var hw Phone = new(Huawei)
   hw.call()
}


func getMax(a, b int) (int) {
	if a > b {
		return a
	} else {
		return b
	}
} 

func swap(x string, y string) (string, string) {
	return y, x
}