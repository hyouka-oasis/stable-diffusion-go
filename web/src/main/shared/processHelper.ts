import process from 'child_process';
import { PlatformHelper } from './PlatformHelper';
import { error, log } from "../shared/debugLog";

const windowPidRegex = /(?<=\.(exe)?\s*C?)(\d+)/gi;

export class ProcessHelper {
    /**
     * 获取进程id
     * */
    static getProgramPid(programName: string): string[] {
        const platform = PlatformHelper.validatePlatform();
        const findPortCmd = platform.isWindow ? `tasklist | findstr ${programName}` : `pgrep ${programName}`;
        log('打印找PID的指令', findPortCmd);
        let info = '';
        try {
            info = process.execSync(findPortCmd).toString('utf8');
        } catch (e: any) {
            error("获取进程PID失败", e);
            return [];
        }
        if (platform.isWindow) {
            const replaceInfo = info?.match(windowPidRegex);
            log('window下找的的进程', replaceInfo);
            if (!replaceInfo) {
                return [];
            }
            return replaceInfo;
        }
        return info?.split('\n')?.filter(pid => pid !== '');
    }

    /**
     * 通过PID关闭进程
     * */
    static killProgramByPid(ids: string | string[]) {
        const pids = Array.isArray(ids) ? ids : [ ids ];
        const platform = PlatformHelper.validatePlatform();
        for (let i = 0; i < pids.length; i++) {
            const pid = pids[i];
            const cmdTaskKill = platform.isWindow ? `taskkill /F /pid ${pid}` : `kill ${pid}`;
            try {
                process.execSync(cmdTaskKill);
                log("成功关闭进程", pid);
            } catch (e: any) {
                error("关闭进程失败进程ID", pid, '错误信息', e);
            }
        }
    }

    /**
     * 获取后端程序的pid 这一步一定确保不能抛出异常 最多就是打log
     */
    static getOasisServerPids() {
        const clientName = PlatformHelper.getGoServerName();
        if (!clientName) {
            return;
        }
        const platform = PlatformHelper.validatePlatform();
        const pids = ProcessHelper.getProgramPid(platform.isWindow ? clientName.replace('.exe', '') : clientName);
        // 过滤一遍空的避免有问题
        return pids?.filter(i => i);
    }
}
