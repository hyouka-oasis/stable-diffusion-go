import ast

import edge_tts
import argparse
import asyncio
import aiofiles
import re
import os
from datetime import datetime

async def time_difference(time1, time2, time_format=r"%H:%M:%S.%f"):
    time1 = datetime.strptime(time1, time_format)
    time2 = datetime.strptime(time2, time_format)
    print()
    # 计算时间差
    delta = time2 - time1
    time_diff = str(delta)[:11]
    return time_diff

async def srt_to_list(filename):
    subtitles = []  # 存储最终结果的列表
    text = []  # 临时存储当前字幕块的文本行
    time_code = None  # 初始化时间码变量

    with open(filename, "r", encoding="utf-8") as file:
        for line in file:
            line = line.strip()  # 移除行首尾的空白字符

            if "-->" in line:  # 检测时间码行
                if text:  # 如果前一个字幕块的文本已经读取，存储前一个字幕块
                    subtitles.append((time_code, " ".join(text)))
                    text = []  # 重置文本列表为下一个字幕块做准备
                time_code = line  # 更新时间码

            elif line:  # 非空行即为字幕文本行
                text.append(line)

        # 添加文件末尾的最后一个字幕块（如果有）
        if text:
            subtitles.append((time_code, " ".join(text)))
    return subtitles

class SubMarker(edge_tts.SubMaker):

    async def remove_non_chinese_chars(self, text):
        # 使用正则表达式匹配非中文字符和非数字
        # pattern = re.compile(r"[^\u4e00-\u9fff0-9]+")
        pattern = re.compile(r"[^\u4e00-\u9fffA-Za-z0-9]+")
        # 使用空字符串替换匹配到的非中文字符和非数字
        cleaned_text = re.sub(pattern, "", text)
        return cleaned_text

    async def generate_cn_subs(self, text) -> str:
        print(text, "获取到的文本")
        # 定义要指定的字符
        punctuation = ["，", "。", "！", "？", "；", "：", "”", ",", "!", "…", "+", "-"]

        def clause():
            start = 0
            i = 0
            text_list = []
            while i < len(text):
                if text[i] in punctuation:
                    try:
                        while text[i] in punctuation:
                            i += 1
                    except IndexError:
                        pass
                    text_list.append(text[start:i])
                    start = i
                i += 1
            return text_list

        self.text_list = clause()
        if len(self.subs) != len(self.offset):
            raise ValueError("文字长度已经达到最长了")
        data = "WEBVTT\r\n\r\n"
        i = 0
        for text in self.text_list:
            text = await self.remove_non_chinese_chars(text)
            try:
                start_time = self.offset[i][0]
            except IndexError:
                return data
            try:
                while self.subs[i + 1] in text:
                    i += 1
            except IndexError:
                pass
            data += edge_tts.submaker.formatter(start_time, self.offset[i][1], text)
            i += 1
        return data


"""
text_path 文本路径
mp3_path 输出的音频路径
srt_path 输出的字幕路径
"""
async def edge_tts_create_srt(text_path, mp3_path, srt_path, *edge_tts_args) -> None:
    # 打开文件
    with open(text_path, "r", encoding="utf-8") as f:
        content = f.read()

    communicate = edge_tts.Communicate(
        text=content,
        voice="Microsoft Server Speech Text to Speech Voice (en-US, AriaNeural)" if edge_tts_args[0] is None else edge_tts_args[0],
        rate="+0%" if edge_tts_args[1] is None else edge_tts_args[1],
        volume="+0%" if edge_tts_args[2] is None else edge_tts_args[2],
        pitch="+0Hz" if edge_tts_args[3] is None else edge_tts_args[3],
    )
    sub_marker = SubMarker()
    #写入音频文件
    async with aiofiles.open(mp3_path, "wb") as file:
        async for chunk in communicate.stream():
            if chunk["type"] == "audio":
                await file.write(chunk["data"])
            elif chunk["type"] == "WordBoundary":
                sub_marker.create_sub((chunk["offset"], chunk["duration"]), chunk["text"])
    #写入字幕文件
    async with aiofiles.open(srt_path, "w", encoding="utf-8") as file:
        content_to_write = await sub_marker.generate_cn_subs(content)
        await file.write(content_to_write)

# 生成字幕时间列表
async def create_processing_time(srt_path, text_path, txt_time_path):
    subtitles = await srt_to_list(srt_path)
    with open(text_path, "r", encoding="utf-8") as f:
        section_list = f.readlines()
        section_time_list = []
        index_ = 0
        time = "00:00:00.000"
        for si, section in enumerate(section_list):
            if len(section_list) == si + 1:
                # 最后这段不处理 默认使用剩余所有time
                next_start_time = subtitles[-1][0].split(" --> ")[1]
                diff = await time_difference(time, next_start_time)
                section_time_list.append(diff)
                break
            content_ = await SubMarker().remove_non_chinese_chars(section)
            for i, v in enumerate(subtitles):
                if i <= index_:
                    continue
                if v[1] not in content_:
                    next_start_time = v[0].split(" --> ")[0]
                    diff = await time_difference(time, next_start_time)
                    section_time_list.append(diff)
                    index_ = i
                    time = next_start_time
                    break
        with open(os.path.join(txt_time_path), "w", encoding="utf-8") as f3:
            f3.write(str(section_time_list))

# 初始化edge_tts
async def create_voice_caption():
    parser = argparse.ArgumentParser()
    # 从go那边获取过来的文本路径
    parser.add_argument("--book_path", help="字幕的文本路径地址")
    parser.add_argument("--audi_srt_map_path", help="字幕时间数组文本路径")
    parser.add_argument("--audio_path", help="输出的音频路径地址")
    parser.add_argument("--audio_srt_path", help="输出的字幕路径地址")
    parser.add_argument("--voice", help="角色")
    parser.add_argument("--rate", help="语速")
    parser.add_argument("--volume", help="音量")
    parser.add_argument("--pitch", help="分贝")
    args = parser.parse_args()
    book_path = args.book_path
    if book_path is None:
        raise Exception("输出路径不能为空")
    audi_srt_map_path = args.audi_srt_map_path
    if audi_srt_map_path is None:
        raise Exception("字幕切片文本路径不能为空")
    audio_path = args.audio_path
    if audio_path is None:
        raise Exception("音频输出路径不能为空")
    audio_srt_path = args.audio_srt_path
    if audio_srt_path is None:
        raise Exception("字幕输出路径不能为空")
    voice, rate, volume, pitch = args.voice, args.rate, args.volume, args.pitch
    # 通过edge-tts生成音频和字幕
    await edge_tts_create_srt(book_path, audio_path, audio_srt_path, voice, rate, volume, pitch)
    # 通过字幕生成时间表
    await create_processing_time(audio_srt_path, book_path, audi_srt_map_path)


if __name__ == "__main__":
    # with open("1.txt", "r", encoding="utf-8") as f:
    #     content = f.read()
    # time_list = ast.literal_eval(content)
    # print(time_list)
    # for index, (image_path, duration) in enumerate(
    #         zip(["1.png", "2.png", "3.png", "4.png", "5.png"], time_list)
    # ):
    #     print(image_path, duration, index, "时间列表")
    asyncio.run(create_voice_caption())
