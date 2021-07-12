 ### Installation
 ```go get github.com/oguzcansaribiyik/struct-comparison```
 
 ### Example Usage
  ```
    package main

    import (
      "fmt"
      comparison "github.com/oguzcansaribiyik/struct-comparison"
    ) 

    type User struct {
      Name string
      Job  Job
      Age  int
    }

    type Job struct {
      Name string
    }

    func main() {
      oldStruct := User{
        Name: "Oğuzcan",
        Job: Job{
          Name: "Student",
        },
        Age: 19,
      }

      newStruct := User{
        Name: "Oğuzcan",
        Job: Job{
          Name: "Software Developer",
        },
        Age: 23,
      }

      var response []comparison.Response
      comparison.Compare(oldStruct, newStruct, &response)
      for _, v := range response {
        fmt.Printf("Field : %v - Old Value : %v - New Value : %v\n", v.FieldName, v.OldValue, v.NewValue)
        fmt.Println("-----------------")
      }
    }
  ```
 ### Console Output
  ```
  Field : Job.Name - Old Value : Student - New Value : Software Developer
  -----------------
  Field : User.Age - Old Value : 19 - New Value : 23
  -----------------
  ```