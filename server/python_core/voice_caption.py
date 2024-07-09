import edge_tts
import argparse
import asyncio
import aiofiles
import re



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
        voice=edge_tts_args.voice,
        rate="+30%",
        volume="+100%"
    )
    sub_marker = SubMarker()
    async with aiofiles.open(mp3_path, "wb") as file:
        async for chunk in communicate.stream():
            if chunk["type"] == "audio":
                await file.write(chunk["data"])
            elif chunk["type"] == "WordBoundary":
                sub_marker.create_sub((chunk["offset"], chunk["duration"]), chunk["text"])
    async with aiofiles.open(srt_path, "w", encoding="utf-8") as file:
        content_to_write = await sub_marker.generate_cn_subs(content)
        await file.write(content_to_write)


# 初始化edge_tts
async def create_voice_caption():
    parser = argparse.ArgumentParser()
    # 从go那边获取过来的文本路径
    parser.add_argument("--text_path", help="需要转字幕的文本路径地址")
    parser.add_argument("--mp3_path", help="输出的音频路径地址")
    parser.add_argument("--srt_path", help="输出的字幕路径地址")
    parser.add_argument("--voice", help="角色")
    parser.add_argument("--rate", help="语速")
    parser.add_argument("--volume", help="音量")
    parser.add_argument("--pitch", help="分贝")
    args = parser.parse_args()
    text_path = args.text_path
    if text_path is None:
        raise Exception("文件路径不能为空")
    mp3_path = args.mp3_path
    if mp3_path is None:
        raise Exception("音频输出路径不能为空")
    srt_path = args.srt_path
    if srt_path is None:
        raise Exception("字幕输出路径不能为空")
    voice, rate, volume, pitch = args.voice, args.rate, args.volume, args.pitch
    await edge_tts_create_srt(text_path, mp3_path, srt_path, voice, rate, volume, pitch)


asyncio.run(create_voice_caption())
