import { useLocation } from "react-router";
import styled from "styled-components";
import { Button, Form, message, Tooltip, UploadFile, UploadProps } from "antd";
import { useEffect, useState } from "react";
import { extractTheCharacterProjectDetailParticipleList, getProjectDetail, getProjectDetailInfo, translateProjectDetailParticipleList, updateProjectDetail, updateProjectDetailParticipleList, uploadProjectDetail } from "../../api/projectApi.ts";
import { ProjectDetailInfo, ProjectDetailResponse } from "../../api/response/projectResponse.ts";
import { EditableProTable, ModalForm, ProColumns, ProForm, ProFormDigit, ProFormUploadButton } from "@ant-design/pro-components";
import { Content, OnChange, TextContent } from "vanilla-jsoneditor";
import ThemeVanillaJsonModal from "../../components/json-edit/ThemeVanillaJsonModal.tsx";
import { SettingOutlined } from "@ant-design/icons";
import { stableDiffusionText2Image } from "../../api/stableDiffusionApi.ts";

const ProjectDetailPageWrap = styled.div`
`;

const ProjectDetailPage = () => {
    const location = useLocation();
    const [ form ] = Form.useForm();
    const [ projectDetail, setProjectDetail ] = useState<ProjectDetailResponse>();
    const [ openStableDiffusionConfig, setOpenStableDiffusionConfig ] = useState<boolean>(false);
    const [ content, setContent ] = useState<Content>({
        json: JSON.stringify({}),
        text: undefined
    });
    const [ editThemeFormatRight, setEditThemeFormatRight ] = useState<boolean>(true);
    const state = location.state;

    const getProjectDetailConfig = async (id: number) => {
        const detail = await getProjectDetail({
            projectId: id
        });
        setProjectDetail(detail);
        form.setFieldsValue({ ...detail.participleConfig });
    };

    const onUploadOkHandler = async (values: Partial<ProjectDetailResponse & {
        file: UploadFile[];
    }>) => {
        const { file, ...args } = values;
        if (projectDetail) {
            await uploadProjectDetail({
                id: projectDetail.id,
                file: file?.[0]?.originFileObj,
                ...args,
            });
            await getProjectDetailConfig(state.id);
        }
    };


    const extractTheCharacter = async () => {
        if (!projectDetail) return;
        await extractTheCharacterProjectDetailParticipleList({
            id: projectDetail?.id
        });
        await getProjectDetailConfig(state.id);
    };

    const translatePrompt = async (data: {
        id?: number;
        projectDetailId?: number;
    }) => {
        if (!projectDetail) return;
        await translateProjectDetailParticipleList(data);
        message.success("翻译成功");
        await getProjectDetailConfig(state.id);
    };

    const setStableDiffusionJson = () => {
        setContent({ json: JSON.parse(projectDetail?.stableDiffusionConfig ?? "{}") });
    };

    const handleChange: OnChange = (newContent, _, status) => {
        setContent(newContent as { text: string });
        if (status?.contentErrors && Object.keys(status.contentErrors).length > 0) {
            setEditThemeFormatRight(false);
        } else {
            setEditThemeFormatRight(true);
        }
    };

    const onJsonOkHandler = async () => {
        if (!editThemeFormatRight) {
            message.error("json错误!");
            return;
        }
        const data = (content as TextContent).text;
        await updateProjectDetail({
            id: projectDetail?.id,
            stableDiffusionConfig: data
        });
        message.success("更新成功");
        setOpenStableDiffusionConfig(false);
        await getProjectDetailConfig(state.id);
    };

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

    const text2imageHandler = async () => {
        const ids = projectDetail?.projectDetailInfoList?.map(i => i.id) ?? [];
        const projectDetailStableDiffusionConfig = projectDetail?.stableDiffusionConfig ?? "{}";
        for (const id of ids) {
            const data = await getProjectDetailInfo({ id });
            const stableDiffusionParams: {
                [key: string]: any
            } = {};
            stableDiffusionParams["prompt"] = data.prompt;
            stableDiffusionParams["negativePrompt"] = data.negativePrompt;
            const jsonConfig = JSON.parse(projectDetailStableDiffusionConfig);
            for (const key in jsonConfig) {
                stableDiffusionParams[key] = jsonConfig[key];
            }
            const stableDiffusionData = await stableDiffusionText2Image(stableDiffusionParams);
            console.log(stableDiffusionData);
        }
        // await stableDiffusionText2Image({ ids, projectDetailId: projectDetail?.id ?? 0 });
    };

    const columns: ProColumns<ProjectDetailInfo>[] = [
        {
            dataIndex: "text",
            title: "文本",
            valueType: "textarea",
            width: 300,
            fixed: "left",
        },
        {
            dataIndex: "prompt",
            title: "正向提示词",
            valueType: "textarea",
            width: 300
        },
        {
            dataIndex: "negativePrompt",
            title: "反向提示词",
            valueType: "textarea",
            width: 300
        },
        {
            dataIndex: "role",
            title: "人物",
            valueType: "textarea",
            width: 100,
            tooltip: '多个人物名称通过","拼接',
        },
        {
            title: "操作",
            align: "center",
            width: 170,
            valueType: 'option',
            fixed: "right",
            render(text, record, _, action) {
                return (
                    <>
                        <Button type={"link"} onClick={() => {
                            action?.startEditable?.(record.id);
                        }}>
                            编辑
                        </Button>
                        <Button type={"link"} onClick={() => translatePrompt({ id: record?.id })}>
                            翻译
                        </Button>
                    </>
                );
            }
        },
    ];

    useEffect(() => {
        if (openStableDiffusionConfig) {
            setStableDiffusionJson();
        }
    }, [ openStableDiffusionConfig ]);

    useEffect(() => {
        if (state.id) {
            getProjectDetailConfig(state.id);
        }
    }, [ state ]);

    return (
        <ProjectDetailPageWrap>
            <ThemeVanillaJsonModal
                title={"stable-diffusion配置"}
                open={openStableDiffusionConfig}
                onCancel={() => setOpenStableDiffusionConfig(false)}
                onOk={onJsonOkHandler}
                content={content}
                onChange={handleChange}
                onImportHandler={handleUpload}
                onExportHandler={handleDownload}
            />
            <EditableProTable
                rowKey={"id"}
                editable={{
                    onSave: async (_, data) => {
                        await updateProjectDetailParticipleList({
                            ...data
                        });
                        message.success("保存成功");
                        await getProjectDetailConfig(state.id);
                    },
                }}
                recordCreatorProps={false}
                value={projectDetail?.projectDetailInfoList ?? []}
                columns={columns}
                virtual={true}
                scroll={{ y: 650, x: 800 }}
                headerTitle={projectDetail?.fileName}
                pagination={false}
                search={false}
                toolBarRender={() => [
                    <ModalForm
                        disabled={!projectDetail}
                        title="导入文件"
                        trigger={
                            <Button type="primary">
                                上传文本
                            </Button>
                        }
                        form={form}
                        onFinish={onUploadOkHandler}
                    >
                        <ProForm.Group>
                            <ProFormDigit
                                width="md"
                                name="minWords"
                                label="最小文字数量"
                                placeholder="请输入最小文字数量"
                                min={10}
                                rules={[ { required: true, message: '请输入最小文字数量' } ]}
                            />

                            <ProFormDigit
                                width="md"
                                name="maxWords"
                                label="最大文字数量"
                                placeholder="请输入最大文字数量"
                                min={10}
                                rules={[ { required: true, message: '请输入最大文字数量' } ]}
                            />
                        </ProForm.Group>
                        <ProForm.Group>
                            <ProFormUploadButton
                                width="md"
                                fieldProps={{
                                    name: 'file',
                                    beforeUpload: () => false,
                                    maxCount: 1,
                                }}
                                name="file"
                                label="文件"
                                placeholder="文件"
                                rules={[ { required: true, message: '请上传文件' } ]}
                                accept={".txt"}
                            />
                        </ProForm.Group>
                    </ModalForm>,
                    <Button onClick={extractTheCharacter}>
                        角色提取
                    </Button>,
                    <Button onClick={() => translatePrompt({ projectDetailId: projectDetail?.id })}>
                        翻译
                    </Button>,
                    <Button onClick={text2imageHandler}>
                        生成图片
                    </Button>,
                    <Tooltip title={"配置stable-diffusion请求参数"}>
                        <SettingOutlined onClick={() => setOpenStableDiffusionConfig(true)}/>
                    </Tooltip>,
                ]}
            />
        </ProjectDetailPageWrap>
    );
};

export default ProjectDetailPage;
