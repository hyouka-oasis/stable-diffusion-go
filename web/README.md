## electron-react-webpack

<p>
通过electron-mon(二次封装的node-mon)实现electron主模块更新，
通过webpack.dll更快的编译模块
</p>

## 安装
```bash
npm install
```

## 启动
```bash
npm run start
```

## 打包web
```bash
npm run build:renderer
```

## 打包electron
```bash
npm run build:main
```

## 打包
```bash
npm run build
```

## 目录结构
```
|-- electron-react-webpack
    |-- electron-builder.json5
    |-- package.json
    |-- tsconfig.json
    |-- .erb 服务启动目录
    |   |-- configs
    |   |   |-- .eslintrc
    |   |   |-- webpack.config.base.ts
    |   |   |-- webpack.config.eslint.ts
    |   |   |-- webpack.config.loader.ts
    |   |   |-- webpack.config.main.prod.ts
    |   |   |-- webpack.config.preload.dev.ts
    |   |   |-- webpack.config.renderer.dev.dll.ts
    |   |   |-- webpack.config.renderer.dev.ts
    |   |   |-- webpack.config.renderer.prod.ts
    |   |   |-- webpack.paths.ts
    |   |-- dll 热更新依赖目录
    |   |   |-- preload.js
    |   |   |-- renderer.dev.dll.js
    |   |   |-- renderer.json
    |   |-- scripts 插件目录
    |       |-- .eslintrc
    |       |-- check-build-exists.ts
    |       |-- check-native-dep.js
    |       |-- check-node-env.js
    |       |-- check-port-in-use.js
    |       |-- clean.js
    |       |-- delete-source-maps.js
    |       |-- electron-rebuild.js
    |       |-- link-modules.ts
    |       |-- notarize.js
    |-- assets 资源目录
    |   |-- assets.d.ts
    |   |-- entitlements.mac.plist
    |   |-- icon.icns
    |   |-- icon.ico
    |   |-- icon.png
    |   |-- icon.svg
    |   |-- icons
    |       |-- 1024x1024.png
    |       |-- 128x128.png
    |       |-- 16x16.png
    |       |-- 24x24.png
    |       |-- 256x256.png
    |       |-- 32x32.png
    |       |-- 48x48.png
    |       |-- 512x512.png
    |       |-- 64x64.png
    |       |-- 96x96.png
    |-- release
    |   |-- app
    |       |-- package-lock.json
    |       |-- package.json
    |       |-- dist
    |           |-- main
    |           |   |-- main.js
    |           |   |-- main.js.map
    |           |   |-- preload.js
    |           |   |-- preload.js.map
    |           |-- renderer
    |               |-- index.html
    |               |-- renderer.js
    |               |-- renderer.js.LICENSE.txt
    |               |-- renderer.js.map
    |-- src 主代码目录
        |-- main
        |   |-- main.ts
        |   |-- preload.ts
        |   |-- util.ts
        |-- renderer
            |-- App.tsx
            |-- index.ejs
            |-- index.tsx
            |-- preload.d.ts
            |-- assets
                |-- test.less
                |-- test.module.less
                |-- audio
                    |-- t-rex-roar.mp3
                    |-- test.mp3
```