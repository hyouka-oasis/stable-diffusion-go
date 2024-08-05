const winConfig = require("./electron-builder-win");

module.exports = {
    ...winConfig,
    extraResources: winConfig.extraResources.concat(
        // {
        //     // TODO 暂时arm的用不了
        //     "from": "./oasis-server-windows-amd.exe",
        //     "to": "./oasis-server/oasis-server-windows-amd.exe"
        // },
    )
};
