import { EditableProTable, ModalForm, ProColumns, ProFormTextArea } from "@ant-design/pro-components";
import { stableDiffusionApi } from "renderer/api";
import { Button, Form, message } from "antd";
import { useContext, useEffect, useState } from "react";
import { AppGlobalContext } from "renderer/shared/context/appGlobalContext";
import { StableDiffusionNegativePromptResponse } from "renderer/api/response/stableDiffusionResponse";

const NegativePrompts = () => {
    const { openMessageBox } = useContext(AppGlobalContext);
    const [ negativePromptsList, setNegativePromptsList ] = useState<StableDiffusionNegativePromptResponse[]>([]);
    const [ form ] = Form.useForm();

    const getNegativePromptList = async () => {
        const data = await stableDiffusionApi.getNegativePromptList({
            page: 1,
            pageSize: -1,
        });
        setNegativePromptsList(data.list);
    };

    const onOkHandler = async (values: Partial<StableDiffusionNegativePromptResponse>) => {
        await stableDiffusionApi.createNegativePrompt({
            ...values,
        });
        message.success("新增成功");
        await getNegativePromptList();
        return true;
    };

    const onDeleteHandler = async (id: number) => {
        await stableDiffusionApi.deleteNegativePrompt({
            id
        });
        message.success("删除成功");
        await getNegativePromptList();
    };

    const columns: ProColumns<StableDiffusionNegativePromptResponse>[] = [
        {
            dataIndex: "name",
            title: "名称",
            valueType: "textarea",
        },
        {
            dataIndex: "text",
            title: "正向提示词",
            valueType: "textarea",
        },

        {
            title: "操作",
            align: "center",
            valueType: 'option',
            fixed: "right",
            render(text, record, _, action) {
                return (
                    <>
                        <Button
                            danger type={"link"} onClick={() => onDeleteHandler(record.id)}>
                            删除
                        </Button>
                        <Button
                            type={"link"} onClick={() => {
                                action?.startEditable?.(record.id);
                            }}>
                            编辑
                        </Button>
                    </>
                );
            }
        },
    ];

    useEffect(() => {
        getNegativePromptList();
    }, []);

    return (
        <EditableProTable
            rowKey={"id"}
            editable={{
                onSave: async (_, data) => {
                    await stableDiffusionApi.updateNegativePrompt({
                        ...data
                    });
                    openMessageBox({ type: "success", message: "保存成功" });
                    await getNegativePromptList();
                },
            }}
            recordCreatorProps={false}
            value={negativePromptsList}
            columns={columns}
            virtual={true}
            pagination={false}
            search={false}
            toolBarRender={() => [
                <ModalForm
                    key={"settings"}
                    title="新建反向提示词"
                    trigger={
                        <Button type="primary">
                            新建反向提示词
                        </Button>
                    }
                    form={form}
                    onFinish={onOkHandler}
                >
                    <ProFormTextArea
                        name={"name"}
                        label="别名"
                        placeholder="请输入别名"
                        rules={[ { required: true, message: '请输入别名' } ]}
                    />
                    <ProFormTextArea
                        name={"text"}
                        label="反向提示词"
                        placeholder="请输入反向提示词"
                        rules={[ { required: true, message: '请输入反向提示词' } ]}
                    />
                </ModalForm>
            ]}
        />
    );
};

export default NegativePrompts;
