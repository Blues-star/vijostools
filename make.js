const { prompt } = require('enquirer');
const fs = require('fs')
const path = require('path')
function deleteFolder(path) {
    if (fs.existsSync(path)) {
        fs.readdirSync(path).forEach(function (file) {
            let curPath = path + "/" + file;
            if (fs.statSync(curPath).isDirectory()) deleteFolder(curPath);
            else fs.unlinkSync(curPath);
        });
        fs.rmdirSync(path);
    }
}
function mkdirSync(dirname) {
    if (fs.existsSync(dirname)) return true;
    else if (mkdirSync(path.dirname(dirname)))
        fs.mkdirSync(dirname);
    return true;
}
function getInputFiles(files) {
    const regs = [
        [/^([a-zA-Z]*)([0-9]+).in$/, a => { return a[1] + a[2] + '.ans' }, a => { return Number(a[2]) }],
        [/^([a-zA-Z]*)([0-9]+).in$/, a => { return a[1] + a[2] + '.out' }, a => { return Number(a[2]) }],
        [/^([a-zA-Z0-9]*)\.in([0-9]*)$/, a => { return a[1] + '.ou' + a[2]; }, a => { return Number(a[2]) }],
        [/^(input)([0-9]*).txt$/, a => { return 'output' + a[2] + '.txt' }, a => { return Number(a[2]) }],
    ]
    let cases = [];
    for (let i in files)
        for (let j in regs)
            if (regs[j][0].test(files[i])) {
                let data = regs[j][0].exec(files[i]);
                if (fs.existsSync('./' + files[i]) && fs.existsSync('./' + regs[j][1](data))) {
                    cases.push({ input: files[i], output: regs[j][1](data), sort: regs[j][2](data) });
                    break;
                }
            }
    cases.sort((a, b) => { return a.sort - b.sort });
    return cases;
}
prompt([
    {
        type: 'select',
        name: 'mem_limit',
        message: 'Memory_Limit:',
        choices: ['128', '256', '512', '64', '32']
    },
    {
        type: 'select',
        name: 'time_limit',
        message: 'Timt_Limit:',
        choices: ['1', '2', '3', '4', '5']
    }
]).then(async res => {
    const AdmZip = require('adm-zip');
    let zip = new AdmZip();
    const cases = getInputFiles(fs.readdirSync('.'));
    const cnt = cases.length;
    let config = cnt + '\n';
    const sc = Math.floor(100 / cnt);
    const d = cnt - 100 % sc;
    deleteFolder('./Data');
    mkdirSync('./Data/Input');
    mkdirSync('./Data/Output');
    for (let i in cases) {
        fs.renameSync('./' + cases[i].input, './Data/Input/' + cases[i].input);
        fs.renameSync('./' + cases[i].output, './Data/Output/' + cases[i].output);
        config += cases[i].input + '|' + cases[i].output + '|' + res.time_limit + '|' + (i >= d ? sc + 1 : sc) + '|' + res.memory_limit + '\n';
    }
    fs.writeFileSync('./Data/Config.ini', config);
    zip.addLocalFolder('./Data/Input');
    zip.addLocalFolder('./Data/Output');
    zip.addLocalFile('./Data/Config.ini');
    await new Promise((resolve, reject) => {
        zip.writeZip('./data.zip', resolve);
    });
    console.log('Done!');
});
