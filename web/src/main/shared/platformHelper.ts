const { platform, arch } = process;

export class PlatformHelper {

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
