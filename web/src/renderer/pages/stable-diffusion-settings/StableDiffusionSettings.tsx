import styled from "styled-components";
import { ModalForm, ProColumns, ProForm, ProFormText, ProTable } from "@ant-design/pro-components";
import { useContext, useEffect, useState } from "react";
import { Button, Form, UploadProps } from "antd";
import { StableDiffusionNegativePromptResponse } from "renderer/api/response/stableDiffusionResponse";
import { stableDiffusionApi } from "renderer/api";
import VanillaUploadJson from "renderer/components/json-edit/VanillaUploadJson";
import { Content, OnChange, TextContent } from "vanilla-jsoneditor";
import { ReactSmoothScrollbar } from "renderer/components/smooth-scroll/SmoothScroll";
import { ProjectDetailResponse } from "renderer/api/response/projectResponse";
import { AppGlobalContext } from "renderer/shared/context/appGlobalContext";

const StableDiffusionSettingsWrap = styled.div`
`;

const StableDiffusionSettings = () => {
    const [ settingsList, setSettingsList ] = useState<StableDiffusionNegativePromptResponse[]>([]);
    const [ form ] = Form.useForm();
    const [ content, setContent ] = useState<Content>({
        json: {},
        text: undefined
    });
    const [ editThemeFormatRight, setEditThemeFormatRight ] = useState<boolean>(true);
    const [ modalOpen, setModalOpen ] = useState<boolean>(false);
    const [ settingsDetail, setSettingsDetail ] = useState<StableDiffusionNegativePromptResponse>();
    const { openMessageBox } = useContext(AppGlobalContext);

    const handleUpload: UploadProps["onChange"] = async ({ fileList }) => {
        const [ file ] = fileList;
        const json = await file.originFileObj?.text();
        setContent({ json: JSON.parse(json!) });
    };

    const handleDownload = () => {
        const file = new File([ `${JSON.stringify((content as any)?.json)}` ], "stable-diffusion-api.json", {
            type: "text/json; charset=utf-8;",
        });
        const tmpLink = document.createElement("a");
        const objectUrl = URL.createObjectURL(file);

        tmpLink.href = objectUrl;
        tmpLink.download = file.name;
        document.body.appendChild(tmpLink);
        tmpLink.click();

        document.body.removeChild(tmpLink);
        URL.revokeObjectURL(objectUrl);
    };

    const getSettingsList = async () => {
        const list = await stableDiffusionApi.getStableDiffusionSettingsList({
            page: 1,
            pageSize: -1,
        });
        setSettingsList(list.list);
    };

    const handleChange: OnChange = (newContent, _, status) => {
        setContent(newContent as { text: string });
        if (status?.contentErrors && Object.keys(status.contentErrors).length > 0) {
            setEditThemeFormatRight(false);
        } else {
            setEditThemeFormatRight(true);
        }
    };

    const onSettingsOkHandler = (values: Partial<ProjectDetailResponse>): Promise<boolean> => {
        return new Promise(resolve => {
            if (!editThemeFormatRight) {
                openMessageBox({ type: "error", message: "json错误!" });
                return;
            }
            const data = (content as TextContent).text;
            (settingsDetail ? stableDiffusionApi.updateStableDiffusionSettings : stableDiffusionApi.createStableDiffusionSettings)(Object.assign({
                text: data,
                ...values,
            }, settingsDetail ? {
                id: settingsDetail?.id
            } : {})).then(res => {
                openMessageBox({ type: "success", message: "更新配置成功" });
                getSettingsList();
                resolve(true);
            });
        });
    };

    const columns: ProColumns<StableDiffusionNegativePromptResponse>[] = [
        {
            dataIndex: "index",
            title: "序号",
            align: "center",
            width: 100,
            render(_, _1, index) {
                return (
                    <span>{index + 1}</span>
                );
            }
        },
        {
            dataIndex: "name",
            title: "名称",
        },
        {
            dataIndex: "text",
            title: "内容",
            ellipsis: true,
        },
        {
            title: "操作",
            align: "center",
            valueType: 'index',
            fixed: "right",
            render(text, record, _, action, s) {
                return (
                    <>
                        <Button
                            type={"link"}
                            onClick={() => {
                                setSettingsDetail(record);
                                form.setFieldsValue({ ...record });
                                setContent({
                                    json: JSON.parse(record.text ?? "{}"),
                                    text: undefined,
                                });
                                setModalOpen(true);
                            }}
                        >
                            编辑
                        </Button>
                        <Button
                            type={"link"}
                            danger
                            onClick={async () => {
                                await stableDiffusionApi.deleteStableDiffusionSettings({
                                    ids: [ record.id ]
                                });
                                await getSettingsList();
                            }}
                        >
                            删除
                        </Button>
                    </>
                );
            }
        },
    ];

    useEffect(() => {
        getSettingsList();
    }, []);

    useEffect(() => {
        if (!modalOpen) {
            setSettingsDetail(undefined);
        }
    }, [ modalOpen ]);

    return (
        <StableDiffusionSettingsWrap>
            <ProTable
                rowKey={"id"}
                dataSource={settingsList ?? []}
                columns={columns}
                pagination={false}
                search={false}
                toolBarRender={() => [
                    <ModalForm
                        open={modalOpen}
                        title="创建/修改配置"
                        trigger={
                            <Button type="primary" onClick={() => setModalOpen(true)}>
                                创建/修改配置
                            </Button>
                        }
                        onOpenChange={setModalOpen}
                        form={form}
                        onFinish={onSettingsOkHandler}
                    >
                        <ReactSmoothScrollbar style={{ maxHeight: "calc(100vh - 320px)" }}>
                            <ProForm.Group>
                                <ProFormText
                                    width="md"
                                    name="name"
                                    label="配置名称"
                                    placeholder="请输入配置名称"
                                    rules={[ { required: true, message: '请输入配置名称' } ]}
                                />
                            </ProForm.Group>
                            <VanillaUploadJson
                                content={content}
                                onChange={handleChange}
                                onImportHandler={handleUpload}
                                onExportHandler={handleDownload}
                            />
                        </ReactSmoothScrollbar>
                    </ModalForm>,
                ]}
            />
        </StableDiffusionSettingsWrap>
    );
};

export default StableDiffusionSettings;
