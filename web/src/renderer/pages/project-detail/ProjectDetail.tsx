import { useLocation, useNavigate } from "react-router";
import styled from "styled-components";
import { Button, Divider, Form, Input, Radio, Space, Spin, Tooltip, Upload, UploadProps } from "antd";
import { useContext, useEffect, useRef, useState } from "react";
import { EditableProTable, ModalForm, ProColumns, ProForm, ProFormDigit, ProFormSelect, ProFormText, ProFormUploadButton } from "@ant-design/pro-components";
import { Content, OnChange, TextContent } from "vanilla-jsoneditor";
import { CloseOutlined, ExclamationOutlined, LoadingOutlined, NotificationOutlined, SettingOutlined } from "@ant-design/icons";
import { Info, ProjectDetailResponse } from "renderer/api/response/projectResponse";
import { stableDiffusionText2Image } from "renderer/api/stableDiffusionApi";
import { blobToFile, dataURLtoBlob } from "renderer/utils/utils";
import { uploadFile } from "renderer/api/fileApi";
import { createAudioSrt } from "renderer/api/audioSrtApi";
import { host } from "renderer/request/request";
import { audioList } from "renderer/utils/audio-list";
import VanillaUploadJson from "renderer/components/json-edit/VanillaUploadJson";
import { infoApi, projectDetailApi, stableDiffusionApi, videoApi } from "renderer/api";
import { AppGlobalContext } from "renderer/shared/context/appGlobalContext";
import { FileResponse } from "renderer/api/response/fileResponse";
import { RcFile } from "antd/lib/upload";
import { navBarHeight } from "renderer/shared";
import { ReactSmoothScrollbar } from "renderer/components/smooth-scroll/SmoothScroll";
import { ipcApi } from "renderer/ipc/BasicRendererIpcAdapter";
import FileSvg from "renderer/assets/svg-com/file.svg";

const ProjectDetailPageWrap = styled.div`
`;

const ImagesActionWrap = styled.div`
    display: flex;
    align-items: center;
    flex-direction: column;

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

const AudioActionWrap = styled.div``;
const AudioActionItemWrap = styled.div`
    display: flex;
    align-items: center;
    justify-content: space-between;
`;

const percentageRegex = /^[-+]\d+(%%?)$/;
const pitchRegex = /^[+]\d+Hz$/;

const ProjectDetailPage = () => {
    const location = useLocation();
    const navigate = useNavigate();
    const [ form ] = Form.useForm();
    const [ currentAudioForm ] = Form.useForm();
    const [ uploadForm ] = Form.useForm();
    const [ projectDetail, setProjectDetail ] = useState<ProjectDetailResponse>();
    const [ content, setContent ] = useState<Content>({
        json: JSON.stringify({}),
        text: undefined
    });
    const [ editThemeFormatRight, setEditThemeFormatRight ] = useState<boolean>(true);
    const [ tableHeight, setTableHeight ] = useState<number>(350);
    const [ tableLoading, setTableLoading ] = useState<boolean>(true);
    const [ renderLoading, setRenderLoading ] = useState<boolean>(false);
    const { openMessageBox } = useContext(AppGlobalContext);
    const state = location.state;
    const mp3Ref = useRef<HTMLAudioElement | null>(null);
    const intervalRef = useRef<NodeJS.Timeout | null>(null);
    // const [ fontColor, setFontColor ] = useState<string>();

    /**
     * 获取项目
     * @param id
     */
    const getProjectDetailConfig = async (id: number) => {
        setTableLoading(true);
        const detail = await projectDetailApi.getProjectDetail({
            id: id
        });
        setProjectDetail(detail);
        form.setFieldsValue({
            ...detail,
        });
        setTableLoading(false);
    };
    /**
     * 启动查询任务
     */
    const removeInterval = () => {
        if (!intervalRef.current) {
            return;
        }
        clearInterval(intervalRef.current);
        intervalRef.current = null;
    };
    /**
     * 启动查询任务
     */
    const startInterval = () => {
        if (intervalRef.current) {
            openMessageBox({ type: "warning", message: "有正在执行的任务，请稍后执行。" });
            return;
        }
        intervalRef.current = setInterval(() => {
            getProjectDetailConfig(state.id);
        }, 1000);
    };
    /**
     * 项目详情配置
     * @param values
     */
    const onSettingsOkHandler = (values: Partial<ProjectDetailResponse>): Promise<boolean> => {
        return new Promise(resolve => {
            if (projectDetail) {
                if (!editThemeFormatRight) {
                    openMessageBox({ type: "error", message: "json错误!" });
                    return;
                }
                const data = (content as TextContent).text;
                if (!percentageRegex.test(values?.audioConfig?.rate ?? "0")) {
                    openMessageBox({ type: "error", message: "请输入正确的语速值!" });
                    resolve(false);
                    return;
                }
                if (!percentageRegex.test(values?.audioConfig?.volume ?? "0")) {
                    openMessageBox({ type: "error", message: "请输入正确的音量值!" });
                    resolve(false);
                    return;
                }
                if (!pitchRegex.test(values?.audioConfig?.pitch ?? "0")) {
                    openMessageBox({ type: "error", message: "请输入正确的分贝值!" });
                    resolve(false);
                    return;
                }
                projectDetailApi.updateProjectDetail({
                    id: projectDetail.id,
                    stableDiffusionConfig: data,
                    ...values,
                }).then(res => {
                    openMessageBox({ type: "success", message: "更新配置成功" });
                    getProjectDetailConfig(state.id);
                    resolve(true);
                });
            }
        });
    };
    /**
     * 上传文本
     * @param values
     */
    const onUploadOkHandler = async (values: any): Promise<boolean> => {
        return new Promise((resolve) => {
            if (!projectDetail) {
                resolve(true);
                return;
            }
            projectDetailApi.uploadProjectDetail({
                id: projectDetail.id,
                file: values.files?.[0]?.originFileObj as RcFile,
                saveType: values?.saveType,
                whetherParticiple: values?.whetherParticiple
            }).then(async () => {
                openMessageBox({ type: "success", message: "上传成功" });
                resolve(true);
                await getProjectDetailConfig(state.id);
            });
        });
    };
    /**
     * 进行人物提取
     */
    const extractTheRole = async () => {
        if (!projectDetail) return;
        await infoApi.extractTheCharacterProjectDetailParticipleList({
            id: projectDetail?.id
        });
        await getProjectDetailConfig(state.id);
    };
    /**
     * 进行人物提取
     */
    const keywordsExtract = async () => {
        if (!projectDetail) return;
        await infoApi.keywordsExtractInfoList({
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
        setRenderLoading(true);
        await infoApi.translateProjectDetailParticipleList(data);
        setRenderLoading(false);
        openMessageBox({ type: "success", message: "翻译成功" });
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
    const text2imageBatchHandler = async (infoDetail?: Info) => {
        const handler = async (info: Info) => {
            setRenderLoading(true);
            let selectedId: null | number = null;
            const stableDiffusionImages: Info["stableDiffusionImages"] = [];
            const images = await stableDiffusionText2Image({
                ids: [ info.id ],
                projectDetailId: projectDetail?.id,
            }).catch(() => {
                setRenderLoading(false);
                return;
            });
            if (images?.length) {
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
                        projectDetailId: info.projectDetailId,
                        InfoId: info.id,
                        name: upload.name,
                        key: upload.key,
                        url: upload.url,
                        tag: upload.tag,
                        fileId: upload.id,
                    });
                }
                await infoApi.updateProjectDetailInfo(Object.assign({
                    id: info.id,
                    stableDiffusionImages,
                }, (selectedId && !info.stableDiffusionImageId) ? {
                    stableDiffusionImageId: selectedId
                } : {}));
                await getProjectDetailConfig(state.id);
            }
            setRenderLoading(false);
        };
        if (infoDetail) {
            await handler(infoDetail);
        } else {
            for (const info of (projectDetail?.infoList ?? [])) {
                await handler(info);
            }
        }
    };
    const text2imageHandler = async (file: RcFile, info: Info) => {
        setRenderLoading(true);
        const upload = await uploadFile({
            file,
            fileType: "stable-diffusion"
        });
        const data = {
            projectDetailId: info.projectDetailId,
            InfoId: info.id,
            name: upload.name,
            key: upload.key,
            url: upload.url,
            tag: upload.tag,
            fileId: upload.id,
        };
        await stableDiffusionApi.addImage(data);
        if (!info.stableDiffusionImageId) {
            await infoApi.updateProjectDetailInfo({
                id: info.id,
                stableDiffusionImageId: upload.id
            });
        }
        openMessageBox({ type: "success", message: "上传成功" });
        setRenderLoading(false);
        await getProjectDetailConfig(state.id);
        return false;
    };
    /**
     * 生成视频
     */
    const createInfoVideo = async (id?: number) => {
        const ids = id ? [ id ] : [];
        if (projectDetail) {
            setRenderLoading(true);
            await videoApi.createInfoVideo({
                ids,
                projectDetailId: projectDetail?.id,
            });
            setRenderLoading(false);
            openMessageBox({ type: "success", message: "生成成功" });
        }
    };
    /**
     * 改变选中的图片
     * @param selectedId
     * @param infoId
     */
    const onStableDiffusionImagesOnChange = async (selectedId: number, infoId: number) => {
        await infoApi.updateProjectDetailInfo(Object.assign({
            id: infoId,
            stableDiffusionImageId: selectedId
        }));
        await getProjectDetailConfig(state.id);
    };
    /**
     * 生成音频和字幕
     */
    const createAudioAndSrtHandler = async () => {
        if (projectDetail) {
            setRenderLoading(true);
            await createAudioSrt({
                id: projectDetail?.id
            });
            setRenderLoading(false);
            openMessageBox({ type: "success", message: "生成成功" });
        }
    };

    const onDetailProjectInfo = async (id: number) => {
        await infoApi.deleteInfo({ id });
        await getProjectDetailConfig(state.id);
    };

    const onInfoAudioOkHandler = async (values: any, id: number) => {
        if (!percentageRegex.test(values?.rate ?? "0")) {
            openMessageBox({ type: "error", message: "请输入正确的语速值!" });
            return;
        }
        if (!percentageRegex.test(values?.volume ?? "0")) {
            openMessageBox({ type: "error", message: "请输入正确的音量值!" });
            return;
        }
        if (!pitchRegex.test(values?.pitch ?? "0")) {
            openMessageBox({ type: "error", message: "请输入正确的分贝值!" });
            return;
        }
        await infoApi.updateProjectDetailInfo({
            id: id,
            audioConfig: {
                ...values,
            }
        });
        openMessageBox({ type: "success", message: "更新配置成功" });
        await getProjectDetailConfig(state.id);
        return true;
    };

    const getDocumentHeight = () => {
        const documentHeight = document.body.clientHeight;
        setTableHeight(documentHeight - navBarHeight() - 200);
    };

    const onPromptUrlPathSelect = async () => {
        const folderValues = await ipcApi.fileAdapter.onFolderSelect({
            properties: [ "openFile" ],
            filters: [
                {
                    name: "promptUrl",
                    extensions: [ "txt" ]
                },
            ]
        });
        if (!folderValues.data.canceled) {
            const selectPath = folderValues.data.filePaths[0];
            form.setFieldValue("promptUrl", selectPath);
        }
    };

    const onFilePathSelect = async () => {
        const folderValues = await ipcApi.fileAdapter.onFolderSelect({
            properties: [ "openFile" ],
            filters: [
                {
                    name: "font",
                    extensions: [ "ttf" ]
                },
            ]
        });
        if (!folderValues.data.canceled) {
            const selectPath = folderValues.data.filePaths[0];
            form.setFieldValue([ "videoConfig", "fontFile" ], selectPath);
        }
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
            dataIndex: "keywordsText",
            title: "中文关键词",
            valueType: "textarea",
            width: 300
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
            tooltip: '多个人物名称通过","拼接。会自动和loras列表中存在的lora关联到prompt中',
        },
        {
            dataIndex: "stableDiffusionImages",
            title: "图片列表",
            editable: false,
            width: 300,
            render(values, record) {
                return (
                    <ImagesActionWrap>
                        <Radio.Group onChange={(e) => onStableDiffusionImagesOnChange(e.target.value, record.id)} value={record.stableDiffusionImageId}>
                            <Space direction="vertical">
                                {
                                    (values as FileResponse[]).map(file => {
                                        return (
                                            <Radio key={file.id} value={file.fileId}>
                                                <span
                                                    className={"action-delete"} onClick={async (e) => {
                                                        e.preventDefault();
                                                        e.stopPropagation();
                                                        await stableDiffusionApi.stableDiffusionDeleteImage({ ids: [ file.id ] });
                                                        await getProjectDetailConfig(state.id);
                                                    }}>
                                                    <CloseOutlined/>
                                                </span>
                                                <img width={150} src={`${host()}/${file?.url}`} alt=""/>
                                            </Radio>
                                        );
                                    })
                                }
                            </Space>
                        </Radio.Group>
                        <Upload showUploadList={false} maxCount={1} accept={"image/*"} beforeUpload={(file) => text2imageHandler(file, record)}>
                            <Button>上传图片</Button>
                        </Upload>
                    </ImagesActionWrap>
                );
            }
        },
        {
            dataIndex: "audioConfig",
            title: () => {
                return (
                    <div style={{ display: "flex", alignItems: "center", justifyContent: "space-between" }}>
                        <span>音频设置</span>
                        <Button
                            type={"link"} onClick={async () => {
                                setRenderLoading(true);
                                await infoApi.updateAudio({ projectDetailId: projectDetail?.id });
                                openMessageBox({ type: "success", message: "一键设置成功" });
                                await getProjectDetailConfig(state.id);
                                setRenderLoading(false);
                            }}>一键设置</Button>
                    </div>
                );
            },
            width: 300,
            editable: false,
            render(values, record) {
                const currentAudio = audioList.find(audio => audio.value === record?.audioConfig?.voice);
                return <ModalForm
                    key={"settings"}
                    title="配置参数"
                    trigger={
                        <AudioActionWrap
                            onClick={() => {
                                currentAudioForm.setFieldsValue({
                                    ...(values as Info["audioConfig"])
                                });
                            }}>
                            <AudioActionItemWrap>
                                <span>声音: </span>
                                <span>{record?.audioConfig?.voice}{currentAudio && currentAudio.mp3 && <NotificationOutlined
                                    style={{
                                        marginLeft: "8px"
                                    }}
                                    onClick={(e) => {
                                        e.stopPropagation();
                                        if (mp3Ref.current) {
                                            mp3Ref.current.src = currentAudio.mp3;
                                            mp3Ref.current.play();
                                        }
                                    }}
                                />}</span>
                            </AudioActionItemWrap>
                            <AudioActionItemWrap>
                                <span>语速: </span>
                                <span>{record?.audioConfig?.rate}</span>
                            </AudioActionItemWrap>
                            <AudioActionItemWrap>
                                <span>音量: </span>
                                <span>{record?.audioConfig?.volume}</span>
                            </AudioActionItemWrap>
                            <AudioActionItemWrap>
                                <span>分贝: </span>
                                <span>{record?.audioConfig?.pitch}</span>
                            </AudioActionItemWrap>
                        </AudioActionWrap>
                    }
                    form={currentAudioForm}
                    onFinish={(values: any) => onInfoAudioOkHandler(values, record.id)}
                >
                    <ProForm.Group>
                        <ProFormText
                            width="md"
                            name={"rate"}
                            label="音频语速"
                            placeholder="请输入音频语速"
                            tooltip={"格式为(+-)0%,如windows不能正常生成改为(+-)0%%"}
                        />
                        <ProFormText
                            width="md"
                            name={"volume"}
                            label="音量"
                            tooltip={"格式为(+-)0%,如windows不能正常生成改为(+-)0%%"}
                            placeholder="请输入音量"
                        />
                        <ProFormText
                            width="md"
                            name={"pitch"}
                            label="分贝"
                            tooltip={"格式为(+-)0Hz"}
                            placeholder="请输入分贝"
                        />
                        <ProFormSelect
                            width="md"
                            name={"voice"}
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
                                            {option.data.mp3 && <NotificationOutlined
                                                onClick={(e) => {
                                                    e.stopPropagation();
                                                    if (mp3Ref.current) {
                                                        mp3Ref.current.src = option.data.mp3;
                                                        mp3Ref.current.play();
                                                    }
                                                }}
                                            />}
                                        </Space>
                                    );
                                }
                            }}
                        />
                    </ProForm.Group>
                </ModalForm>;
            }
        },
        {
            title: "操作",
            align: "center",
            valueType: 'option',
            fixed: "right",
            render(text, record, _, action, s) {
                return (
                    <Spin spinning={false}>
                        <Button
                            type={"link"} onClick={() => {
                                action?.startEditable?.(record.id);
                            }}>
                            编辑
                        </Button>
                        <Button type={"link"} onClick={() => translatePrompt({ id: record?.id })}>
                            翻译
                        </Button>
                        <Button
                            type={"link"} onClick={async () => {
                                setRenderLoading(true);
                                await createAudioSrt({
                                    id: projectDetail?.id,
                                    infoId: record.id
                                });
                                setRenderLoading(false);
                            }}>
                            生成音频
                        </Button>
                        <Button
                            type={"link"} onClick={async () => {
                                await text2imageBatchHandler(record);
                            }}>
                            生成图片
                        </Button>
                        <Button
                            type={"link"} onClick={async () => {
                                await createInfoVideo(record.id);
                            }}>
                            生成视频
                        </Button>
                    </Spin>
                );
            }
        },
    ];

    useEffect(() => {
        if (state.id) {
            getProjectDetailConfig(state.id);
        }
    }, [ state ]);

    useEffect(() => {
        getDocumentHeight();
        window.addEventListener("resize", getDocumentHeight);
        return () => {
            window.removeEventListener("resize", getDocumentHeight);
        };
    }, []);

    if (!projectDetail) return null;

    return (
        <ProjectDetailPageWrap>
            <audio ref={mp3Ref} style={{ display: "none" }}/>
            <Spin spinning={tableLoading}>
                <EditableProTable
                    rowKey={"id"}
                    editable={{
                        onDelete: async (_, data) => {
                            await onDetailProjectInfo(data.id);
                        },
                        onSave: async (_, data) => {
                            await infoApi.updateProjectDetailInfo({
                                ...data
                            });
                            openMessageBox({ type: "success", message: "保存成功" });
                            await getProjectDetailConfig(state.id);
                        },
                    }}
                    recordCreatorProps={false}
                    value={projectDetail?.infoList ?? []}
                    columns={columns}
                    virtual={true}
                    scroll={{ y: tableHeight, x: 400 }}
                    headerTitle={<div>
                        <Button
                            style={{ marginRight: "8px" }} onClick={() => {
                                navigate("/");
                            }}>返回</Button>
                        {projectDetail?.fileName}
                    </div>}
                    pagination={false}
                    search={false}
                    toolBarRender={() => [
                        <Tooltip key={"loading"} title={"生成状态中，建议不要点击修改操作。没有做这部分的处理。避免没必要的错误"}>
                            {
                                renderLoading && <div style={{ display: "flex", alignItems: "center", justifyContent: "center" }}>
                                    <Spin indicator={<LoadingOutlined spin/>} size={"small"}/>
                                    <ExclamationOutlined/>
                                </div>
                            }
                        </Tooltip>,
                        <ModalForm
                            key={"upload"}
                            title="配置参数"
                            trigger={
                                <Button>上传文本</Button>
                            }
                            form={uploadForm}
                            onFinish={onUploadOkHandler}
                        >
                            <ProFormSelect
                                name={"saveType"}
                                label="文件添加类型"
                                placeholder="请选择添加类型"
                                options={[
                                    {
                                        value: "create",
                                        label: "新增"
                                    },
                                    {
                                        value: "update",
                                        label: "追加"
                                    }
                                ]}
                                rules={[ { required: true, message: '请选择添加类型' } ]}
                            />
                            <ProFormSelect
                                name={"whetherParticiple"}
                                label="是否进行分词"
                                placeholder="请选择是否进行分词"
                                tooltip={"是则按照设置的长度进行分割，反之只是去除多余非正常字符"}
                                options={[
                                    {
                                        value: "no",
                                        label: "否"
                                    },
                                    {
                                        value: "yes",
                                        label: "是"
                                    }
                                ]}
                                rules={[ { required: true, message: '请选择是否进行分词' } ]}
                            />
                            <ProFormUploadButton
                                name={"files"}
                                label="文件"
                                accept={".txt"}
                                fieldProps={{
                                    maxCount: 1,
                                    beforeUpload: () => false
                                }}
                                placeholder="文件"
                                rules={[ { required: true, message: '请上传文本' } ]}
                            />
                        </ModalForm>,
                        <Button key={"keywords"} disabled={!projectDetail?.infoList?.length} onClick={keywordsExtract}>
                            关键字提取
                        </Button>,
                        <Tooltip key={"role"} title={"当前版本没有开放分词词典,所以并不准确"}>
                            <Button key={"extract"} disabled={!projectDetail?.infoList?.length} onClick={extractTheRole}>
                                角色提取
                            </Button>
                        </Tooltip>,
                        <Button key={"translate"} disabled={!projectDetail?.infoList?.length} onClick={() => translatePrompt({ projectDetailId: projectDetail?.id })}>
                            翻译
                        </Button>,
                        <Button key={"audio"} disabled={!projectDetail?.infoList?.length} onClick={createAudioAndSrtHandler}>
                            生成音频和字幕
                        </Button>,
                        <Button key={"stable-diffusion"} disabled={!projectDetail?.infoList?.length} onClick={() => text2imageBatchHandler()}>
                            生成图片
                        </Button>,
                        <Tooltip key={"video"} title={"公开版本不添加视频合并之间的过渡动画"}>
                            <Button disabled={!projectDetail?.infoList?.length} onClick={() => createInfoVideo()}>
                                生成视频
                            </Button>
                        </Tooltip>,
                        <ModalForm
                            key={"settings"}
                            title="配置参数"
                            trigger={
                                <Tooltip title={"配置stable-diffusion请求参数"}>
                                    <SettingOutlined onClick={setStableDiffusionJson}/>
                                </Tooltip>
                            }
                            form={form}
                            onFinish={onSettingsOkHandler}
                        >
                            <ReactSmoothScrollbar style={{ maxHeight: "calc(100vh - 320px)" }}>
                                <ProForm.Group title={"项目基础配置"}>
                                    <ProFormSelect
                                        width="md"
                                        name={"breakAudio"}
                                        label="是否跳过存在的音频"
                                        placeholder="请选择是否跳过存在的音频"
                                        rules={[ { required: true, message: '请选择是否跳过存在的音频' } ]}
                                        options={[
                                            {
                                                label: "是",
                                                // @ts-ignore
                                                value: true,
                                            },
                                            {
                                                label: "否",
                                                // @ts-ignore
                                                value: false,
                                            }
                                        ]}
                                    />
                                    <ProFormSelect
                                        width="md"
                                        name={"breakVideo"}
                                        label="是否跳过存在的视频"
                                        placeholder="请选择是否跳过存在的视频"
                                        rules={[ { required: true, message: '请选择是否跳过存在的视频' } ]}
                                        options={[
                                            {
                                                label: "是",
                                                // @ts-ignore
                                                value: true,
                                            },
                                            {
                                                label: "否",
                                                // @ts-ignore
                                                value: false,
                                            }
                                        ]}
                                    />
                                    <ProFormSelect
                                        width="md"
                                        name={"concatAudio"}
                                        label="合并音频"
                                        placeholder="请选择是否合并音频"
                                        rules={[ { required: true, message: '请选择是否合并音频' } ]}
                                        options={[
                                            {
                                                label: "是",
                                                // @ts-ignore
                                                value: true,
                                            },
                                            {
                                                label: "否",
                                                // @ts-ignore
                                                value: false,
                                            }
                                        ]}
                                    />
                                    <ProFormSelect
                                        width="md"
                                        name={"concatVideo"}
                                        label="合并视频"
                                        placeholder="请选择是否合并视频"
                                        rules={[ { required: true, message: '请选择是否合并视频' } ]}
                                        options={[
                                            {
                                                label: "是",
                                                // @ts-ignore
                                                value: true,
                                            },
                                            {
                                                label: "否",
                                                // @ts-ignore
                                                value: false,
                                            }
                                        ]}
                                    />
                                    <ProFormSelect
                                        width="md"
                                        name={"openSubtitles"}
                                        label="开启字幕"
                                        placeholder="请选择是否开启字幕"
                                        rules={[ { required: true, message: '请选择是否开启字幕' } ]}
                                        options={[
                                            {
                                                label: "是",
                                                // @ts-ignore
                                                value: true,
                                            },
                                            {
                                                label: "否",
                                                // @ts-ignore
                                                value: false,
                                            }
                                        ]}
                                    />
                                    <ProFormSelect
                                        width="md"
                                        name={"language"}
                                        label="语言设置"
                                        placeholder="请选择语言"
                                        options={[
                                            {
                                                value: "zh",
                                                label: "中文"
                                            },
                                            {
                                                value: "en",
                                                label: "英语"
                                            }
                                        ]}
                                        rules={[ { required: true, message: '请选择语言' } ]}
                                    />
                                    <ProFormSelect
                                        width="md"
                                        name={"openContext"}
                                        label="开启上下文"
                                        placeholder="请选择是否开启上下文"
                                        options={[
                                            {
                                                // @ts-ignore
                                                value: true,
                                                label: "是"
                                            },
                                            {
                                                // @ts-ignore
                                                value: false,
                                                label: "否"
                                            }
                                        ]}
                                        rules={[ { required: true, message: '请选择是否开启上下文' } ]}
                                    />
                                    <Form.Item rules={[ { required: false, message: '自定义prompt路径' } ]} label={"自定义prompt路径"} name={"promptUrl"}>
                                        <Input className={"sava-path"} addonAfter={<div className={"file-svg"} onClick={onPromptUrlPathSelect}><FileSvg/></div>}/>
                                    </Form.Item>
                                </ProForm.Group>
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
                                        tooltip={"格式为(+-)0%,如windows不能正常生成改为(+-)0%%"}
                                    />
                                    <ProFormText
                                        width="md"
                                        name={[ "audioConfig", "volume" ]}
                                        label="音量"
                                        tooltip={"格式为(+-)0%,如windows不能正常生成改为(+-)0%%"}
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
                                                        {option.data.mp3 && <NotificationOutlined
                                                            onClick={(e) => {
                                                                e.stopPropagation();
                                                                if (mp3Ref.current) {
                                                                    mp3Ref.current.src = option.data.mp3;
                                                                    mp3Ref.current.play();
                                                                }
                                                            }}
                                                        />}
                                                    </Space>
                                                );
                                            }
                                        }}
                                    />
                                </ProForm.Group>
                                <ProForm.Group title={"视频配置"}>
                                    <ProFormDigit
                                        width="md"
                                        name={[ "videoConfig", "fontSize" ]}
                                        label="字幕大小"
                                        min={1}
                                        placeholder="请填写字幕大小"
                                        rules={[ { required: true, message: '请填写字幕大小' } ]}
                                    />
                                    <ProFormSelect
                                        width="md"
                                        name={[ "videoConfig", "fontColor" ]}
                                        tooltip={"公开版只支持白色,想自己改的话修改yaml文件"}
                                        label="字幕颜色"
                                        placeholder="请选择字幕颜色"
                                        rules={[ { required: true, message: '请选择字幕颜色' } ]}
                                        options={[
                                            {
                                                label: "白色",
                                                value: "FFFFFF"
                                            }
                                        ]}
                                    />
                                    {/*<Form.Item label={"字幕颜色"} name={[ "videoConfig", "fontColor" ]}>*/}
                                    {/*    <ColorPicker*/}
                                    {/*        format={"hex"} onChange={(color) => {*/}
                                    {/*            setFontColor(color.toHex());*/}
                                    {/*        }}*/}
                                    {/*    />*/}
                                    {/*</Form.Item>*/}
                                    <ProFormDigit
                                        width="md"
                                        name={[ "videoConfig", "animationSpeed" ]}
                                        tooltip={"最佳1.2"}
                                        label="动画浮动比例"
                                        placeholder="请填写动画浮动比例"
                                        rules={[ { required: true, message: '请填写动画浮动比例' } ]}
                                    />
                                    <ProFormSelect
                                        width="md"
                                        name={[ "videoConfig", "position" ]}
                                        label="字幕位置"
                                        placeholder="请选择字幕位置"
                                        rules={[ { required: true, message: '请选择字幕位置' } ]}
                                        options={[
                                            {
                                                label: "上",
                                                value: 6
                                            },
                                            {
                                                label: "中",
                                                value: 10
                                            },
                                            {
                                                label: "下",
                                                value: 2
                                            }
                                        ]}
                                    />
                                    <ProFormSelect
                                        width="md"
                                        name={[ "videoConfig", "animationName" ]}
                                        tooltip={"精力有限,公开版只做全局随机效果"}
                                        label="动画效果"
                                        placeholder="请选择动画效果"
                                        rules={[ { required: true, message: '请选择动画效果' } ]}
                                        options={[
                                            {
                                                label: "随机动画",
                                                value: "random"
                                            }
                                        ]}
                                    />
                                    <Form.Item noStyle={true} shouldUpdate={true}>
                                        {
                                            (({ getFieldValue }) => {
                                                return <Form.Item tooltip={"不支持"} label={"字幕字体"} name={[ "videoConfig", "fontFile" ]} shouldUpdate={true}>
                                                    <Button onClick={onFilePathSelect} disabled={true}>
                                                        字幕选择,当前字体路径:{getFieldValue([ "videoConfig", "fontFile" ])}
                                                    </Button>
                                                </Form.Item>;
                                            })
                                        }
                                    </Form.Item>
                                </ProForm.Group>
                                <Divider orientation="left">stable-diffusion配置</Divider>
                                <VanillaUploadJson
                                    content={content}
                                    onChange={handleChange}
                                    onImportHandler={handleUpload}
                                    onExportHandler={handleDownload}
                                />
                            </ReactSmoothScrollbar>
                        </ModalForm>
                    ]}
                />
            </Spin>
        </ProjectDetailPageWrap>
    );
};


export default ProjectDetailPage;
