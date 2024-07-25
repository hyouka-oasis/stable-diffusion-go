import { useEffect, useState } from "react";
import { getSettings, updateSettings } from "../../api/settings.ts";
import { Button, Form, Input, Select } from "antd";
import styled from "styled-components";

const translateConfigs = [
    {
        label: "ollama",
        value: "ollama",
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

    const getSettingsConfig = async () => {
        const config = await getSettings();
        setSettingsId(config.id);
        form.setFieldsValue({ ...config });
    };

    const onSubmitHandler = async () => {
        const formValues = await form.validateFields();
        await updateSettings({
            ...formValues,
            id: settingsId
        });
    };

    useEffect(() => {
        getSettingsConfig();
    }, []);
    return (
        <SettingsPageWrap>
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
                        确认
                    </Button>
                </Form.Item>
            </Form>
        </SettingsPageWrap>
    );
};
export default SettingsPage;
