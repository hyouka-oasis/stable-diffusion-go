import { FC, useEffect, useRef } from "react";
import { JSONEditor, JSONEditorPropsOptional, Mode } from "vanilla-jsoneditor";
import { ReactSmoothScrollbar } from "../smooth-scroll/SmoothScroll.tsx";

const VanillaJson: FC<JSONEditorPropsOptional> = (props) => {
    const refContainer = useRef<HTMLDivElement | null>(null);
    const editorRef = useRef<JSONEditor | null>(null);
    useEffect(() => {
        editorRef.current = new JSONEditor({
            target: refContainer.current as HTMLDivElement,
            props: { mode: Mode.text, mainMenuBar: false },
        });
        return () => {
            editorRef.current?.destroy();
        };
    }, []);

    useEffect(() => {
        editorRef.current?.updateProps(props);
    }, [ props ]);
    return (
        <ReactSmoothScrollbar
            style={{
                maxHeight: "calc(100vh - 390px)",
            }}>
            <span/>
            <div ref={refContainer}/>
        </ReactSmoothScrollbar>
    );
};

export default VanillaJson;
