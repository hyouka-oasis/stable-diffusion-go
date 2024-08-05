const {
    basicBuilderConfig,
    buildExtraResourcesConfig
} = require("./electron-builder-config");

module.exports = {
    ...basicBuilderConfig,
    "afterSign": ".erb/scripts/notarize.js",
    "mac": {
        "type": "distribution",
        "hardenedRuntime": true,
        "entitlements": "assets/entitlements.mac.plist",
        "entitlementsInherit": "assets/entitlements.mac.plist",
        "gatekeeperAssess": false
    },
    "dmg": {
        "contents": [
            {
                "x": 130,
                "y": 220
            },
            {
                "x": 410,
                "y": 220,
                "type": "link",
                "path": "/Applications"
            }
        ]
    },
    "directories": {
        "app": "release/app",
        "buildResources": "assets",
        "output": "release/build/mac/${arch}"
    },
    "extraResources": [
        // {
        //     "from": "./src/main/shared/node_font/libs/darwin/fontlist",
        //     "to": "./oasis-server/node_font/darwin/fontlist"
        // }
    ].concat(buildExtraResourcesConfig()),
};
