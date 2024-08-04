import styled from "styled-components";
import { Button, Col, Input, Modal, Row, Table, TableProps } from "antd";
import { useEffect, useState } from "react";
import useMessage from "antd/es/message/useMessage";
import { useNavigate } from "react-router-dom";
import { ProjectResponse } from "renderer/api/response/projectResponse";
import { createProject, deleteProject, getProjectList } from "renderer/api/projectApi";

const ProjectWrap = styled.div`
    .action {
        display: flex;
        align-items: center;
        justify-content: end;
        padding-bottom: 24px;
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

    const onOkHandler = async () => {
        if (!name) {
            messageApi.error({
                content: "名称不能为空!"
            });
            return;
        }
        await createProject({ name });
        messageApi.success({
            content: "创建成功"
        });
        setOpen(false);
        await getProjectListApi();
    };

    const getProjectListApi = async () => {
        const data = await getProjectList({
            page: 1,
            pageSize: -1,
        });
        setProjectList(data.list);
    };

    const onDeleteProject = async (id: number) => {
        await deleteProject({ id });
        messageApi.success({
            content: "删除成功"
        });
        await getProjectListApi();
    };

    const columns: TableProps<ProjectResponse>["columns"] = [
        {
            title: "项目名称",
            dataIndex: "name",
            key: "name",
            ellipsis: true,
        },
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
                        <Button type={"link"} danger onClick={() => onDeleteProject(record.id)}>
                            删除
                        </Button>
                    </TableActionWrap>
                );
            }
        }
    ];

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
            <Table columns={columns} rowKey={"id"} dataSource={projectList}/>
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
