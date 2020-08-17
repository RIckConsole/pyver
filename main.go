package main

import (
       "fmt"
       "os"
       "os/exec"
       "log"
)

func main() {
       usage := "Usage: pyver <set/check/add> <version>"
       if (len(os.Args) < 2) {
              fmt.Println(usage)
              os.Exit(0)
       }
       switch os.Args[1]{
              case "check":
              checkver()
              case "set":
              switchver(os.Args[2])
              case "add":
              addver(os.Args[2])
              default:
              fmt.Println(usage)
       }

}

func checkver(){
       cmd, err := exec.Command("bash", "-c", "update-alternatives --get-selections | grep python | grep 'python.*python' | awk '{print $3}'").Output()
       //cmd, err := exec.Command("python", "--version").Output()
       if err != nil {
              log.Fatal(err)
       }
       fmt.Print(string(cmd))
}

func switchver(version string){
       path := "/usr/bin/"
       cmd := exec.Command("update-alternatives", "--set", "python", path+version)

       _, err := cmd.Output()
       if err == nil {
              fmt.Println("Python version switched to:", version)
       } else {
              fmt.Println("Are you r00t? Or maybe you havent added that version yet...")
       }
}

func addver(version string){
       path := "/usr/bin/"

       cmd := exec.Command("update-alternatives", "--install", "/usr/bin/python", "python", path+version, "1")
       stdout, err := cmd.Output()
       if err == nil {
              fmt.Println("Python version added successfully!")
       } else {
              fmt.Print(stdout)
       }
}
