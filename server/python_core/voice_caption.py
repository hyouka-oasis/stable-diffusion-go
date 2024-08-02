import ast

import edge_tts
import argparse
import asyncio
import aiofiles
import re
import os
from datetime import datetime
from xml.sax.saxutils import unescape
from typing import List
from edge_tts.submaker import formatter


async def spilt_str2(s, t, k):
    """
    :param s: 切片文本
    :param t: 切分前时间
    :param k: 切分最大字数
    :return:  新的切片信息

    @ samples
        s = "并且觉醒天赋 得到力量 对抗凶兽 觉醒天赋 便是人人在十八岁时能以血脉沟通沟通 觉醒天赋"
        t = "00:00:35,184 --> 00:00:42,384"
        k = 15
    """

    async def time2second(ti):
        """
        :param ti: 输入时间， 格式示例：00:02:56,512
        :return: float
        """
        a, b, _c = ti.split(":")
        c, d = _c.split(",")

        a, b, c, d = int(a), int(b), int(c), int(d)

        second = a * 3600 + b * 60 + c + d / 1000

        return second

    async def second2time(si):
        hours = int(si // 3600)
        minutes = int((si % 3600) // 60)
        seconds = int(si % 60)
        milliseconds = round((si % 1) * 1000)

        v = "00"
        u = "000"
        a = v[: 2 - len(str(hours))] + str(hours)
        b = v[: 2 - len(str(minutes))] + str(minutes)
        c = v[: 2 - len(str(seconds))] + str(seconds)
        d = u[: 3 - len(str(milliseconds))] + str(milliseconds)

        return f"{a}:{b}:{c},{d}"

    ss = s.split(" ")
    ss_valid = []

    # todo 将所有片段设置成不超过15
    for _ss in ss:
        if len(_ss) > k:

            # 暴力截断几段
            e = len(_ss) // k + 1
            n_e = len(_ss) // e + 1

            for _i in range(e):
                if _i == e - 1:
                    ss_valid.append(_ss[n_e * _i:])
                else:
                    ss_valid.append(_ss[n_e * _i: n_e * (_i + 1)])
        else:
            ss_valid.append(_ss)

    # todo 片段合并
    tmp = ""
    new_ss = []
    for i in range(len(ss_valid)):
        tmp += ss_valid[i]

        if i < len(ss_valid) - 1:
            if len(tmp + ss_valid[i + 1]) > k:
                new_ss.append(tmp)
                tmp = ""
            else:
                continue
        else:
            new_ss.append(tmp)
            tmp = ""

    # 分配时间戳
    t1, t2 = t.split("-->")
    ft1 = await time2second(t1)
    ft2 = await time2second(t2)
    ftd = ft2 - ft1

    # 转换成秒数
    all_str = " ".join(new_ss)

    tt_s = 0
    line_srt = []
    for z in new_ss:
        tt_e = len(z) + tt_s

        # 文章最后一句异常处理
        if len(all_str) * ftd == 0:
            continue

        t_start = tt_s / len(all_str) * ftd
        t_end = tt_e / len(all_str) * ftd
        t_start = round(t_start, 3)
        t_end = round(t_end, 3)

        rec_s = await second2time(ft1 + t_start)
        rec_e = await second2time(ft1 + t_end)

        cc = (f"{rec_s} --> {rec_e}", z)
        line_srt.append(cc)

        tt_s = tt_e + 1

    return line_srt


async def time_difference(time1, time2, time_format=r"%H:%M:%S,%f"):
    time1 = datetime.strptime(time1, time_format)
    time2 = datetime.strptime(time2, time_format)
    # 计算时间差
    delta = time2 - time1
    time_diff = str(delta).replace(".", ",")[:11]
    return time_diff


async def load_srt_new(filename, flag=True):
    time_format = r"(\d{2}:\d{2}:\d{2}),\d{3} --> (\d{2}:\d{2}:\d{2}),\d{3}"

    n = 0  # srt 文件总行数
    index = 0  # strs 文字串移动下标
    line_tmp = ""  # 每个时间区间后的字数累计
    count_tmp = 0  # 每个时间区间后的字数行计数
    new_srt = []

    async with aiofiles.open(filename, mode="r", encoding="utf-8") as f3:
        f_lines = await f3.readlines()
        for line in f_lines:
            line = line.strip("\n")
            n += 1

            # 写入新的数据
            #   1)当出现在文本末写入一次
            if n == len(f_lines):
                new_srt_line = await spilt_str2(line_tmp, t_line_cur, limit)
                new_srt.append(new_srt_line)

            #   2）当新的一行是数字时，srt语句段写入
            # case1: 判断新的一行是不是数字
            if line.isdigit():
                if flag:
                    print(line)
                if n > 1:
                    new_srt_line = await spilt_str2(line_tmp, t_line_cur, limit)
                    new_srt.append(new_srt_line)
                continue

            # case2: 判断新的一行是不是时间段
            if re.match(time_format, line):
                t_line_cur = line
                # reset line_tmp
                line_tmp = ""
                count_tmp = 0
                continue

            # case3: 判断新的一行是空格时
            if len(line) == 0:
                continue

            # case4: 新的一行不属于上面其中之一
            line_std = line.replace(" ", "")
            if flag:
                print(f"{line}\n{line_std}")

            if count_tmp:
                line_tmp += " " + line_std
            else:
                line_tmp += line_std
            count_tmp += 1

    srt = []
    for _line in new_srt:
        for _l in _line:
            srt.append(_l)
    return srt


async def srt_to_list(filename):
    subtitles = []  # 存储最终结果的列表
    text = []  # 临时存储当前字幕块的文本行
    time_code = None  # 初始化时间码变量

    with open(filename, "r", encoding="utf-8") as file:
        for line in file:
            line = line.strip()  # 移除行首尾的空白字符

            if line.isdigit():  # 跳过字幕编号行
                continue

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

    def generate_subs_pro(self, content_list) -> str:

        if len(self.subs) != len(self.offset):
            raise ValueError("subs and offset are not of the same length")

        data = ""
        sub_state_count = 0
        sub_state_start = -1.0
        sub_state_subs = ""
        index = 0
        content = content_list[index]

        for idx, (offset, subs) in enumerate(zip(self.offset, self.subs)):
            start_time, end_time = offset
            subs = unescape(subs)

            if len(sub_state_subs) > 0:
                sub_state_subs += " "
            sub_state_subs += subs

            if sub_state_start == -1.0:
                sub_state_start = start_time
            sub_state_count += 1
            if idx == len(self.offset) - 1 or self.subs[idx + 1] not in content:
                subs = sub_state_subs

                split_subs: List[str] = [
                    subs[i: i + 100] for i in range(0, len(subs), 100)
                ]
                for i in range(len(split_subs) - 1):
                    sub = split_subs[i]
                    split_at_word = True
                    if sub[-1] == " ":
                        split_subs[i] = sub[:-1]
                        split_at_word = False

                    if sub[0] == " ":
                        split_subs[i] = sub[1:]
                        split_at_word = False

                    if split_at_word:
                        split_subs[i] += "-"

                data += formatter(
                    start_time=sub_state_start,
                    end_time=end_time,
                    subdata="\r\n".join(split_subs),
                )
                sub_state_count = 0
                sub_state_start = -1
                sub_state_subs = ""
                if index < len(content_list) - 1:
                    index += 1
                    content = content_list[index]
            else:
                content = re.sub(subs, '', content, count=1)
        return data


# 通过edge-tts生成字幕文件
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
    # 写入音频文件
    async with aiofiles.open(mp3_path, "wb") as file:
        async for chunk in communicate.stream():
            if chunk["type"] == "audio":
                await file.write(chunk["data"])
            elif chunk["type"] == "WordBoundary":
                sub_marker.create_sub((chunk["offset"], chunk["duration"]), chunk["text"])
    if language == "zh":
        # 写入字幕文件
        async with aiofiles.open(vtt_path, "w", encoding="utf-8") as file:
            content_to_write = await sub_marker.generate_cn_subs(content)
            await file.write(content_to_write)
        # vtt -》 srt
        idx = 1  # 字幕序号
        with open(srt_path, "w", encoding="utf-8") as f_out:
            for line in open(vtt_path, encoding="utf-8"):
                if "-->" in line:
                    f_out.write("%d\n" % idx)
                    idx += 1
                    line = line.replace(".", ",")  # 这行不是必须的，srt也能识别'.'
                if idx > 1:  # 跳过header部分
                    f_out.write(line)
    else:
        # 定义一个空列表来存储文件的每一行
        lines = []

        # 打开文件并读取每一行
        with open(participle_book_path, 'r', encoding='utf-8') as file:
            for line in file:
                # 去除每行末尾的换行符，并添加到列表中
                lines.append(line.strip())

        # 使用列表推导式去除空行
        lines = [line for line in lines if line and not line.isdigit()]

        with open(srt_tmp_path, "w", encoding="utf-8") as f_out:
            f_out.write(sub_marker.generate_subs_pro(lines))


# 生成字幕时间列表
async def create_processing_time(text_path, txt_time_path):
    subtitles = await srt_to_list(audio_srt_path)
    if language == "zh":
        with open(text_path, "r", encoding="utf-8") as f:
            section_list = f.readlines()
        section_time_list = []
        index_ = 0
        time = "00:00:00,000"
        for si, section in enumerate(section_list):
            if len(section_list) == si + 1:
                # 最后这段不处理 默认使用剩余所有time
                next_start_time = subtitles[-1][0].split(" --> ")[1]
                diff = await time_difference(time, next_start_time)
                diff = diff.replace(",", ".")
                section_time_list.append(diff)
                break
            content_ = await SubMarker().remove_non_chinese_chars(section)
            for i, v in enumerate(subtitles):
                if i <= index_:
                    continue
                if v[1] not in content_:
                    next_start_time = v[0].split(" --> ")[0]
                    diff = await time_difference(time, next_start_time)
                    diff = diff.replace(",", ".")
                    section_time_list.append(diff)
                    index_ = i
                    time = next_start_time
                    break
        with open(os.path.join(txt_time_path), "w", encoding="utf-8") as f3:
            f3.write(str(section_time_list))
    else:
        time_list = []
        init_time = "00:00:00.000"
        for subtitle in subtitles:
            duration = await time_difference(init_time, subtitle[0].split(" --> ")[1], time_format=r"%H:%M:%S.%f")
            time_list.append(duration)
            init_time = subtitle[0].split(" --> ")[1]
        with open(os.path.join(txt_time_path), "w", encoding="utf-8") as f3:
            f3.write(str(time_list))


# 保存字幕
async def save_srt(filename, srt_list):
    async with aiofiles.open(filename, mode="w", encoding="utf-8") as f:
        for _li, _l in enumerate(srt_list):
            if _li == len(srt_list) - 1:
                info = "{}\n{}\n{}".format(_li + 1, _l[0], _l[1])
            else:
                info = "{}\n{}\n{}\n\n".format(_li + 1, _l[0], _l[1])
            await f.write(info)


# 生成新字幕
async def srt_regen_new(srt_tmp_path, srt_path, flag):
    srt_list = await load_srt_new(srt_tmp_path, flag)
    await save_srt(srt_path, srt_list)


# 初始化edge_tts
async def create_voice_caption():
    # 通过edge-tts生成音频和字幕
    await edge_tts_create_srt(participle_book_path, audio_path, srt_tmp_path, voice, rate, volume, pitch)
    if language == "zh":
        await srt_regen_new(srt_tmp_path, audio_srt_path, False)
    else:
        os.replace(srt_tmp_path, audio_srt_path)
    # 通过字幕生成时间表
    await create_processing_time(participle_book_path, audi_srt_map_path)


if __name__ == "__main__":
    # audio_srt_path = "D:\\ComicTweetsGo\\server\\神秘复苏2-10\\participle\\test.wav"
    # participle_book_path = "../测试/participle/神秘复苏.txt"
    # audi_srt_map_path = "../测试/participle/神秘复苏time.txt"
    # asyncio.run(map3_to_srt(audio_srt_path))

    # audio_path = "F:\\stable-diffusion-go\\server\\神秘复苏16\\participle\\神秘复苏16.mp3"
    # participle_book_path = "F:\\stable-diffusion-go\\server\\神秘复苏16\\participle\\神秘复苏16.txt"
    # audi_srt_map_path = "F:\\stable-diffusion-go\\server\\神秘复苏16\\神秘复苏16map.txt"
    # audio_srt_path = "F:\\stable-diffusion-go\\server\\神秘复苏16\\participle\\神秘复苏16.srt"
    # voice, rate, volume, pitch, language, limit = "zh-CN-YunxiNeural", "+10%", "+100%", "+0Hz", "zh", 10
    # asyncio.run(create_processing_time(audio_srt_path, participle_book_path, audi_srt_map_path))
    parser = argparse.ArgumentParser()
    # 从go那边获取过来的文本路径
    parser.add_argument("--participle_book_path", help="字幕的文本路径地址")
    parser.add_argument("--audi_srt_map_path", help="字幕时间数组文本路径")
    parser.add_argument("--audio_path", help="输出的音频路径地址")
    parser.add_argument("--audio_srt_path", help="输出的字幕路径地址")
    parser.add_argument("--voice", help="角色")
    parser.add_argument("--rate", help="语速")
    parser.add_argument("--volume", help="音量")
    parser.add_argument("--pitch", help="分贝")
    parser.add_argument("--language", help="语言")
    parser.add_argument("--limit", help="每一行限制最大数")
    args = parser.parse_args()
    participle_book_path = args.participle_book_path
    if participle_book_path is None:
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
    voice, rate, volume, pitch, language, limit = args.voice, args.rate, args.volume, args.pitch, args.language, int(args.limit)
    vtt_path = audio_srt_path.replace(".srt", ".vtt")
    srt_tmp_path = audio_srt_path.replace(".srt", ".tmp.srt")
    asyncio.run(create_voice_caption())
