import { ArgsProps, IconType, NotificationPlacement } from "antd/es/notification/interface";
import React from "react";

export type MessageType = "notification" | "message";

export interface MessageBoxConfigProps extends Omit<ArgsProps, "message" | "description"> {
    /**
     * 提示文案
     */
    message?: string | Error | React.ReactNode;
    /**
     * 提示内容
     */
    description?: string;
    /**
     * 显示位置
     */
    placement?: NotificationPlacement;
    /**
     * 提示类型
     */
    type?: IconType;
    /**
     * 是否需要关闭
     */
    closeIcon?: boolean | React.ReactNode;
    /**
     * 通知框类型
     */
    messageType?: MessageType;
    /**
     * 样式
     */
    messageBoxStyle?: React.CSSProperties;
    /**
     * 类名
     */
    messageBoxClassname?: string;
    /**
     * 通知提醒框底部按钮
     */
    notificationBtn?: React.ReactNode;
}


export interface AppGlobalContextProps {

    openMessageBox(config: MessageBoxConfigProps): (() => void) | void;
}

const AppGlobalContext = React.createContext<AppGlobalContextProps>({
    openMessageBox(config: MessageBoxConfigProps) {
    },
});


export {
    AppGlobalContext
};
