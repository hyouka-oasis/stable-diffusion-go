const { BUILDER_NAME } = require("./system-config");
const basicBuilderConfig = {
    "productName": BUILDER_NAME,
    "appId": "com.hyouka.stable.diffusion",
    "asar": true,
    "files": [
        "dist",
        "node_modules",
        "package.json"
    ],
    "asarUnpack": [ "**\\*.{node,dll}", "app.asar.unpacked" ],
    "publish": [
        {
            "provider": "generic",
            "url": ""
        }
    ]
};

const buildExtraPortPathList = [
    {
        "from": "./port.json",
        "to": "./server/port.json"
    },
];

const buildExtraResourcesConfig = () => {
    return [
        ...buildExtraPortPathList,
    ];
};

module.exports = {
    basicBuilderConfig,
    buildExtraResourcesConfig
};
