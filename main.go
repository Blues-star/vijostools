package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var mp = make(map[string]int)
var inputlist, outputlist []string
var filelist []string
var name string
var tot int
var problem string
var maxcnt = 0
var ind [55]int
var start = 516874
var end = 0
var sed, memory = 0, 0
var perscore, lastscore = 0, 0
var inputexp, outputexp string

func check(filename os.FileInfo) bool {
	if filename.IsDir() == true {
		return false
	}
	if strings.Contains(filename.Name(), ".ini") {
		return false
	}
	if strings.Contains(filename.Name(), ".in") || strings.Contains(filename.Name(), ".out") {
		return true
	}
	return false
}
func add(title string) {
	if _, ok := mp[title]; ok == true {
		mp[title]++
	} else {
		mp[title] = 1
	}
}
func copyfile(src, des string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}
	defer srcFile.Close()

	desFile, err := os.Create(des)
	if err != nil {
		fmt.Println(err)
	}
	defer desFile.Close()

	return io.Copy(desFile, srcFile)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func initfolder() {
	_, ierr := os.Stat("./input")
	if os.IsNotExist(ierr) {
		os.Mkdir("./input", os.ModePerm)
	}
	_, oerr := os.Stat("./output")
	if os.IsNotExist(oerr) {
		os.Mkdir("./output", os.ModePerm)
	}
}
func Zip(srcFile []string, destZip string) error {
	zipfile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()
	for _, file_or_folder := range srcFile {
		filepath.Walk(file_or_folder, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}
			header.Name = strings.TrimPrefix(path, filepath.Dir(file_or_folder)+"/")
			// header.Name = path
			if info.IsDir() {
				header.Name += "/"
			} else {
				header.Method = zip.Deflate
			}
			writer, err := archive.CreateHeader(header)
			if err != nil {
				return err
			}
			if !info.IsDir() {
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()
				_, err = io.Copy(writer, file)
			}
			return err
		})
	}

	return err
}
func main() {
	fmt.Println("vijos数据打包器-golang重制版1.0,powered by 星夜的蓝天")
	fmt.Println("请将本程序放置于数据文件的根目录下，该程序会自动识别.in输入文件与.out或.ans输出文件")
	fmt.Println("请确保输入输出文件名形如 XXX1.in XXX1.out或XXX1.ans")
	fmt.Println("请输入时限，单位秒,如果输入0，则默认为1秒钟:")
	fmt.Scanln(&sed)
	if sed == 0 {
		sed = 1
	}
	fmt.Println("请输入内存容量，单位MB，如果输入0，则默认128MB:")
	fmt.Scanln(&memory)
	if memory == 0 {
		memory = 128 * 1024
	} else {
		memory *= 1024
	}
	var orifilelist, _ = ioutil.ReadDir("./")
	for _, fname := range orifilelist {
		if check(fname) == true {
			filelist = append(filelist, fname.Name())
		}
	}
	sort.Strings(filelist)
	for _, fi := range filelist {
		if strings.Contains(fi, ".in") {
			reg := regexp.MustCompile(`[a-zA-Z]{1,}\d`)
			temp := reg.FindAllString(fi, -1)[0]
			//name = fi[0 : len(fi)-4]
			name = temp[0 : len(temp)-1]
			add(name)
			inputlist = append(inputlist, fi)
		} else {
			reg := regexp.MustCompile(`[a-zA-Z]{1,}\d`)
			temp := reg.FindAllString(fi, -1)[0]
			//name = fi[0 : len(fi)-5]
			name = fi[0 : len(temp)-1]
			add(name)
			outputlist = append(outputlist, fi)
		}
	}
	fmt.Println("文件列表处理完成")
	for i := range mp {
		if maxcnt < mp[i] {
			maxcnt = mp[i]
			problem = i
		}
	}
	initfolder()
	if len(inputlist) == len(outputlist) {
		fmt.Printf("匹配到%d对数据点\n", len(inputlist))
		//fmt.Printf("match %d", len(inputlist))
	}
	reg := regexp.MustCompile(`\d{1,}\.`)
	for i := 0; i < len(inputlist); i++ {
		temp := reg.FindAllString(inputlist[i], -1)[0]
		index := temp[0 : len(temp)-1]
		intindex, _ := strconv.Atoi(index)
		if intindex < start {
			start = intindex
		}
		if intindex > end {
			end = intindex
		}
		ind[intindex]++
		copyfile(inputlist[i], "./input/"+inputlist[i])
		copyfile(outputlist[i], "./output/"+outputlist[i])
	}
	outputexp = outputlist[0][len(outputlist[0])-4:]
	//such .out or .ans
	fmt.Println("检测到答案文件后缀名为", outputexp)
	if end-start+1 == len(inputlist) {
		//fmt.Print("match")
		fmt.Println("数据点二次检查通过")
	}
	if 100%(end-start+1) != 0 {
		fmt.Print("警告，数据点个数非100的因子，可能造成数据点分数不完全相同")
		perscore = 100 / (end - start + 1)
		lastscore = 100 - (end-start)*perscore
	} else {
		perscore = 100 / (end - start + 1)
		lastscore = perscore
	}
	var content = strconv.Itoa(end-start+1) + "\n"
	for i := start; i <= end-1; i++ {
		content += (problem + strconv.Itoa(i) + ".in|" + problem + strconv.Itoa(i) + outputexp + "|" + "1|" + strconv.Itoa(perscore) + "|" + strconv.Itoa(memory) + "\n")
	}
	content += (problem + strconv.Itoa(end) + ".in|" + problem + strconv.Itoa(end) + outputexp + "|" + "1|" + strconv.Itoa(lastscore) + "|" + strconv.Itoa(memory) + "\n")
	os.Stdout, _ = os.OpenFile("./config.ini", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	fmt.Print(content)
	file := []string{"input", "output", "config.ini"}
	Zip(file, problem+".zip")
}
