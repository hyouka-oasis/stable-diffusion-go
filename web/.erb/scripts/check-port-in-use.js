import chalk from 'chalk';
import detectPort from 'detect-port';

const port = process.env.PORT || '2580';

detectPort(port, (err, availablePort) => {
    if (port !== String(availablePort)) {
        throw new Error(chalk.whiteBright.bgRed.bold(`端口 "${port}" 已经被占用. 请更换端口后重新启动`));
    } else {
        process.exit(0);
    }
});
