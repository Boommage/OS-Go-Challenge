package main

import (
"fmt"
	"bufio"
	"strings"
	"strconv"
	"log"
	"os"
)

//error handler for redundancy
func errc(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

//global map variable
var kv = map[int64]string{}

//Checks if key is valid
func keyCheck(str string, key *int64) bool {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Invald key")
		return false
	}
	*key = int64(num)
	return true
}

//places or replaces (key,value) pair into map
func put(str []string) {
	var key int64
	if len(str) != 2 {
		fmt.Println("bad command")
		return
	}
	if !(keyCheck(str[0], &key)) {
		return
	}
	kv[key] = str[1]
}

//prints (key,value) pair from map based on key
func get(str []string) {
	var key int64
	if len(str) != 1 {
		fmt.Println("bad command")
		return
	}
	if !(keyCheck(str[0], &key)) {
		return
	}
	val := kv[key]
	if val != "" {
		fmt.Printf("%d,%s\n",key,val)
		return
	}
	fmt.Println(key,"not found")
}

//removes (key,value) pair from map based on key
func del(str []string) {
	var key int64
	if len(str) != 1 {
		fmt.Println("bad command")
		return
	}
	if !(keyCheck(str[0], &key)) {
		return
	}
	val := kv[key]
	if val != "" {
		delete(kv, key)
		return
	}
	fmt.Println(key,"not found")
}

//prints all (key,value) pairs present in the map
func all() {
	for key, val := range kv {
		fmt.Printf("%d,%s\n",key,val)
	} 
}

//loads (key,value) pair data from local file into map
func readMap() {
	file, err := os.Open("mapFile.txt")
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		put(strings.Split(scanner.Text(),","))
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner encountered an error: %s", err)
	}
}

//saves (key,value) pair data from map to a local file 
func writeMap() {
	file, err := os.Create("mapFile.txt")
	errc(err)
	defer file.Close()
	for key, val := range kv {
	message := fmt.Sprintf("%d,%s\n",key,val)
	_, err = file.WriteString(message)
	errc(err)
	}
}

func main() {
	//persistance - loading data
	readMap()

	if len(os.Args) < 1 {
		os.Exit(0)
	}

	//for each argument - does not end until each argument is read
	for _, arg := range os.Args[1:] {
		cmd := strings.Split(arg,",")

		//Switch statement to check cmd (p, g, d, c, or a)
		switch cmd[0] {
		case"p":
			put(cmd[1:])
		case "g":
			get(cmd[1:])
		case "d":
			del(cmd[1:])
		case "c":
			if len(cmd) > 1 {
				fmt.Println("bad command")
				continue
			}
			clear(kv)
		case "a":
			all()
		default:
			fmt.Println("bad command")
		}
	}
	//persistance - saving data
	writeMap()	
}
