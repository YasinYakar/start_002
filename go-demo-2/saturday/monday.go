package saturday

import "fmt"
import "time"

func monday() {

today:= time.Now().Weekday()
switch time.Monday {

case today+0:{fmt.Println("Today")}
case today+1:{fmt.Println("Tomorrow")}
case today+2:{fmt.Println("in two days")}
default : {fmt.Println("too far away")}
}




}