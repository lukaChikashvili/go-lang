package main
import "fmt"

func main() {

	// define variables
     var firstName string
	 var lastName string
	 var email string
	 var userTickets uint

	 var remainingTickets uint = 50
	 var bookings [50]string


// ask user to enter info
  fmt.Println("Enter your first name")
  fmt.Scan(&firstName)

  fmt.Println("Enter your last name")
  fmt.Scan(&lastName)

  fmt.Println("Enter your email ")
  fmt.Scan(&email)
  
  fmt.Println("Enter your tickets ")
  fmt.Scan(&userTickets)

  remainingTickets = remainingTickets - userTickets
   bookings[0] = firstName + " " + lastName

   fmt.Println("thank you", firstName, " ", lastName, "i will write on email", email)
   fmt.Println(remainingTickets, "tickets are remained")
   fmt.Println(bookings)
   

 





	
	 
}
