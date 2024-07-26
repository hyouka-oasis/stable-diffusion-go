import { useLocation } from "react-router";
import styled from "styled-components";
import { Button, Form, InputNumber, Modal, Upload } from "antd";
import { useEffect, useState } from "react";
import { getProjectDetail, updateProjectDetail } from "../../api/projectApi.ts";

const ProjectDetailPageWrap = styled.div`
    .action {
        display: flex;
    }
`;

const ProjectDetailPage = () => {
    const location = useLocation();
    const [ uploadOpen, setUploadOpen ] = useState<boolean>(false);
    const [ form ] = Form.useForm();
    const [ projectDetail, setProjectDetail ] = useState<any>();
    const state = location.state;

    const getProjectDetailConfig = async (id) => {
        const detail = await getProjectDetail({
            projectId: id
        });
        setProjectDetail(detail);
    };

    const onUploadOkHandler = async () => {
        const uploadValues = await form.validateFields();
        const { file, ...args } = uploadValues;
        updateProjectDetail({
            id: projectDetail.id,
            file: file.file,
            ...args
        });
    };

    useEffect(() => {
        if (state.id) {
            getProjectDetailConfig(state.id);
        }
    }, [ state ]);

    return (
        <ProjectDetailPageWrap>
            <Modal open={uploadOpen} onOk={onUploadOkHandler} title={"导入文件"}>
                <Form layout={"vertical"} form={form}>
                    <Form.Item noStyle={true} label={"文本配置"}>
                        <Form.Item name={"minWords"} label={"最小文字数量"} rules={[ { required: true, message: '请输入最小文字数量' } ]}>
                            <InputNumber min={10}/>
                        </Form.Item>
                        <Form.Item name={"maxWords"} label={"最大文字数量"} rules={[ { required: true, message: '请输入最大文字数量' } ]}>
                            <InputNumber min={10}/>
                        </Form.Item>
                    </Form.Item>
                    <Form.Item name={"file"} label={"文件"} valuePropName={"filelist"} rules={[ { required: true, message: '请上传文件' } ]}>
                        <Upload accept={".txt"} beforeUpload={() => {
                            return false;
                        }}>
                            <Button>
                                导入文本
                            </Button>
                        </Upload>
                    </Form.Item>
                </Form>
            </Modal>
            <div className={'action'}>
                <Button type="primary" onClick={() => setUploadOpen(true)}>
                    开始处理
                </Button>
            </div>
        </ProjectDetailPageWrap>
    );
};

export default ProjectDetailPage;
