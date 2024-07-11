import argparse

import aiofiles
import asyncio

PUNCTUATION = ["，", "。", "！", "？", "；", "：", "”", ",", "!", "…"]


# 根据小说文本进行分词
async def combine_strings(strings):
    combined = []
    current_srt = ""
    for s in strings:
        if min_words <= len(current_srt + s) <= max_words:
            combined.append(current_srt + s + "\n")
            current_srt = ""
        elif len(current_srt) > max_words:
            combined.append(current_srt + "\n")
            current_srt = s
        else:
            current_srt += s
    if current_srt:
        combined.append(current_srt + "\n")
    return combined


async def clause(text):
    start = 0
    i = 0
    text_list = []
    while i < len(text):
        if text[i] in PUNCTUATION:
            try:
                while text[i] in PUNCTUATION:
                    i += 1
            except IndexError:
                pass
            text_list.append(text[start:i].strip())
            start = i
        i += 1
    return text_list


# 分词和切割文本内容
async def participle(book_path: str, out_book_path: str):
    # 读取文件并且去除多余的空格，换行等
    async with aiofiles.open(book_path, "r", encoding="utf-8") as file:
        content = await file.read()
    novel = content.replace("\n", "").replace("\r", "").replace("\r\n", "").replace("\u2003", "")
    async with aiofiles.open(out_book_path, "w", encoding="utf-8") as file:
        # 先进行特殊字符进行校验
        text_list = await clause(novel)
        # 再通过最大，最小长度进行拼接
        result = await combine_strings(text_list)
        await file.writelines(result)
    return


async def main():
    book_path = args.book_path
    if book_path is None:
        raise Exception("源文件路径不能为空")
    participle_book_path = args.participle_book_path
    if participle_book_path is None:
        raise Exception("输出路径不能为空")
    await participle(book_path, participle_book_path)


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    # 从go那边获取过来的文本路径
    parser.add_argument("--book_path", help="原文件路径")
    parser.add_argument("--participle_book_path", help="输出路径")
    parser.add_argument("--max_words", help="最大长度")
    parser.add_argument("--min_words", help="最大长度")
    args = parser.parse_args()
    min_words = 30 if args.min_words is None else int(args.min_words)
    max_words = 30 if args.max_words is None else int(args.max_words)
    asyncio.run(main())
