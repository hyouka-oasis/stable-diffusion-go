import styled from "styled-components";
import { ModalForm, ProColumns, ProForm, ProFormText, ProFormUploadButton, ProTable } from "@ant-design/pro-components";
import { useEffect, useState } from "react";
import { LorasResponse } from "../../api/response/lorasResponse.ts";
import { Button, Form, UploadFile } from "antd";
import { createStableDiffusionLoras, getStableDiffusionLorasList } from "../../api/stableDiffusionApi.ts";
import { uploadFile } from "../../api/fileApi.ts";
import { RcFile } from "antd/lib/upload";
import { baseURL } from "../../utils/request.ts";

const LorasWrap = styled.div`
`;

const LorasPage = () => {
    const [ lorasList, setLorasList ] = useState<LorasResponse[]>([]);
    const [ form ] = Form.useForm();

    const onCreateOkHandler = async (values: Partial<LorasResponse & {
        file: UploadFile[];
    }>) => {
        const { file: rcFile, ...args } = values;
        let fileId;
        if (values.file) {
            const file = await uploadFile({
                file: values.file?.[0]?.originFileObj as RcFile
            });
            fileId = file.id;
            await createStableDiffusionLoras({
                imageId: fileId,
                ...args,
            });
        } else {
            await createStableDiffusionLoras({
                ...args,
            });
        }
    };


    const getLorasList = async () => {
        const list = await getStableDiffusionLorasList({
            page: 1,
            pageSize: -1,
        });
        setLorasList(list.list);
    };

    const columns: ProColumns<LorasResponse>[] = [
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
            title: "lora标签",
        },
        {
            dataIndex: "roles",
            title: "人物",
        },
        {
            dataIndex: "url",
            title: "图片",
            render(value) {
                return <img style={{
                    width: "100px",
                    height: "100px"
                }} src={`${baseURL}/${value}`}/>;
            }
        },
        {
            dataIndex: "action",
            title: "操作",
            align: "center",
            width: 200,
            render() {
                return (
                    <Button.Group>
                        <Button>
                            修改
                        </Button>
                        <Button danger>
                            删除
                        </Button>
                    </Button.Group>
                );
            }
        },
    ];

    useEffect(() => {
        getLorasList();
    }, []);
    return (
        <LorasWrap>
            <ProTable
                rowKey={"id"}
                dataSource={lorasList ?? []}
                columns={columns}
                search={false}
                toolBarRender={() => [
                    <ModalForm
                        title="创建/修改lora"
                        trigger={
                            <Button type="primary">
                                创建lora
                            </Button>
                        }
                        form={form}
                        onFinish={onCreateOkHandler}
                    >
                        <ProForm.Group>
                            <ProFormText
                                width="md"
                                name="name"
                                label="lora名称"
                                placeholder="请输入lora名称"
                                rules={[ { required: true, message: '请输入lora名称' } ]}
                            />

                            <ProFormText
                                width="md"
                                name="roles"
                                label="角色名称"
                                placeholder="请输入角色名称"
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
                                label="图片"
                                placeholder="图片"
                                accept={".png,.jpeg,.jpg"}
                            />
                        </ProForm.Group>
                    </ModalForm>,
                ]}
            />
        </LorasWrap>
    );
};

export default LorasPage;
