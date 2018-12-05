import os
import re
import shutil
import zipfile


reg = "[a-zA-z0-9]{1,}\d{1,2}\.(?:in|out)"
reg = re.compile(reg)
ma = "\d{1,2}"
mareg = re.compile(ma)
tireg = "[a-zA-z]{1,}\d{1}"
tireg = re.compile(tireg)
curpath = os.getcwd()
curpathfile = os.listdir(os.getcwd())
#print(os.path.join(curpath, 'input'))
mp = {}

title = ""
minn = 5464
maxx = 0
sed = 1
val = 10
memory = 131072
tot = 0
flag = False



def gettitle(name):
    global title, minn, maxx
    if title == "":
        if (('0' <= name[0]) and (name[0] <= '9')):
            title = "Untitle"
        else:
            title = re.findall(tireg, name)[0]
            title = title[:-1]
    minn = min(int(re.findall(mareg, name)[0]), minn)
    print(minn)
    maxx = max(int(re.findall(mareg, name)[0]), maxx)
    print(maxx)


def getline(ind):
    if (ind != maxx) and (flag == False):
        return title + str(ind) + ".in" + "|" + title + str(
            ind) + ".out" + "|" + str(sed) + "|" + str(val) + "|" + str(
                memory) + "\n"
    else:
        return title + str(ind) + ".in" + "|" + title + str(
            ind) + ".out" + "|" + str(sed) + "|" + str(100 - tot) + "|" + str(
                memory) + "\n"


#print("请将本程序放置于数据文件的根目录下，该程序会自动识别.in输入文件与.out输出文件")
print("Please place the program under the root directory of the data file, the program will automatically identify")
#sed = int(input("请输入时限，单位秒:"))Please enter a time limit, the unit of a second
sed = int(input("Please enter a time limit(/s),if input 0, the default setting is 1s:"))
# memory = int(input("请输入内存容量，如果默认输入0，则默认128MB:"))Please enter the memory capacity, if the default input 0, the default 128 MB
memory = int(input("Please enter the memory limit (/MB), if input 0, the default setting is 128 MB"))
if (memory == 0):
    memory = 131072
else:
    memory *= 1024

if (sed == 0):
    sed = 1
inputf = os.path.join(curpath, 'input')
outputf = os.path.join(curpath, 'output')
inputfolder = os.path.exists(inputf)
outpurfolder = os.path.exists(outputf)

if not inputfolder:
    os.makedirs(inputf)
if not outpurfolder:
    os.makedirs(outputf)

for p in curpathfile:
    if (".out" in p):
        gettitle(p)
        break


with zipfile.ZipFile(title+".zip","w") as k:
    #k.write("input")
    #k.write("output")
    with open("config.ini", "w") as e:
        for p in curpathfile:
            if ".ini" in p:
                continue
            if os.path.isfile(p):
                if (".in" in p):
                    print(p)
                    gettitle(p)
                    # print(title)
                    shutil.copy(p, inputf)
                    k.write("input//"+p)
                if (".out" in p):
                    print(p)
                    shutil.copy(p, outputf)
                    k.write("output//"+p)
        e.write(str(maxx + 1 - minn) + "\n")
        val = 100 // (maxx + 1 - minn)
        if ((100 % (maxx + 1 - minn)) != 0):
            print("Warning, the number of data points isn't 100's factor, may cause data point score is not exactly the same")
            # print("警告，数据点非100的因子，可能造成数据点分数不完全相同")

        for i in range(minn, maxx + 1):
            e.write(getline(i))
            tot += val
    k.write("config.ini")

print(title)
print("The title\"{}\"'s data is already generated,total {} ,total of 100 points".format(title,maxx + 1 - minn))

# print("题目\"{}\"数据生成完毕,总计{}个测试点，总分100分".format(title,maxx + 1 - minn))
