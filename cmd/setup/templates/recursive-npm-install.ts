import * as path from 'path';
import { execSync } from 'child_process';
import * as fs from 'fs';

console.log('Running npm install for all necessary directories');

const flatten = (lists: any[]) => {
    return lists.reduce((a: any[], b: any[]) => a.concat(b), []);
}

const ignoreFolders = ['node_modules', '.git', 'dist', 'terraform'];

const isDirectory = (source: string) => {
    return fs.lstatSync(source).isDirectory() &&
        ignoreFolders.reduce((acc, folderName) => source.indexOf(folderName) + acc, 0) < 0
}
const getDirectories = (source: string) => fs.readdirSync(source).map(name => {
    return path.join(source, name)
}).filter(isDirectory)


const getDirectoriesRecursive = (source: string) => {
    return [source, ...flatten(getDirectories(source).map(getDirectoriesRecursive))];
}

const directories: string[] =
    getDirectoriesRecursive(__dirname)
        .filter((directory: string) => {
            return fs.existsSync(path.join(directory, 'package.json'));
        });

console.log(directories);

directories.forEach(directory => {
    console.log('Installing ' + directories + '/package.json')
    execSync('npm install', { cwd: directory })
})
