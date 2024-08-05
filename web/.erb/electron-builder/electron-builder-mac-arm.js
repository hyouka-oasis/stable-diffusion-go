const macConfig = require("./electron-builder-mac-config");

module.exports = {
    ...macConfig,
    extraResources: macConfig.extraResources.concat([
        // {
        //     "from": "./oasis-server-darwin-arm",
        //     "to": "./oasis-server/oasis-server-darwin-arm"
        // }
    ])
};
