const linuxConfig = require("./electron-builder-linux-config");

module.exports = {
    ...linuxConfig,
    extraResources: linuxConfig.extraResources.concat([
        {
            "from": "./oasis-server-linux-amd",
            "to": "./oasis-server/oasis-server-linux-amd"
        }
    ])
};
