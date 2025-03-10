import styled from "styled-components";
import { ProColumns, ProTable } from "@ant-design/pro-components";
import { useEffect, useState } from "react";
import { Button } from "antd";
import { getFileList } from "renderer/api/fileApi";
import { LorasResponse } from "renderer/api/response/lorasResponse";
import { host } from "renderer/request/request";

const FilePageWrap = styled.div`
`;

const FilePage = () => {
    const [ fileList, setFileList ] = useState<any[]>([]);


    const getFileListHandler = async () => {
        const list = await getFileList({
            page: 1,
            pageSize: -1,
        });
        setFileList(list.list);
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
            title: "附件名称",
        },
        {
            dataIndex: "url",
            title: "图片",
            render(value) {
                return (
                    <img
                        style={{
                            width: "100px",
                            height: "100px"
                        }}
                        src={`${host()}/${value}`}></img>
                );
            }
        },
        {
            dataIndex: "action",
            title: "操作",
            align: "center",
            width: 70,
            render() {
                return (
                    <Button>
                        删除
                    </Button>
                );
            }
        },
    ];

    useEffect(() => {
        getFileListHandler();
    }, []);
    return (
        <FilePageWrap>
            <ProTable
                rowKey={"id"}
                dataSource={fileList ?? []}
                columns={columns}
                search={false}
                scroll={{ y: 650, x: 800 }}
                pagination={false}
            />
        </FilePageWrap>
    );
};

export default FilePage;
