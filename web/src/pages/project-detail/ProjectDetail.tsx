import { useLocation } from "react-router";
import styled from "styled-components";
import { Button, Form, Modal, Upload } from "antd";
import { useState } from "react";

const ProjectDetailPageWrap = styled.div`
    .action {
        display: flex;
    }
`;

const ProjectDetailPage = () => {
    const location = useLocation();
    const [ uploadOpen, setUploadOpen ] = useState<boolean>(false);
    const [ form ] = Form.useForm();

    const onUploadOkHandler = async () => {
        const uploadValues = await form.validateFields();
        console.log(uploadValues);
    };

    return (
        <ProjectDetailPageWrap>
            <Modal open={uploadOpen} onOk={onUploadOkHandler}>
                <Form layout={"vertical"} form={form}>
                    <Form.Item name={"upload"} label={"文件"} valuePropName={"filelist"}>
                        <Upload beforeUpload={() => {
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
