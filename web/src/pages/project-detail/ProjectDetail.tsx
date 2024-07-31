import { useLocation } from "react-router";
import styled from "styled-components";
import { Button, Divider, Form, message, Tooltip, Upload, UploadProps } from "antd";
import { useEffect, useState } from "react";
import { extractTheCharacterProjectDetailParticipleList, getProjectDetail, getProjectDetailInfo, translateProjectDetailParticipleList, updateProjectDetail, updateProjectDetailInfo, uploadProjectDetail } from "../../api/projectApi.ts";
import { ProjectDetailInfo, ProjectDetailResponse } from "../../api/response/projectResponse.ts";
import { EditableProTable, ModalForm, ProColumns, ProForm, ProFormDigit } from "@ant-design/pro-components";
import { Content, OnChange, TextContent } from "vanilla-jsoneditor";
import { SettingOutlined } from "@ant-design/icons";
import { stableDiffusionText2Image } from "../../api/stableDiffusionApi.ts";
import { blobToFile, dataURLtoBlob } from "../../utils/utils.ts";
import { uploadFile } from "../../api/fileApi.ts";
import { FileResponse } from "../../api/response/fileResponse.ts";
import { baseURL } from "../../utils/request.ts";
import VanillaUploadJson from "../../components/json-edit/VanillaUploadJson.tsx";
import { RcFile } from "antd/lib/upload";

const ProjectDetailPageWrap = styled.div`
`;

const ProjectDetailPage = () => {
    const location = useLocation();
    const [ form ] = Form.useForm();
    const [ projectDetail, setProjectDetail ] = useState<ProjectDetailResponse>();
    const [ content, setContent ] = useState<Content>({
        json: JSON.stringify({}),
        text: undefined
    });
    const [ editThemeFormatRight, setEditThemeFormatRight ] = useState<boolean>(true);
    const state = location.state;
    /**
     * 获取项目
     * @param id
     */
    const getProjectDetailConfig = async (id: number) => {
        const detail = await getProjectDetail({
            projectId: id
        });
        setProjectDetail(detail);
        form.setFieldsValue({ ...detail });
    };
    /**
     * 项目详情配置
     * @param values
     */
    const onSettingsOkHandler = async (values: Partial<ProjectDetailResponse>) => {
        if (projectDetail) {
            if (!editThemeFormatRight) {
                message.error("json错误!");
                return;
            }
            const data = (content as TextContent).text;
            await updateProjectDetail({
                id: projectDetail.id,
                ...values,
                stableDiffusionConfig: data
            });
            message.success("更新配置成功");
            await getProjectDetailConfig(state.id);
            return true;
        }
    };
    /**
     * 上传文本
     * @param file
     */
    const onBeforeUpload = async (file: RcFile) => {
        if (!projectDetail) {
            return;
        }
        const maxWords = projectDetail?.participleConfig?.maxWords;
        const minWords = projectDetail?.participleConfig?.minWords;
        await uploadProjectDetail({
            id: projectDetail.id,
            file: file,
            maxWords,
            minWords
        });
        await getProjectDetailConfig(state.id);
        return false;
    };
    /**
     * 进行人物提取
     */
    const extractTheRole = async () => {
        if (!projectDetail) return;
        await extractTheCharacterProjectDetailParticipleList({
            id: projectDetail?.id
        });
        await getProjectDetailConfig(state.id);
    };
    /**
     * 进行翻译
     * @param data
     */
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
        setContent({ json: JSON.parse(!projectDetail?.stableDiffusionConfig ? "{}" : projectDetail?.stableDiffusionConfig) });
    };

    const handleChange: OnChange = (newContent, _, status) => {
        setContent(newContent as { text: string });
        if (status?.contentErrors && Object.keys(status.contentErrors).length > 0) {
            setEditThemeFormatRight(false);
        } else {
            setEditThemeFormatRight(true);
        }
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
    /**
     * 进行文本转图片
     */
    const text2imageHandler = async () => {
        const ids = projectDetail?.projectDetailInfoList?.map(i => i.id) ?? [];
        const projectDetailStableDiffusionConfig = projectDetail?.stableDiffusionConfig ?? "{}";
        console.log(!!projectDetailStableDiffusionConfig);
        for (const id of ids) {
            const stableDiffusionImages: ProjectDetailInfo["stableDiffusionImages"] = [];
            const data = await getProjectDetailInfo({ id });
            const stableDiffusionParams: {
                [key: string]: any
            } = {};
            stableDiffusionParams["prompt"] = data.prompt;
            stableDiffusionParams["negativePrompt"] = data.negativePrompt;
            const jsonConfig = JSON.parse(!projectDetailStableDiffusionConfig ? "{}" : projectDetailStableDiffusionConfig);
            for (const key in jsonConfig) {
                stableDiffusionParams[key] = jsonConfig[key];
            }
            const images = await stableDiffusionText2Image({
                id,
                projectDetailId: projectDetail?.id,
            });
            if (images.length) {
                for (let i = 0; i < images.length; i++) {
                    const image = images[i];
                    const blob = dataURLtoBlob(`data:image/png;base64,${image}`);
                    const file = blobToFile(blob, `${id}-${i}.png`);
                    const upload = await uploadFile({
                        file,
                        fileType: "stable-diffusion"
                    });
                    stableDiffusionImages.push({
                        projectDetailInfoId: id,
                        name: upload.name,
                        key: upload.key,
                        url: upload.url,
                        tag: upload.tag,
                    });
                }
                await updateProjectDetailInfo({
                    id: id,
                    stableDiffusionImages,
                });
            }
            await getProjectDetailConfig(state.id);
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
            dataIndex: "stableDiffusionImages",
            title: "图片列表",
            editable: false,
            width: 100,
            render(values: FileResponse[]) {
                console.log(values);
                return (
                    <div style={{ width: "100%" }}>
                        <img width={100} src={`${baseURL}/${values?.[0]?.url}`} alt=""/>
                    </div>
                );
            }
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
        if (state.id) {
            getProjectDetailConfig(state.id);
        }
    }, [ state ]);

    if (!projectDetail) return null;

    return (
        <ProjectDetailPageWrap>
            <EditableProTable
                rowKey={"id"}
                editable={{
                    onSave: async (_, data) => {
                        await updateProjectDetailInfo({
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
                    <Upload
                        beforeUpload={onBeforeUpload}
                        maxCount={1}
                        showUploadList={false}
                        accept={".txt"}
                    >
                        <Button>
                            上传文件
                        </Button>
                    </Upload>,
                    <Button onClick={extractTheRole}>
                        角色提取
                    </Button>,
                    <Button onClick={() => translatePrompt({ projectDetailId: projectDetail?.id })}>
                        翻译
                    </Button>,
                    <Button onClick={text2imageHandler}>
                        生成图片
                    </Button>,
                    <ModalForm
                        title="配置参数"
                        trigger={
                            <Tooltip title={"配置stable-diffusion请求参数"}>
                                <SettingOutlined onClick={setStableDiffusionJson}/>
                            </Tooltip>
                        }
                        form={form}
                        onFinish={onSettingsOkHandler}
                    >
                        <ProForm.Group>
                            <ProFormDigit
                                width="md"
                                name={[ "participleConfig", "minWords" ]}
                                label="最小文字数量"
                                placeholder="请输入最小文字数量"
                                min={10}
                                rules={[ { required: true, message: '请输入最小文字数量' } ]}
                            />

                            <ProFormDigit
                                width="md"
                                name={[ "participleConfig", "maxWords" ]}
                                label="最大文字数量"
                                placeholder="请输入最大文字数量"
                                min={10}
                                rules={[ { required: true, message: '请输入最大文字数量' } ]}
                            />
                        </ProForm.Group>
                        <Divider orientation="left">stable-diffusion配置</Divider>
                        <VanillaUploadJson
                            content={content}
                            onChange={handleChange}
                            onImportHandler={handleUpload}
                            onExportHandler={handleDownload}
                        />
                    </ModalForm>
                ]}
            />
        </ProjectDetailPageWrap>
    );
};

export default ProjectDetailPage;
