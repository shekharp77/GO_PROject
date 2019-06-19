package main

import (
	"GO_PROject/DB"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	printHi()
	startProcess()
}

var (
	projectname string
)

func printHi() {
	fmt.Println("###############################")
	fmt.Println("#                             #")
	fmt.Println("#    WELCOME TO GoPROject     #")
	fmt.Println("#                             #")
	fmt.Println("###############################")
}

func createProjectDirectory() {
	os.Mkdir(projectname, 0700)
}

func createMainFile() {
	path := fmt.Sprintf("%s/%s.go",getProjectDirectory(),projectname)
	err3 := ioutil.WriteFile(path, []byte(""), 0700)
	if err3 != nil {
		fmt.Println("Oh No!  its an error ;(")
		fmt.Println(err3)
		return
	}
}

func createDB() {
	os.Mkdir(getProjectDirectory() + "/" + "Database", 0700)
}

func getCurrentPath() (string, error) {
	dir, err := os.Getwd()
	return dir, err
}

func getProjectDirectory() string{
	dir, err2 := getCurrentPath()
	projectDir := dir + "/" + projectname
	if err2 != nil {
		fmt.Println("Oh No!  its an error ;(")
		fmt.Println(err2)
		return ""
	}
	return projectDir
}

func Execute(url string) {
	out, err := exec.Command("go","get", url).Output()
	fmt.Println(out)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	fmt.Println("Command Successfully Executed")
}

func startProcess() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Project Name ")
	projectname, _ = reader.ReadString('\n')
	projectname = strings.Replace(projectname, "\n", "", -1)


	createProjectDirectory()
	createMainFile()

	fmt.Print("Do you want to use database?(y/n)")
	dbans, _ := reader.ReadString('\n')
	dbans = strings.Replace(dbans, "\n", "", -1)

	if dbans == "y" {
		createDB()
		DB.InitDBoptions()
		fmt.Print("which db you want to use?")
		fmt.Println(DB.DBmap)
		db, _ := reader.ReadString('\n')
		db = strings.Replace(db, "\n", "", -1)
		switch db {
		case "0": installMySql()
		case "1": installPostgres()
		}
	}
}


func installMySql(){
	Execute("github.com/go-sql-driver/mysql")
	path := fmt.Sprintf("%s/Database/DBConnection.go",getProjectDirectory())
	file, err := os.Create(path)
	file.Chmod(777)
	defer file.Close()
	if err != nil {
		fmt.Println("Oh No!  its an error ;(")
		fmt.Println(err)
		return
	}

	data, err := ioutil.ReadFile(getProjectDirectory()+"/../DB/mysql.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	l, err := file.WriteString(string(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Close()
}



func installPostgres(){
	Execute("github.com/lib/pq")
	path := fmt.Sprintf("%s/Database/DBConnection.go",getProjectDirectory())
	file, err := os.Create(path)
	file.Chmod(777)
	defer file.Close()
	if err != nil {
		fmt.Println("Oh No!  its an error ;(")
		fmt.Println(err)
		return
	}

	data, err := ioutil.ReadFile(getProjectDirectory()+"/../DB/postgres.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	l, err := file.WriteString(string(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Close()
}
