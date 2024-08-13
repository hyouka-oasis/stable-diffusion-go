import { useEffect, useState } from "react";
import { Button, Form, Input, message, Select } from "antd";
import styled from "styled-components";
import { ollamaApi, settingsApi } from "renderer/api";
import FileSvg from "renderer/assets/svg-com/file.svg";
import { ipcApi } from "renderer/ipc/BasicRendererIpcAdapter";
import { DefaultOptionType } from "rc-select/lib/Select";

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
        label: "阿里云",
        value: "aliyun",
    },
    {
        label: "chatgpt",
        value: "chatgpt",
    }
];

const SettingsPageWrap = styled.div`
    .sava-path {
        input {
            pointer-events: none;
        }
    }

    .file-svg {
        cursor: pointer;
    }
`;

const SettingsPage = () => {
    const [ form ] = Form.useForm();
    const [ settingsId, setSettingsId ] = useState<number>();
    const [ messageApi, messageContext ] = message.useMessage();
    const [ modelList, setModelList ] = useState<DefaultOptionType[]>([]);
    const getSettingsConfig = async () => {
        const config = await settingsApi.getSettings();
        setSettingsId(config.id);
        form.setFieldsValue({ ...config });
    };

    const onSubmitHandler = async () => {
        const formValues = await form.validateFields();
        if (settingsId) {
            await settingsApi.updateSettings({
                ...formValues,
                id: settingsId
            });
        } else {
            await settingsApi.createSettings({
                ...formValues,
            });
        }
        messageApi.success({
            content: settingsId ? "修改成功" : "创建成功"
        });
        await getSettingsConfig();
    };

    const onFilePathSelect = async () => {
        const folderValues = await ipcApi.fileAdapter.onFolderSelect({
            properties: [ "openDirectory" ],
        });
        if (!folderValues.data.canceled) {
            const selectPath = folderValues.data.filePaths[0];
            form.setFieldsValue({ savePath: selectPath });
        }
    };

    const getOllamaModelList = async (open) => {
        if (open) {
            const data = await ollamaApi.getOllamaModelList();
            setModelList(() => {
                return data?.list?.map(item => ({
                    label: item.name,
                    value: item.model
                }));
            });
        }
    };

    useEffect(() => {
        getSettingsConfig();
    }, []);
    return (
        <SettingsPageWrap>
            {messageContext}
            <Form form={form} layout="vertical">
                <Form.Item rules={[ { required: true, message: '请输入项目保存路径' } ]} label={"项目保存路径"} name={"savePath"}>
                    <Input className={"sava-path"} addonAfter={<div className={"file-svg"} onClick={onFilePathSelect}><FileSvg/></div>}/>
                </Form.Item>
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
                                            <Select options={modelList} onDropdownVisibleChange={getOllamaModelList}/>
                                        </Form.Item>
                                    </>
                                );
                            if (value === "aliyun")
                                return (
                                    <>
                                        <Form.Item rules={[ { required: true, message: '请填写阿里云ACCESS_KEY_ID' } ]} name={[ "aliyunConfig", "keyId" ]} label={"ACCESS_KEY_ID"}>
                                            <Input.Password visibilityToggle={false}/>
                                        </Form.Item>
                                        <Form.Item rules={[ { required: true, message: '请填写阿里云ACCESS_KEY_SECRET' } ]} name={[ "aliyunConfig", "keySecret" ]} label={"ACCESS_KEY_SECRET"}>
                                            <Input.Password visibilityToggle={false}/>
                                        </Form.Item>
                                    </>
                                );
                            return null;
                        }
                    }
                </Form.Item>
                <Form.Item style={{ textAlign: "center" }}>
                    <Button type="primary" htmlType="submit" onClick={onSubmitHandler}>
                        {settingsId ? "修改" : "创建"}
                    </Button>
                </Form.Item>
            </Form>
        </SettingsPageWrap>
    );
};
export default SettingsPage;
