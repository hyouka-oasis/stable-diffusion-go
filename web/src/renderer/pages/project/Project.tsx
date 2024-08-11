import styled from "styled-components";
import { Button, Col, Input, Modal, Row, Table, TableColumnsType } from "antd";
import { useContext, useEffect, useState } from "react";
import useMessage from "antd/es/message/useMessage";
import { useNavigate } from "react-router-dom";
import { ProjectResponse } from "renderer/api/response/projectResponse";
import { projectApi, projectDetailApi } from "renderer/api";
import { EditableProTable, ProColumns } from "@ant-design/pro-components";
import { AppGlobalContext } from "renderer/shared/context/appGlobalContext";

const ProjectWrap = styled.div`
    .action {
        display: flex;
        align-items: center;
        justify-content: end;
        padding-bottom: 24px;
    }

    .ant-pro-card-body {
        padding: 0;
    }
`;

const TableActionWrap = styled.div`

`;

const ProjectPage = () => {
    const [ open, setOpen ] = useState<boolean>(false);
    const [ name, setName ] = useState<string>();
    const [ messageApi, messageContext ] = useMessage();
    const [ projectList, setProjectList ] = useState<ProjectResponse[]>([]);
    const navigate = useNavigate();
    const { openMessageBox } = useContext(AppGlobalContext);

    const onOkHandler = async () => {
        if (!name) {
            messageApi.error({
                content: "名称不能为空!"
            });
            return;
        }
        await projectApi.createProject({ name });
        messageApi.success({
            content: "创建成功"
        });
        setOpen(false);
        await getProjectListApi();
    };
    /**
     * 获取列表
     */
    const getProjectListApi = async () => {
        const data = await projectApi.getProjectList({
            page: 1,
            pageSize: -1,
        });
        setProjectList(data.list);
    };
    /**
     * 删除项目
     * @param id
     */
    const onDeleteProject = async (id: number) => {
        await projectApi.deleteProject({ id });
        messageApi.success({
            content: "删除成功"
        });
        await getProjectListApi();
    };
    /**
     * 删除详情
     * @param id
     */
    const onDeleteProjectDetail = async (id: number) => {
        await projectDetailApi.deleteProjectDetail({ id });
        messageApi.success({
            content: "删除成功"
        });
        await getProjectListApi();
    };


    const columns: ProColumns<ProjectResponse>[] = [
        {
            title: "项目名称",
            dataIndex: "name",
            key: "name",
            ellipsis: true,
        },
        {
            title: "操作",
            valueType: 'option',
            align: "center",
            width: 220,
            render(text, record, _, action) {
                return (
                    <TableActionWrap>
                        <Button
                            type={"link"} onClick={() => {
                                action?.startEditable?.(record.id);
                            }}>
                            编辑
                        </Button>
                        <Button
                            type={"link"} onClick={async () => {
                                const projectDetail = await projectDetailApi.createProjectDetail({
                                    projectId: record.id
                                });
                                navigate("/detail", {
                                    state: {
                                        id: projectDetail.id
                                    }
                                });
                            }}>
                            新增
                        </Button>
                        <Button type={"link"} danger onClick={() => onDeleteProject(record.id)}>
                            删除
                        </Button>
                    </TableActionWrap>
                );
            }
        }
    ];

    const expandedRowRender = (recode: any) => {
        const columns: TableColumnsType<ProjectResponse> = [
            { title: '书本名称', dataIndex: 'fileName', key: 'fileName' },
            {
                title: "操作",
                dataIndex: "description",
                width: 180,
                render(_, record) {
                    return (
                        <TableActionWrap>
                            <Button
                                type={"link"} onClick={() => {
                                    navigate("/detail", {
                                        state: {
                                            id: record.id
                                        }
                                    });
                                }}>
                                编辑
                            </Button>
                            <Button type={"link"} danger onClick={() => onDeleteProjectDetail(record.id)}>
                                删除
                            </Button>
                        </TableActionWrap>
                    );
                }
            }
        ];
        return <Table columns={columns} dataSource={recode.list} pagination={false}/>;
    };

    useEffect(() => {
        getProjectListApi();
    }, []);

    return (
        <ProjectWrap>
            {messageContext}
            <div className={'action'}>
                <Button onClick={() => setOpen(true)}>
                    新增项目
                </Button>
            </div>
            <EditableProTable
                recordCreatorProps={false}
                columns={columns}
                rowKey={"id"}
                value={projectList}
                editable={{
                    onDelete: async (_, record) => {
                        await onDeleteProject(record.id);
                    },
                    onSave: async (_, data) => {
                        await projectApi.updateProject({
                            ...data
                        });
                        openMessageBox({ type: "success", message: "保存成功" });
                        await getProjectListApi();
                    },
                }}
                expandable={{
                    expandedRowRender: expandedRowRender,
                }}
            />
            <Modal open={open} title={"新增项目"} onOk={onOkHandler} onCancel={() => setOpen(false)}>
                <Row justify="center" align="middle">
                    <Col span={6}>
                        项目名称
                    </Col>
                    <Col span={18}>
                        <Input value={name} onChange={(e) => setName(e.target.value)}/>
                    </Col>
                </Row>
            </Modal>
        </ProjectWrap>
    );
};

export default ProjectPage;
