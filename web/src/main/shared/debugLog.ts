import logger from 'electron-log';

/**
 * 最大大小 50m
 */
logger.transports.file.maxSize = 50 * 1024 * 1024;
logger.transports.file.format = "[{y}-{m}-{d} {h}:{i}:{s}.{ms}] [{level}]{scope} {text}";
logger.transports.file.level = 'debug';
const date = new Date();
const fileName = date.getFullYear() + "-" + (date.getMonth() + 1) + "-" + date.getDate();
logger.transports.file.fileName = `${fileName}.log`;

const log = logger.log;
const error = logger.error;
const info = logger.info;
const debug = logger.debug;
export {
    log,
    error,
    info,
    debug
};
