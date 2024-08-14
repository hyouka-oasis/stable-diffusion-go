import { Button, Upload, UploadProps } from "antd";
import { ExportOutlined, ImportOutlined } from "@ant-design/icons";
import { JSONEditorPropsOptional } from "vanilla-jsoneditor";
import { FC } from "react";
import styled from "styled-components";
import VanillaJson from "renderer/components/json-edit/VanillaJson";

const VanillaUploadJsonActionWrap = styled.div`
    margin-top: 24px;
    margin-bottom: -44px;
`;

export interface VanillaUploadJsonProps extends Pick<JSONEditorPropsOptional, "content" | "onChange"> {
    onImportHandler?: UploadProps["onChange"];
    onExportHandler?: () => void;
}

const VanillaUploadJson: FC<VanillaUploadJsonProps> = (props) => {
    const { content, onChange, onImportHandler, onExportHandler } = props;
    return (
        <>
            <VanillaJson content={content} onChange={onChange}/>
            <VanillaUploadJsonActionWrap>
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
            </VanillaUploadJsonActionWrap>
        </>
    );
};

export default VanillaUploadJson;
