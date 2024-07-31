import { FC } from "react";
import { Modal, ModalProps } from "antd";
import VanillaUploadJson, { VanillaUploadJsonProps } from "./VanillaUploadJson.tsx";

interface ThemeVanillaJsonModalProps extends Pick<ModalProps, "onOk" | "open" | "onCancel" | "title">, VanillaUploadJsonProps {

}

const VanillaJsonModal: FC<ThemeVanillaJsonModalProps> = (props) => {
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
            <VanillaUploadJson content={content} onChange={onChange} onImportHandler={onImportHandler} onExportHandler={onExportHandler}/>
        </Modal>
    );
};

export default VanillaJsonModal;
