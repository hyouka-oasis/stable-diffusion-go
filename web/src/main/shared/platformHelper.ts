const { platform, arch } = process;

export class PlatformHelper {
    static getGoServerName() {
        if (platform === "win32") {
            return "stable-diffusion-server.exe";
        }
        return "stable-diffusion-server";
    }

    static getInstructionSet() {
        if (arch === "arm" || arch === "arm64") {
            return "arm";
        } else {
            return "amd";
        }
    }

    /**
     * 获取平台操作系统的名称
     * */
    static getPlatformName(): NodeJS.Platform {
        return platform;
    }

    /**
     * 校验当前平台
     * */
    static validatePlatform(): ValidatePlatformResult {
        const name = platform.toLocaleLowerCase();
        const mainOS = [ "darwin", "win32", "linux" ];
        const isMac = name === "darwin";
        const isWindow = name === "win32";
        const isLinux = name === "linux";
        return {
            isMac,
            isWindow,
            isLinux,
            isOther: ![ isMac, isWindow, isLinux ].some(it => it),
        };
    }
}

export interface ValidatePlatformResult {
    /**
     * Mac OS
     * */
    isMac: boolean;
    /**
     * Window OS
     * */
    isWindow: boolean;
    /**
     * Linux OS
     * */
    isLinux: boolean;
    /**
     * 其他类型的操作系统
     * */
    isOther: boolean;
}
