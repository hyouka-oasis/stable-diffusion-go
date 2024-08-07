import styled from "styled-components";
import { ModalForm, ProColumns, ProForm, ProFormText, ProFormUploadButton, ProTable } from "@ant-design/pro-components";
import { useEffect, useState } from "react";
import { Button, Form, UploadFile } from "antd";
import { RcFile } from "antd/lib/upload";
import { LorasResponse } from "renderer/api/response/lorasResponse";
import { uploadFile } from "renderer/api/fileApi";
import { createStableDiffusionLoras, getStableDiffusionLorasList } from "renderer/api/stableDiffusionApi";
import { baseURL } from "renderer/request/request";

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
                file: values.file?.[0]?.originFileObj as RcFile,
                fileType: "lora"
            });
            fileId = file.id;
            await createStableDiffusionLoras({
                imageId: fileId,
                ...args,
            });
            await getLorasList();
            return true;
        } else {
            await createStableDiffusionLoras({
                ...args,
            });
            await getLorasList();
            return true;
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
                return <img
                    style={{
                        width: "100px",
                        height: "100px"
                    }} src={`${baseURL}/${value}`}
                />;
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
