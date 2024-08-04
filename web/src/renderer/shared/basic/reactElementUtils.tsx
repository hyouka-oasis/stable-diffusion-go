import React from "react";

export interface InsertNavbarChildType {
    key: string;
    label?: string | React.ReactNode;
}

export interface ConversionReactChildOptions {
    selfChildren?: React.ReactNode;

    onTabClickHandler?(tabKey: string, event?: React.MouseEvent, child?: InsertNavbarChildType): void;
}

/**
 * 转化node
 * @param customChild
 * @param options
 */
export const conversionReactChildren = (customChild: InsertNavbarChildType | InsertNavbarChildType[] | React.ReactNode, options?: ConversionReactChildOptions) => {
    if (React.isValidElement(customChild)) {
        return customChild;
    }
    const arrayCustomChild = Array.isArray(customChild) ? customChild : [ customChild ];
    if (options?.selfChildren) {
        return options?.selfChildren;
    }
    return (arrayCustomChild as InsertNavbarChildType[]).map((child) => {
        return conversionReactChild(child.label, {
            key: child.key,
            onTabClickHandler: (tabKey, event) => options?.onTabClickHandler?.(tabKey, event, child)
        });
    });
};

export const conversionReactChild = (node: InsertNavbarChildType["label"], options: ConversionReactChildOptions & Pick<InsertNavbarChildType, "key">): React.ReactNode | null => {
    if (typeof node === "string") {
        return <React.Fragment key={options.key}>
            {<span onClick={(e) => options?.onTabClickHandler?.(options.key, e)}>{node}</span>}
        </React.Fragment>;
    }
    if (React.isValidElement(node)) {
        return React.cloneElement(node as React.ReactElement, { key: options.key, onClick: (e) => options?.onTabClickHandler?.(options.key, e) });
    }
    if (typeof node === "function") {
        return <span onClick={(e) => options?.onTabClickHandler?.(options.key, e)} key={options.key}>{node()}</span>;
    }
    return null;
};
