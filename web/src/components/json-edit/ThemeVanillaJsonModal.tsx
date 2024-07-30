import { FC } from "react";
import { Button, Modal, ModalProps, Upload, UploadProps } from "antd";
import { ExportOutlined, ImportOutlined } from "@ant-design/icons";
import ThemeVanillaJson from "./ThemeVanillaJson";
import styled from "styled-components";
import { JSONEditorPropsOptional } from "vanilla-jsoneditor";

const ThemeEditorJsonActionWrap = styled.div`
    margin-top: 24px;
    margin-bottom: -44px;
`;

interface ThemeVanillaJsonModalProps extends Pick<ModalProps, "onOk" | "open" | "onCancel" | "title">, Pick<JSONEditorPropsOptional, "content" | "onChange"> {
    onImportHandler?: UploadProps["onChange"];
    onExportHandler?: () => void;
}

const ThemeVanillaJsonModal: FC<ThemeVanillaJsonModalProps> = (props) => {
    const {
        open, title = "主题配置", onCancel, onOk,
        onImportHandler, onExportHandler, content, onChange
    } = props;
    return (
        <Modal
            open={open}
            title={title}
            onCancel={onCancel}
            onOk={onOk}
            width={"50%"}
        >
            <ThemeVanillaJson content={content} onChange={onChange}/>
            <ThemeEditorJsonActionWrap>
                <Upload
                    customRequest={() => {
                    }}
                    maxCount={1}
                    showUploadList={false}
                    accept=".json"
                    onChange={onImportHandler}
                >
                    <Button style={{ marginRight: "8px" }}>
                        <ImportOutlined/>
                        导入
                    </Button>
                </Upload>
                <Button onClick={onExportHandler}>
                    <ExportOutlined/>
                    导出
                </Button>
            </ThemeEditorJsonActionWrap>
        </Modal>
    );
};

export default ThemeVanillaJsonModal;
