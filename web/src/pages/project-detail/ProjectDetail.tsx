import { useLocation } from "react-router";
import styled from "styled-components";
import { Button, Divider, Form, message, Radio, Space, Tooltip, Upload, UploadProps } from "antd";
import { useEffect, useRef, useState } from "react";
import { deleteInfo, extractTheCharacterProjectDetailParticipleList, getProjectDetail, getProjectDetailInfo, translateProjectDetailParticipleList, updateProjectDetail, updateProjectDetailInfo, uploadProjectDetail } from "../../api/projectApi.ts";
import { Info, ProjectDetailResponse } from "../../api/response/projectResponse.ts";
import { EditableProTable, ModalForm, ProColumns, ProForm, ProFormDigit, ProFormSelect, ProFormText } from "@ant-design/pro-components";
import { Content, OnChange, TextContent } from "vanilla-jsoneditor";
import { CloseOutlined, NotificationOutlined, SettingOutlined } from "@ant-design/icons";
import { stableDiffusionText2Image } from "../../api/stableDiffusionApi.ts";
import { blobToFile, dataURLtoBlob } from "../../utils/utils.ts";
import { uploadFile } from "../../api/fileApi.ts";
import { FileResponse } from "../../api/response/fileResponse.ts";
import { baseURL } from "../../utils/request.ts";
import VanillaUploadJson from "../../components/json-edit/VanillaUploadJson.tsx";
import { RcFile } from "antd/lib/upload";
import { audioList } from "../../utils/audio-list.ts";
import { createAudioSrt } from "../../api/audioSrt.ts";

const ProjectDetailPageWrap = styled.div`
`;

const ImagesActionWrap = styled.div`
    .ant-radio-wrapper {
        position: relative;

        &:hover {
            .action-delete {
                display: block;
            }
        }

        .action-delete {
            position: absolute;
            right: 10px;
            top: 0;
            color: red;
            cursor: pointer;
            display: none;
        }
    }


`;

const percentageRegex = /^[-+]\d+%$/;
const pitchRegex = /^[-+]\d+Hz$/;

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
    const mp3Ref = useRef<HTMLAudioElement | null>(null);
    /**
     * 获取项目
     * @param id
     */
    const getProjectDetailConfig = async (id: number) => {
        const detail = await getProjectDetail({
            projectId: id
        });
        setProjectDetail(detail);
        form.setFieldsValue({
            ...detail,
        });
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
            if (!percentageRegex.test(values?.audioConfig?.rate ?? "0")) {
                message.error("请输入正确的语速值!");
                return;
            }
            if (!percentageRegex.test(values?.audioConfig?.volume ?? "0")) {
                message.error("请输入正确的音量值!");
                return;
            }
            if (!pitchRegex.test(values?.audioConfig?.pitch ?? "0")) {
                message.error("请输入正确的分贝值!");
                return;
            }
            await updateProjectDetail({
                id: projectDetail.id,
                stableDiffusionConfig: data,
                ...values,
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
        const projectDetailStableDiffusionConfig = projectDetail?.stableDiffusionConfig ?? "{}";
        for (const info of (projectDetail?.infoList ?? [])) {
            let selectedId: null | number = null;
            const stableDiffusionImages: Info["stableDiffusionImages"] = [];
            const data = await getProjectDetailInfo({ id: info.id });
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
                id: info.id,
                projectDetailId: projectDetail?.id,
            });
            if (images.length) {
                for (let i = 0; i < images.length; i++) {
                    const image = images[i];
                    const blob = dataURLtoBlob(`data:image/png;base64,${image}`);
                    const file = blobToFile(blob, `stable-diffusion-${info.id}-${i}.png`);
                    const upload = await uploadFile({
                        file,
                        fileType: "stable-diffusion"
                    });
                    if (i === 0) {
                        selectedId = upload.id;
                    }
                    stableDiffusionImages.push({
                        InfoId: info.id,
                        name: upload.name,
                        key: upload.key,
                        url: upload.url,
                        tag: upload.tag,
                        fileId: upload.id,
                    });
                }
                await updateProjectDetailInfo(Object.assign({
                    id: info.id,
                    stableDiffusionImages,
                }, (selectedId && !info.stableDiffusionImageId) ? {
                    stableDiffusionImageId: selectedId
                } : {}));
            }
            await getProjectDetailConfig(state.id);
        }
    };

    const onStableDiffusionImagesOnChange = async (selectedId: number, infoId: number) => {
        await updateProjectDetailInfo(Object.assign({
            id: infoId,
            stableDiffusionImageId: selectedId
        }));
        await getProjectDetailConfig(state.id);
    };
    /**
     * 生成音频和字幕
     */
    const createAudioAndSrtHandler = async () => {
        if (projectDetail)
            await createAudioSrt({
                id: projectDetail?.id
            });
    };

    const onDetailProjectInfo = async (id: number) => {
        await deleteInfo({ id });
        await getProjectDetailConfig(state.id);
    };

    const columns: ProColumns<Info>[] = [
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
            width: 200,
            // @ts-ignore
            render(values: FileResponse[], record) {
                return (
                    <ImagesActionWrap>
                        <Radio.Group onChange={(e) => onStableDiffusionImagesOnChange(e.target.value, record.id)} value={record.stableDiffusionImageId}>
                            <Space direction="vertical">
                                {
                                    values.map(file => {
                                        return (
                                            <Radio key={file.id} value={file.id}>
                                                <span className={"action-delete"} onClick={(e) => {
                                                    e.preventDefault();
                                                    e.stopPropagation();
                                                }}>
                                                    <CloseOutlined/>
                                                </span>
                                                <img width={150} src={`${baseURL}/${file?.url}`} alt=""/>
                                            </Radio>
                                        );
                                    })
                                }
                            </Space>
                        </Radio.Group>
                    </ImagesActionWrap>
                );
            }
        },
        {
            title: "操作",
            align: "center",
            valueType: 'option',
            fixed: "right",
            // @ts-ignore
            render(text, record, _, action) {
                return (
                    <>
                        <Button danger type={"link"} onClick={() => onDetailProjectInfo(record.id)}>
                            删除
                        </Button>
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
            <audio ref={mp3Ref} style={{ display: "none" }}/>
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
                value={projectDetail?.infoList ?? []}
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
                    <Button disabled={!projectDetail?.infoList?.length} onClick={extractTheRole}>
                        角色提取
                    </Button>,
                    <Button disabled={!projectDetail?.infoList?.length} onClick={() => translatePrompt({ projectDetailId: projectDetail?.id })}>
                        翻译
                    </Button>,
                    <Button disabled={!projectDetail?.infoList?.length} onClick={createAudioAndSrtHandler}>
                        生成音频和字幕
                    </Button>,
                    <Button disabled={!projectDetail?.infoList?.length} onClick={text2imageHandler}>
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
                        <ProForm.Group title={"分词配置"}>
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
                        <ProForm.Group title={"音频设置"}>
                            <ProFormDigit
                                width="md"
                                name={[ "audioConfig", "srtLimit" ]}
                                label="字幕最大长度"
                                placeholder="请输入最大文字数量"
                                min={10}
                            />
                            <ProFormText
                                width="md"
                                name={[ "audioConfig", "rate" ]}
                                label="音频语速"
                                placeholder="请输入音频语速"
                                tooltip={"格式为(+-)0%"}
                            />
                            <ProFormText
                                width="md"
                                name={[ "audioConfig", "volume" ]}
                                label="音量"
                                tooltip={"格式为(+-)0%"}
                                placeholder="请输入音量"
                            />
                            <ProFormText
                                width="md"
                                name={[ "audioConfig", "pitch" ]}
                                label="分贝"
                                tooltip={"格式为(+-)0Hz"}
                                placeholder="请输入分贝"
                            />
                            <ProFormSelect
                                width="md"
                                name={[ "audioConfig", "voice" ]}
                                label="音频角色"
                                placeholder="请选择音频角色"
                                options={audioList}
                                fieldProps={{
                                    optionRender(option) {
                                        return (
                                            <Space>
                                                <div style={{ width: '265px', overflow: "hidden", textOverflow: "ellipsis" }}>
                                                    {option.data.name + "-" + option.data.value}
                                                </div>
                                                {option.data.mp3 && <NotificationOutlined onClick={(e) => {
                                                    e.stopPropagation();
                                                    if (mp3Ref.current) {
                                                        mp3Ref.current.src = option.data.mp3;
                                                        mp3Ref.current.play();
                                                    }
                                                }}/>}
                                            </Space>
                                        );
                                    }
                                }}
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
