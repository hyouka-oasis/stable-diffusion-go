import argparse
import jieba
import jieba.analyse
import asyncio
from collections import Counter
import aiofiles

async def extract_potential_names():
    words = jieba.cut(text)
    keywords = [word for word in words if word in jieba.analyse.extract_tags(text)]
    # word_count = Counter()
    # for word in words:
    #     # 过滤掉单个字
    #     if len(word) > 1:
    #         word_count[word] += 1
    # all_words = [word for word, count in word_count.most_common()]
    async with aiofiles.open(save_path, "w", encoding="utf-8") as file:
        await file.write(",".join(keywords))

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    # 从go那边获取过来的文本路径
    parser.add_argument("--text", help="文本内容")
    parser.add_argument("--save_path", help="保存路径")
    args = parser.parse_args()
    text = args.text
    save_path = args.save_path
    if save_path is None:
        raise Exception("保存路径不能为空")
    asyncio.run(extract_potential_names())
