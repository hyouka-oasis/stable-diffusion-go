const {
    basicBuilderConfig,
    buildExtraResourcesConfig
} = require("./electron-builder-config");


module.exports = {
    ...basicBuilderConfig,
    "icon": "assets/icon.ico",
    "nsis": {
        "artifactName": "${productName}-${version}-${arch}-installer.${ext}",
        "oneClick": false,
        "allowElevation": true,
        "allowToChangeInstallationDirectory": true,
        "createDesktopShortcut": true,
        "createStartMenuShortcut": true,
        "shortcutName": "${productName}"
    },
    "win": {
        "target": [
            {
                "target": "nsis",
            }
        ]
    },
    "directories": {
        "app": "release/app",
        "buildResources": "assets",
        "output": "release/build/windows/${arch}"
    },
    "extraResources": [
        // {
        //     "from": "./src/main/shared/node_font/libs/win/fonts.vbs",
        //     "to": "./oasis-server/node_font/win/fonts.vbs"
        // }
    ].concat(buildExtraResourcesConfig()),
};
