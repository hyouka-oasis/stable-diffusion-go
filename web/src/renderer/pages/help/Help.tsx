import { List } from "antd";

const Help = () => {
    const data = [
        "如何使用",
        "1.安装完成后,进入安装目录下面有个resource/oasis-server目录,用你熟悉的终端进入改目录,执行pip install -r requirements.txt(pip,以及python如何安装请自行百度)",
        "2.生成视频前请确保本地安装了ffmpeg",
        "3.进行生成操作时,请不要操作其他编辑操作,避免数据读取错误问题,开放版本不考虑回写数据库操作.",
        "4.lora列表存在的，在项目当中会将名称一致的添加到正向关键词当中去",
        "5.该项目全部基于本地,不用担心信息泄露的问题",
    ];
    return (
        <div>
            <List
                bordered
                dataSource={data}
                renderItem={(item) => (
                    <List.Item>
                        {item}
                    </List.Item>
                )}
            />
            如果还有什么不会，请右下角鼠标移入联系我微信,QQ不太上<br/>
        </div>
    );
};
export default Help;
