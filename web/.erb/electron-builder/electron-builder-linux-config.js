const {
    basicBuilderConfig,
    buildExtraResourcesConfig
} = require("./electron-builder-config");


module.exports = {
    ...basicBuilderConfig,
    "linux": {
        "target": [ "deb", "AppImage" ],
        "artifactName": "${productName}-${version}-${arch}-installer.${ext}",
        "icon": "assets/icon.icns",
    },
    "directories": {
        "app": "release/app",
        "buildResources": "assets",
        "output": "release/build/linux/${arch}"
    },
    "extraResources": buildExtraResourcesConfig(),
};
