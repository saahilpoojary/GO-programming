package main
import "fmt"
func pythtriplet(n int)(int,int,int){
    for a:=1;a<=n/3;a++ {
        for b:=a;b<=n/2;b++{
            c:=n-(a+b)
            if a*a + b*b== c*c{
                return a,b,c
            }
            
        }
    }
    return -1,-1,-1
}
func main(){
    var n int
    fmt.Println("Enter the number")
    fmt.Scan(&n)
    a,b,c:=pythtriplet(n)
    if a!=-1 && b!=-1 && c!=-1{
        fmt.Printf("The pythagorean triplet for sum %d is %d %d and %d\n",n,a,b,c)
    }else{
        fmt.Printf("There is no pythagorean triplet for %d\n",n)
    }
    
}
