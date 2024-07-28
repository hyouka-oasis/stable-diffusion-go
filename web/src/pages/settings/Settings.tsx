import { useEffect, useState } from "react";
import { createSettings, getSettings, updateSettings } from "../../api/settings.ts";
import { Button, Form, Input, message, Select } from "antd";
import styled from "styled-components";

const translateConfigs = [
    {
        label: "ollama(本地需要运行ollama服务)",
        value: "ollama",
    },
    {
        label: "sd-prompt-translator(需要安装改名称插件,本地直译中文prompt)",
        value: "sd-prompt-translator",
    },
    {
        label: "aliyun",
        value: "aliyun",
    },
    {
        label: "chatgpt",
        value: "chatgpt",
    }
];

const SettingsPageWrap = styled.div`
`;

const SettingsPage = () => {
    const [ form ] = Form.useForm();
    const [ settingsId, setSettingsId ] = useState<number>();
    const [ messageApi, messageContext ] = message.useMessage();
    const getSettingsConfig = async () => {
        const config = await getSettings();
        setSettingsId(config.id);
        form.setFieldsValue({ ...config });
    };

    const onSubmitHandler = async () => {
        const formValues = await form.validateFields();
        if (settingsId) {
            await updateSettings({
                ...formValues,
                id: settingsId
            });
        } else {
            await createSettings({
                ...formValues,
            });
        }
        messageApi.success({
            content: "创建成功"
        });
    };

    useEffect(() => {
        getSettingsConfig();
    }, []);
    return (
        <SettingsPageWrap>
            {messageContext}
            <Form form={form} layout="vertical">
                <Form.Item label={"stable-diffusion配置"}>
                    <Form.Item rules={[ { required: true, message: '请输入url' } ]} label={"url"} name={[ "stableDiffusionConfig", "url" ]}>
                        <Input/>
                    </Form.Item>
                </Form.Item>
                <Form.Item rules={[ { required: true, message: '请选择翻译配置' } ]} label={"翻译配置"} name={"translateType"}>
                    <Select options={translateConfigs}/>
                </Form.Item>
                <Form.Item noStyle shouldUpdate={true}>
                    {
                        ({ getFieldValue }) => {
                            const value = getFieldValue("translateType");
                            if (value === "ollama")
                                return (
                                    <>
                                        <Form.Item rules={[ { required: true, message: '请输入ollama模型名称' } ]} name={[ "ollamaConfig", "url" ]} label={"ollama地址"}>
                                            <Input/>
                                        </Form.Item>
                                        <Form.Item rules={[ { required: true, message: '请输入ollama模型名称' } ]} name={[ "ollamaConfig", "modelName" ]} label={"ollama模型名称"}>
                                            <Input/>
                                        </Form.Item>
                                    </>
                                );
                            return null;
                        }
                    }
                </Form.Item>
                <Form.Item>
                    <Button type="primary" htmlType="submit" onClick={onSubmitHandler}>
                        {settingsId ? "修改" : "创建"}
                    </Button>
                </Form.Item>
            </Form>
        </SettingsPageWrap>
    );
};
export default SettingsPage;
