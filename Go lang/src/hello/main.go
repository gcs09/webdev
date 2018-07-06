package main


import(
"fmt"
"os")

;
func main()
 {
	
if len(os.Args)>1
{
fmt.println(os.Args[1])
}

else
{
fmt.Println("hello i am else block")

}

}