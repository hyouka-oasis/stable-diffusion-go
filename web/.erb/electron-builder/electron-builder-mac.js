const macConfig = require("./electron-builder-mac-config");

module.exports = {
    ...macConfig,
    extraResources: macConfig.extraResources.concat(
        {
            "from": "./mac/stable-diffusion-server",
            "to": "./oasis-server/stable-diffusion-server"
        }
    ),
};
