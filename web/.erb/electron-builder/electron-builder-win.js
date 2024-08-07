const winConfig = require("./electron-builder-win-config");

module.exports = {
    ...winConfig,
    extraResources: winConfig.extraResources.concat([
        {
            "from": "./win/stable-diffusion-server.exe",
            "to": "./oasis-server/stable-diffusion-server.exe"
        }
    ])
};
