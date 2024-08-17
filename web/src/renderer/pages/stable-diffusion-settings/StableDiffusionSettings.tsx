import styled from "styled-components";
import { ModalForm, ProColumns, ProForm, ProFormText, ProTable } from "@ant-design/pro-components";
import { useContext, useEffect, useState } from "react";
import { Button, Form } from "antd";
import { StableDiffusionNegativePromptResponse } from "renderer/api/response/stableDiffusionResponse";
import { stableDiffusionSettingsApi } from "renderer/api";
import { ReactSmoothScrollbar } from "renderer/components/smooth-scroll/SmoothScroll";
import { ProjectDetailResponse } from "renderer/api/response/projectResponse";
import { AppGlobalContext } from "renderer/shared/context/appGlobalContext";
import StableDiffusionForm from "renderer/components/stable-diffusion-form/StableDiffusionForm";

const StableDiffusionSettingsWrap = styled.div`
`;

const StableDiffusionSettings = () => {
    const [ settingsList, setSettingsList ] = useState<StableDiffusionNegativePromptResponse[]>([]);
    const [ form ] = Form.useForm();
    const [ modalOpen, setModalOpen ] = useState<boolean>(false);
    const [ settingsDetail, setSettingsDetail ] = useState<StableDiffusionNegativePromptResponse>();
    const { openMessageBox } = useContext(AppGlobalContext);


    const getSettingsList = async () => {
        const list = await stableDiffusionSettingsApi.getStableDiffusionSettingsList({
            page: 1,
            pageSize: -1,
        });
        setSettingsList(list.list);
    };

    const onSettingsOkHandler = (values: Partial<ProjectDetailResponse>): Promise<boolean> => {
        return new Promise(resolve => {
            (settingsDetail ? stableDiffusionSettingsApi.updateStableDiffusionSettings : stableDiffusionSettingsApi.createStableDiffusionSettings)(Object.assign({
                ...(values.stableDiffusionConfig as any),
            }, settingsDetail ? {
                id: settingsDetail?.id
            } : {})).then(res => {
                openMessageBox({ type: "success", message: "更新配置成功" });
                getSettingsList();
                resolve(true);
            });
        });
    };

    const columns: ProColumns<StableDiffusionNegativePromptResponse>[] = [
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
            title: "名称",
        },
        {
            title: "操作",
            align: "center",
            valueType: 'index',
            fixed: "right",
            render(text, record, _, action, s) {
                return (
                    <>
                        <Button
                            type={"link"}
                            onClick={() => {
                                setSettingsDetail(record);
                                form.setFieldsValue({
                                    stableDiffusionConfig: {
                                        ...record
                                    }
                                });
                                setModalOpen(true);
                            }}
                        >
                            编辑
                        </Button>
                        <Button
                            type={"link"}
                            danger
                            onClick={async () => {
                                await stableDiffusionSettingsApi.deleteStableDiffusionSettings({
                                    ids: [ record.id ]
                                });
                                await getSettingsList();
                            }}
                        >
                            删除
                        </Button>
                    </>
                );
            }
        },
    ];

    useEffect(() => {
        getSettingsList();
    }, []);

    useEffect(() => {
        if (!modalOpen) {
            setSettingsDetail(undefined);
        }
    }, [ modalOpen ]);

    return (
        <StableDiffusionSettingsWrap>
            <ProTable
                rowKey={"id"}
                dataSource={settingsList ?? []}
                columns={columns}
                pagination={false}
                search={false}
                toolBarRender={() => [
                    <ModalForm
                        key={"sd-settings"}
                        open={modalOpen}
                        title="创建/修改配置"
                        trigger={
                            <Button type="primary" onClick={() => setModalOpen(true)}>
                                创建/修改配置
                            </Button>
                        }
                        onOpenChange={setModalOpen}
                        form={form}
                        onFinish={onSettingsOkHandler}
                    >
                        <ReactSmoothScrollbar style={{ maxHeight: "calc(100vh - 320px)" }}>
                            <ProForm.Group title={"基础配置"}>
                                <ProFormText
                                    width="xl"
                                    name={[ "stableDiffusionConfig", "name" ]}
                                    label="sd配置名称"
                                    placeholder="请填写sd配置名称"
                                    rules={[ { required: true, message: '请填写sd配置名称' } ]}
                                />
                            </ProForm.Group>
                            <span/>
                            <StableDiffusionForm/>
                        </ReactSmoothScrollbar>
                    </ModalForm>,
                ]}
            />
        </StableDiffusionSettingsWrap>
    );
};

export default StableDiffusionSettings;
