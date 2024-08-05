const winConfig = require("./electron-builder-win-config");

module.exports = {
    ...winConfig,
    extraResources: winConfig.extraResources.concat([
        // {
        //     "from": "./oasis-server-windows-amd.exe",
        //     "to": "./oasis-server/oasis-server-windows-amd.exe"
        // }
    ])
};
