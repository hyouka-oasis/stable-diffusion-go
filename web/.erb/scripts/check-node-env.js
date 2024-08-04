import chalk from 'chalk';

export default function checkNodeEnv(expectedEnv) {
  if (!expectedEnv) {
    throw new Error('没有预期环境');
  }

  if (process.env.NODE_ENV !== expectedEnv) {
    console.log(
      chalk.whiteBright.bgRed.bold(
        `"process.env.NODE_ENV" 必须是 "${expectedEnv}" 才能被webpack应用`
      )
    );
    process.exit(2);
  }
}
