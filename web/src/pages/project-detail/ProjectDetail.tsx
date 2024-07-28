import { useLocation } from "react-router";
import styled from "styled-components";
import { Button, Dropdown, Form, MenuProps, UploadFile } from "antd";
import { EllipsisOutlined } from "@ant-design/icons";
import { useEffect, useState } from "react";
import { extractTheCharacterProjectDetailParticipleList, getProjectDetail, TranslateProjectDetailParticipleList, updateProjectDetail } from "../../api/projectApi.ts";
import { ProjectDetailParticipleList, ProjectDetailResponse } from "../../api/response/projectResponse.ts";
import { ModalForm, ProColumns, ProForm, ProFormDigit, ProFormUploadButton, ProTable } from "@ant-design/pro-components";

const ProjectDetailPageWrap = styled.div`
`;

const ProjectDetailPage = () => {
    const location = useLocation();
    const [ form ] = Form.useForm();
    const [ projectDetail, setProjectDetail ] = useState<ProjectDetailResponse>();
    const state = location.state;

    const getProjectDetailConfig = async (id: number) => {
        const detail = await getProjectDetail({
            projectId: id
        });
        setProjectDetail(detail);
    };

    const onUploadOkHandler = async (values: Partial<ProjectDetailResponse & {
        file: UploadFile[];
    }>) => {
        const { file, ...args } = values;
        if (projectDetail) {
            await updateProjectDetail({
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

    const translatePrompt = async () => {
        if (!projectDetail) return;
        await TranslateProjectDetailParticipleList({
            id: projectDetail?.id
        });
        await getProjectDetailConfig(state.id);
    };

    const items: MenuProps['items'] = [
        {
            key: 'delete',
            danger: true,
            label: (
                <span>
                    删除
                </span>
            ),
        },
    ];

    const columns: ProColumns<ProjectDetailParticipleList>[] = [
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
            dataIndex: "text",
            title: "文本",
            ellipsis: true,
        },
        {
            dataIndex: "prompt",
            title: "正向提示词",
        },
        {
            dataIndex: "character",
            title: "人物",
        },
        {
            dataIndex: "action",
            title: "操作",
            align: "center",
            width: 70,
            render() {
                return (
                    <Dropdown menu={{ items }} trigger={[ "click" ]}>
                        <EllipsisOutlined/>
                    </Dropdown>
                );
            }
        },
    ];

    useEffect(() => {
        if (state.id) {
            getProjectDetailConfig(state.id);
        }
    }, [ state ]);

    return (
        <ProjectDetailPageWrap>
            <ProTable
                rowKey={"id"}
                dataSource={projectDetail?.participleList ?? []}
                columns={columns}
                virtual={true}
                scroll={{ y: 650 }}
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
                    <Button onClick={translatePrompt}>
                        prompt转换
                    </Button>
                ]}
            />
        </ProjectDetailPageWrap>
    );
};

export default ProjectDetailPage;
