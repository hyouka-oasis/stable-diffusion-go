import { app } from "electron";
import path from "path";

export const resourcePath = app.isPackaged
    ? path.join(process.resourcesPath, "assets")
    : path.join(__dirname, "../../../assets");

// TODO 这边需要调整
export const preloadUrl = app.isPackaged
    ? path.join(__dirname, "preload.js")
    : path.join(__dirname, "../../../.erb/dll/preload.js");

export const appIconPath = app.isPackaged
    ? path.join(resourcePath, "../", "icon.icns")
    : path.join(resourcePath, "icon.icns");
