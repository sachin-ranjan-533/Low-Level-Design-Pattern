package main


func main() {
	root := &directory{name: "root"}
	dir1 := root.createDirectory("dir1")
	dir2 := root.createDirectory("dir2")

	file1 := &file{name: "file1.txt"}
	file2 := &file{name: "file2.txt"}
	file3 := &file{name: "file3.txt"}

	root.addChild(dir1)
	root.addChild(file1)
	dir1.addChild(dir2)
	dir1.addChild(file2)
	dir2.addChild(file3)

	root.ls()
}