package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)


type Directory struct{
	parent        *Directory
	child_dirs    []Directory
	child_files   []Normalfile
	name          string
	fullpath      string
	size          int
}

type Normalfile struct{
	parent    *Directory
	size      int
	name      string
	fullpath  string
}

//var pwd string
var root = Directory{name: "/", fullpath: "/"};
var pwd  = &root

func populate_curr_pwd(line string){
	if (line[0:3] == "dir"){
		fmt.Printf("\t\t DIR, %s", line)

		v := strings.Split(line, " ") //
		if len(v) != 2{
			fmt.Printf("err: check this line %s %v", line, v)
			panic(3)
		}
		dir_name := v[1]
		tmpdir := Directory{
			parent        :pwd,
			name          :dir_name,
			fullpath      :pwd.fullpath  + dir_name + "/",
		}
		pwd.child_dirs = append(pwd.child_dirs, tmpdir)

	} else{
		v := strings.Split(line, " ") //
		if len(v) != 2{
			fmt.Printf("err: check this line %s %v", line, v)
			panic(3)
		}
		file_size, err := strconv.Atoi(v[0])
		if err != nil {
			fmt.Printf("err: while converting str to int %v, %v", v, err)
			panic(4)
		}
		file_name := v[1]
		tmpfile := Normalfile{
			parent: pwd,
			size: file_size,
			name: file_name,
			fullpath: pwd.fullpath + file_name,
		}
		//appending to curren working directory
		pwd.child_files = append(pwd.child_files, tmpfile)
		
		fmt.Printf("\t\t FILE, %s, %v", line, v)

		//var tmp Normalfile {}
	}
}

func tree(printpwd *Directory){ //prints out everything from root
	//ah sheeet this can be recursive hha
	//printpwd = &root
	//pritn out dirs wiith recursive
	for i,L := 0, len(printpwd.child_dirs); i<L; i++{
		fmt.Printf("-d- %d\t%s\n",
			printpwd.child_dirs[i].size,
			printpwd.child_dirs[i].fullpath,
		)
		tree(&printpwd.child_dirs[i])
	}
	for i,L := 0, len(printpwd.child_files); i<L; i++{
		fmt.Printf("-f- %d\t%s\n",
			printpwd.child_files[i].size,
			printpwd.child_files[i].fullpath,
		)
	}
}


func cd(destination string){
	if destination == "/"{
		pwd = &root
	}else if destination == ".."{
		pwd = pwd.parent
		
	}else{
		cd_child(destination)
	}
	
}


func cd_child(destination string){
	for i,L := 0, len(pwd.child_dirs); i<L; i++{
		if pwd.child_dirs[i].name == destination{
			pwd = &pwd.child_dirs[i]
			return
		}
	}

	// if get out of loop means , des not found
	fmt.Printf("err, could find %s dir in pwd of %s\n",
		destination,
		pwd.fullpath,
	)
	for i,L := 0, len(pwd.child_dirs); i<L; i++{
		fmt.Printf("\t\t%d\t%s\n",
			i, pwd.child_dirs[i].name)
	}
	panic(100)
}

func du(dir * Directory) int { //disk-usage like haha
	//first sum of files
	sum:=0
	for i,L := 0, len(dir.child_files); i<L; i++{
		sum += dir.child_files[i].size
	}

	if len(dir.child_dirs) == 0{
		//no further dirs
		dir.size = sum // sum of the files bcuz there's no other dirs
		return dir.size
	}
	
	for i,L := 0, len(dir.child_dirs); i<L; i++{
		sum += du(&dir.child_dirs[i])
	}

	dir.size = sum // sum of the files bcuz there's no other dirs
	return dir.size
}

func run_all_dir(dir * Directory) int {
	sum:=0
	for i,L := 0, len(dir.child_dirs); i<L; i++{
		if dir.child_dirs[i].size < 100000{
			sum += dir.child_dirs[i].size
		}
		sum += run_all_dir(&dir.child_dirs[i])
	}
	return sum
}

var curSelected_forDeletion_size = 70000000;
var curSelected_dirDeletion * Directory;

func run_all_dir_find_minimumdir(dir * Directory, need2free int) {
	for i,L := 0, len(dir.child_dirs); i<L; i++{ 
		if dir.child_dirs[i].size > need2free{
			if dir.child_dirs[i].size < curSelected_forDeletion_size{
				curSelected_forDeletion_size = dir.child_dirs[i].size
				curSelected_dirDeletion = &dir.child_dirs[i]
				run_all_dir_find_minimumdir(&dir.child_dirs[i], need2free)
			}
		} else {
			//folder is too small to run inside
		}
	}
}



func main(){
	fmt.Printf("AoC2022\n");
	fname := "/home/garid/Documents/advent/AoC-2022/day7/input.txt"
	file, err:= os.Open(fname)
	if err != nil{
		fmt.Printf("Pls check %s file, err code %v", fname, err)
		panic(1)
	}
	reader := bufio.NewReader(file)
	for i:=0;;i++{
		line, ret := reader.ReadString('\n')
		if ret == io.EOF{
			fmt.Printf("File has ended\n")
			break
		}
		fmt.Printf("%d\t%s", i, line[:len(line)-1])

		if (line[0] == '$'){
			// means previoues command is done
			line_args := strings.Split(line[:len(line)-1], " ")
			fmt.Printf("\t\t%v\t", line_args)
			if len(line_args) < 2 {
				fmt.Printf("Bad cmd line args %v %s", line_args, line)
				panic(10)
			}
			cmd := line_args[1]
			if cmd == "cd"{
				if len(line_args) != 3{
					fmt.Printf("Bad cmd line args %v %s", line_args, line)
					panic(11)
				}
				destination_wd := line_args[2]
				fmt.Printf("cd to %s", destination_wd)
				cd(destination_wd)
				
			}else if cmd == "ls"{
				;
			}

		} else {
			// probably lising out the fnames
			populate_curr_pwd(line[:len(line)-1])
		}
		fmt.Println()
	}

	//this where finish
	du(&root)
	s := run_all_dir(&root)
	tree(&root)
	fmt.Println("sum: thisp part1:", s) //
	fmt.Println("used:", root.size)

	// now part 2
	freespace := 70000000 - root.size
	needtofree:= 30000000 - freespace

	run_all_dir_find_minimumdir(&root, needtofree)
	fmt.Printf("part 2:\n\n")
	fmt.Printf("we need to free up %v\n", needtofree)
	fmt.Printf("size: %v\n", curSelected_forDeletion_size)
	fmt.Printf("path: %s\n", curSelected_dirDeletion.fullpath)

}
