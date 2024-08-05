const fs = require('fs');
const pathUtils = require('../configs/webpack.paths');
const {
    VERSION,
    PACKAGE_JSON_NAME
} = require('../electron-builder/system-config');

const appPackage = pathUtils.default.appPackagePath;

const rootPackage = pathUtils.default.packagePath;

const readPackageList = [
    appPackage,
    rootPackage
];


const replaceFile = () => {
    for (const filePath of readPackageList) {
        try {
            const fileJson = fs.readFileSync(filePath, 'utf-8');
            const fileJsonParse = JSON.parse(fileJson);
            const newDate = new Date();
            const year = newDate.getFullYear();
            const mouth = newDate.getMonth() + 1;
            const day = newDate.getDate();
            const hours = newDate.getHours();
            const minutes = newDate.getMinutes();
            const seconds = newDate.getSeconds();
            const buildDate = `${year}/${mouth}/${day} ${hours}:${minutes}:${seconds}`;
            fileJsonParse.name = PACKAGE_JSON_NAME;
            fileJsonParse.version = VERSION;
            fileJsonParse.description = `客户端打包时间: ${buildDate}`;
            fs.writeFileSync(filePath, JSON.stringify(fileJsonParse, null, 2));
        } catch (e) {
            console.log(e);
            break;
        }
    }
};
replaceFile();
