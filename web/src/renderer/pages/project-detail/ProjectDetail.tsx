import { useLocation, useNavigate } from "react-router";
import styled from "styled-components";
import { Button, Divider, Form, Modal, Radio, Space, Spin, Tooltip, UploadProps } from "antd";
import { useContext, useEffect, useRef, useState } from "react";
import { EditableProTable, ModalForm, ProColumns, ProForm, ProFormDigit, ProFormSelect, ProFormText, ProFormUploadButton } from "@ant-design/pro-components";
import { Content, OnChange, TextContent } from "vanilla-jsoneditor";
import { CloseOutlined, NotificationOutlined, SettingOutlined } from "@ant-design/icons";
import { Info, ProjectDetailResponse } from "renderer/api/response/projectResponse";
import { stableDiffusionText2Image } from "renderer/api/stableDiffusionApi";
import { blobToFile, dataURLtoBlob } from "renderer/utils/utils";
import { uploadFile } from "renderer/api/fileApi";
import { createAudioSrt } from "renderer/api/audioSrtApi";
import { baseURL } from "renderer/request/request";
import { audioList } from "renderer/utils/audio-list";
import VanillaUploadJson from "renderer/components/json-edit/VanillaUploadJson";
import { projectApi, stableDiffusionApi } from "renderer/api";
import { AppGlobalContext } from "renderer/shared/context/appGlobalContext";
import { FileResponse } from "renderer/api/response/fileResponse";
import { RcFile } from "antd/lib/upload";
import { navBarHeight } from "renderer/shared";
import { ReactSmoothScrollbar } from "renderer/components/smooth-scroll/SmoothScroll";

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

const AudioActionWrap = styled.div``;
const AudioActionItemWrap = styled.div`
    display: flex;
    align-items: center;
    justify-content: space-between;
`;

const percentageRegex = /^[-+]\d+%$/;
const pitchRegex = /^[+]\d+Hz$/;

const ProjectDetailPage = () => {
    const location = useLocation();
    const navigate = useNavigate();
    const [ modalApi, setModalContext ] = Modal.useModal();
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
    const { openMessageBox } = useContext(AppGlobalContext);
    const state = location.state;
    const mp3Ref = useRef<HTMLAudioElement | null>(null);

    const changeInfoLoading = (idConfig: {
        id?: number;
        batch?: boolean;
    }, loading: boolean) => {
        setProjectDetail((detail) => {
            if (!detail) return detail;
            let infoList = detail?.infoList ?? [];
            infoList = infoList.map(info => {
                if (idConfig?.id) {
                    if (info.id === idConfig?.id) {
                        info.loading = loading;
                    }
                }
                if (idConfig?.batch) {
                    info.loading = loading;
                }
                return info;
            });
            return {
                ...detail,
                infoList: infoList
            };
        });
    };

    /**
     * 获取项目
     * @param id
     */
    const getProjectDetailConfig = async (id: number) => {
        setTableLoading(true);
        const detail = await projectApi.getProjectDetail({
            projectId: id
        });
        setProjectDetail(detail);
        form.setFieldsValue({
            ...detail,
        });
        console.log(123);
        setTableLoading(false);
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
                    return;
                }
                if (!percentageRegex.test(values?.audioConfig?.volume ?? "0")) {
                    openMessageBox({ type: "error", message: "请输入正确的音量值!" });
                    return;
                }
                if (!pitchRegex.test(values?.audioConfig?.pitch ?? "0")) {
                    openMessageBox({ type: "error", message: "请输入正确的分贝值!" });
                    return;
                }
                projectApi.updateProjectDetail({
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
            projectApi.uploadProjectDetail({
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
        await projectApi.extractTheCharacterProjectDetailParticipleList({
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
        if (data.id) {
            changeInfoLoading({
                id: data.id
            }, true);
        }
        await projectApi.translateProjectDetailParticipleList(data);
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
            let selectedId: null | number = null;
            const stableDiffusionImages: Info["stableDiffusionImages"] = [];
            const images = await stableDiffusionText2Image({
                ids: [ info.id ],
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
                        projectDetailId: info.projectDetailId,
                        InfoId: info.id,
                        name: upload.name,
                        key: upload.key,
                        url: upload.url,
                        tag: upload.tag,
                        fileId: upload.id,
                    });
                }
                await projectApi.updateProjectDetailInfo(Object.assign({
                    id: info.id,
                    stableDiffusionImages,
                }, (selectedId && !info.stableDiffusionImageId) ? {
                    stableDiffusionImageId: selectedId
                } : {}));
            }
            await getProjectDetailConfig(state.id);
        };
        if (infoDetail) {
            changeInfoLoading({ id: infoDetail.id }, true);
            await handler(infoDetail);
        } else {
            for (const info of (projectDetail?.infoList ?? [])) {
                changeInfoLoading({ batch: true }, true);
                await handler(info);
            }
        }
    };
    /**
     * 生成视频
     */
    const createInfoVideo = async (id?: number) => {
        const ids = id ? [ id ] : [];
        changeInfoLoading({ batch: true }, true);
        if (projectDetail) {
            await projectApi.createInfoVideo({
                ids,
                projectDetailId: projectDetail?.id,
            });
            openMessageBox({ type: "success", message: "生成成功" });
            changeInfoLoading({ batch: true }, false);
        }
    };
    /**
     * 改变选中的图片
     * @param selectedId
     * @param infoId
     */
    const onStableDiffusionImagesOnChange = async (selectedId: number, infoId: number) => {
        await projectApi.updateProjectDetailInfo(Object.assign({
            id: infoId,
            stableDiffusionImageId: selectedId
        }));
        await getProjectDetailConfig(state.id);
    };
    /**
     * 生成音频和字幕
     */
    const createAudioAndSrtHandler = async () => {
        changeInfoLoading({ batch: true }, true);
        if (projectDetail) {
            await createAudioSrt({
                id: projectDetail?.id
            });
            openMessageBox({ type: "success", message: "生成成功" });
            changeInfoLoading({ batch: true }, false);
        }
    };

    const onDetailProjectInfo = async (id: number) => {
        await projectApi.deleteInfo({ id });
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
        await projectApi.updateProjectDetailInfo({
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
                                            <Radio key={file.id} value={file.id}>
                                                <span
                                                    className={"action-delete"} onClick={async (e) => {
                                                        e.preventDefault();
                                                        e.stopPropagation();
                                                        await stableDiffusionApi.stableDiffusionDeleteImage({ ids: [ file.id ] });
                                                        await getProjectDetailConfig(state.id);
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
            dataIndex: "audioConfig",
            title: "音频设置",
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
                            tooltip={"格式为(+-)0%"}
                        />
                        <ProFormText
                            width="md"
                            name={"volume"}
                            label="音量"
                            tooltip={"格式为(+-)0%"}
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
                    <Spin spinning={record.loading}>
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
                                changeInfoLoading({
                                    id: record.id
                                }, true);
                                await createAudioSrt({
                                    id: projectDetail?.id,
                                    infoId: record.id
                                });
                                changeInfoLoading({
                                    id: record.id
                                }, false);
                            }}>
                            生成音频
                        </Button>
                        <Button
                            type={"link"} onClick={async () => {
                                await text2imageBatchHandler(record);
                            }}>
                            生成图片
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
            {setModalContext}
            <audio ref={mp3Ref} style={{ display: "none" }}/>
            <Spin spinning={tableLoading}>
                <EditableProTable
                    rowKey={"id"}
                    editable={{
                        onDelete: async (_, data) => {
                            await onDetailProjectInfo(data.id);
                        },
                        onSave: async (_, data) => {
                            await projectApi.updateProjectDetailInfo({
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
                                        label: "覆盖"
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
                        <Button key={"extract"} disabled={!projectDetail?.infoList?.length} onClick={extractTheRole}>
                            角色提取
                        </Button>,
                        <Button key={"translate"} disabled={!projectDetail?.infoList?.length} onClick={() => translatePrompt({ projectDetailId: projectDetail?.id })}>
                            翻译
                        </Button>,
                        <Button key={"audio"} disabled={!projectDetail?.infoList?.length} onClick={createAudioAndSrtHandler}>
                            生成音频和字幕
                        </Button>,
                        <Button key={"stable-diffusion"} disabled={!projectDetail?.infoList?.length} onClick={() => text2imageBatchHandler()}>
                            生成图片
                        </Button>,
                        <Button key={"video"} disabled={!projectDetail?.infoList?.length} onClick={() => createInfoVideo()}>
                            生成视频
                        </Button>,
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
                                    <ProFormSelect
                                        width="md"
                                        name={"batchAudio"}
                                        label="全量替换音频"
                                        placeholder="请选择是否全量替换音频"
                                        rules={[ { required: true, message: '请选择是否全量替换音频' } ]}
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
